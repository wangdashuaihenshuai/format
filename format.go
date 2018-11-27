package format

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Match struct {
	Method string      `json:"method"`
	Value  interface{} `json:"value"`
}

type BasicType struct {
	Type         string        `json:"type"`
	Enum         []interface{} `json:"enum"`
	Optional     bool          `json:"optional"`
	DefaultValue interface{}   `json:"defaultValue"`
	Rename       string        `json:"rename"`
	EnumFunc     func(data interface{}) bool
	EnumStr      string

	Match     *Match         `json:"match"`
	Select    string         `json:"select"`
	Formats   []*ValueFormat `json:"formats"`
	MatchFunc func(parent interface{}, data interface{}) (*ValueFormat, error)
}

type MapType struct {
	IsFilter bool                    `json:"isFilter"`
	Fields   map[string]*ValueFormat `json:"fields"`
}

type ArrayType struct {
	Format *ValueFormat
}

type BoolType struct {
}

type NumberType struct {
}

type StringType struct {
}

type Formater interface {
	FormatData(data interface{}) (interface{}, error)
}

type ValueFormat struct {
	BasicType
	MapType
	ArrayType
	BoolType
	NumberType
	StringType
}

func (vf *ValueFormat) Init() error {
	err := isInvaildType(vf.Type)
	if err != nil {
		return err
	}
	for i, f := range vf.Formats {
		err := f.Init()
		if err != nil {
			return fmt.Errorf(".formats.$%d%s", i, err.Error())
		}
	}
	for name, f := range vf.Fields {
		err := f.Init()
		if err != nil {
			return fmt.Errorf(".%s%s", name, err.Error())
		}
	}
	if vf.Format != nil {
		err := vf.Format.Init()
		if err != nil {
			return fmt.Errorf(".format%s", err.Error())
		}
	}

	if len(vf.Enum) > 0 {
		EnumMap := make(map[interface{}]bool, len(vf.Enum))
		for i, value := range vf.Enum {
			err := expectType(value, vf.Type)
			if err != nil {
				return fmt.Errorf(" enum type error index %d err: %s", i, err.Error())
			}
			EnumMap[value] = true
		}

		bs, err := json.Marshal(vf.Enum)
		if err != nil {
			return err
		}
		vf.Enum = string(bs)

		vf.EnumFunc = func(data interface{}) bool {
			_, ok := EnumMap[data]
			return ok
		}
	}

	if vf.DefaultValue != nil {
		err := expectType(vf.DefaultValue, vf.Type)
		if err != nil {
			return fmt.Errorf(" defaultValue type error: %s", err.Error())
		}
	}

	if (vf.Select == "" && len(vf.Formats) > 0) || (vf.Select != "" && len(vf.Formats) <= 0) {
		return errors.New(" select or formats config error")
	}

	var selectFunc func(parent interface{}, data interface{}) (interface{}, error)
	if vf.Select == "" {
		selectFunc = func(parent interface{}, data interface{}) (interface{}, error) {
			return data, nil
		}
	} else {
		names := strings.Split(vf.Select, "->")
		namesLength := len(names)
		dataName := "data"

		if namesLength > 1 {
			dataName = names[0]
			names = names[1:]
		}

		selectFunc = func(parent interface{}, data interface{}) (interface{}, error) {
			switch dataName {
			case "data":
				return GetMapValueByNames(data, names)
			case "parent":
				return GetMapValueByNames(parent, names)
			default:
				return nil, fmt.Errorf(" select data key %s not find", dataName)
			}
		}
	}

	length := len(vf.Formats)
	if length <= 1 {
		return nil
	}

	matchEqualMap := make(map[interface{}]*ValueFormat, length)
	for _, format := range vf.Formats {
		if format.Match == nil {
			return errors.New(" formats match is empty")
		}

		switch format.Match.Method {
		case "equal":
			matchEqualMap[format.Match.Value] = format
		default:
			return fmt.Errorf(" not support match method %s", format.Match.Method)
		}
	}
	if len(matchEqualMap) <= 0 {
		return errors.New(" match map is empty")
	}

	vf.MatchFunc = func(parent interface{}, data interface{}) (*ValueFormat, error) {
		value, err := selectFunc(parent, data)
		if err != nil {
			return nil, err
		}
		format, ok := matchEqualMap[value]
		if ok {
			return format, nil
		}
		return nil, errors.New(" not match format")
	}
	return nil
}

func TypeError(expect string, _type string) error {
	return fmt.Errorf(" expect type %s but %s", expect, _type)
}

func (vf *ValueFormat) MatchFormat(parent interface{}, data interface{}) (*ValueFormat, error) {
	length := len(vf.Formats)
	if length == 0 {
		return vf, nil
	}
	if length == 1 {
		return vf.Formats[0], nil
	}
	if vf.MatchFunc == nil {
		return nil, errors.New(" lack of match function")
	}
	return vf.MatchFunc(parent, data)
}

func FormatTypes(format *ValueFormat, _type string) error {
	if format.Type != _type {
		return TypeError(format.Type, _type)
	}
	return nil
}

func FormatEnum(format *ValueFormat, data interface{}) error {
	if len(format.Enum) <= 0 {
		return nil
	}
	if format.EnumFunc == nil {
		return errors.New(" Enum func is not init")
	}
	exist := format.EnumFunc(data)
	if !exist {
		return fmt.Errorf(" value is not in Enum %s", format.EnumStr)
	}
	return nil
}

func (vf *ValueFormat) BoolFormat(data bool, format *ValueFormat) (interface{}, error) {
	err := FormatTypes(format, BOOL)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (vf *ValueFormat) NumberFormat(data interface{}, format *ValueFormat) (interface{}, error) {
	err := FormatTypes(format, NUMBER)
	if err != nil {
		return nil, err
	}
	err = FormatEnum(format, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (vf *ValueFormat) StringFormat(data string, format *ValueFormat) (interface{}, error) {
	err := FormatTypes(format, STRING)
	if err != nil {
		return nil, err
	}
	err = FormatEnum(format, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (vf *ValueFormat) ArrayFormat(data []interface{}, format *ValueFormat) (interface{}, error) {
	err := FormatTypes(format, ARRAY)
	if err != nil {
		return nil, err
	}
	if format.Format == nil {
		return nil, errors.New(" array format is empty")
	}

	res := make([]interface{}, 0, len(data))
	for i, d := range data {
		value, err := format.Format.format(data, d)
		if err != nil {
			return nil, fmt.Errorf(".$%d%s", i, err.Error())
		}
		res = append(res, value)
	}
	return res, nil
}

func (vf *ValueFormat) MapFormat(data map[string]interface{}, format *ValueFormat) (interface{}, error) {
	err := FormatTypes(format, MAP)
	if err != nil {
		return nil, err
	}
	var res = data
	if format.IsFilter {
		res = make(map[string]interface{}, len(data))
	}
	for key, filedFormat := range format.Fields {
		oldKey := key
		oldValue, ok := data[oldKey]
		if filedFormat.Rename != "" {
			key = filedFormat.Rename
		}
		if !ok {
			if !filedFormat.Optional && filedFormat.DefaultValue == nil {
				return nil, fmt.Errorf(".%s not exists", key)
			}
			if filedFormat.DefaultValue != nil {
				res[key] = filedFormat.DefaultValue
			}
		} else {
			value, err := filedFormat.format(data, oldValue)
			if err != nil {
				return nil, fmt.Errorf(".%s%s", oldKey, err.Error())
			}
			res[key] = value
		}
	}
	return res, nil
}

func (vf *ValueFormat) FormatData(data interface{}) (interface{}, error) {
	return vf.format(data, data)
}

func (vf *ValueFormat) format(parent interface{}, data interface{}) (interface{}, error) {
	format, err := vf.MatchFormat(parent, data)
	if err != nil {
		return nil, err
	}
	switch value := data.(type) {
	case string:
		return vf.StringFormat(value, format)
	case float32, float64, int, int8, int16, int32, int64:
		return vf.NumberFormat(value, format)
	case bool:
		return vf.BoolFormat(value, format)
	case []interface{}:
		return vf.ArrayFormat(value, format)
	case map[string]interface{}:
		return vf.MapFormat(value, format)
	default:
		return nil, errors.New(" unkonw parse data type")
	}
}

func NewFormater(config string) (Formater, error) {
	var formater *ValueFormat
	err := json.Unmarshal([]byte(config), &formater)
	if err != nil {
		return nil, err
	}
	err = formater.Init()
	if err != nil {
		return nil, err
	}
	return formater, nil
}

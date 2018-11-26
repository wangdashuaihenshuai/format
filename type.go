package format

import "fmt"

const (
	NUMBER = "number"
	STRING = "string"
	BOOL   = "bool"
	MAP    = "map"
	ARRAY  = "array"
	UNKONW = "unknow"
)

var TYPES = []string{
	NUMBER,
	STRING,
	BOOL,
	MAP,
	ARRAY,
}

func typeOf(data interface{}) string {
	switch data.(type) {
	case string:
		return STRING
	case float32, float64, int, int8, int16, int32, int64:
		return NUMBER
	case bool:
		return BOOL
	case []interface{}:
		return ARRAY
	case map[string]interface{}:
		return MAP
	default:
		return UNKONW
	}
}

func isType(data interface{}, _type string) (bool, error) {
	err := isInvaildType(_type)
	if err != nil {
		return false, err
	}
	t := typeOf(data)
	return t == _type, nil
}

func expectType(data interface{}, _type string) error {
	ok, err := isType(data, _type)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("expect type %s but not", _type)
	}
	return nil
}

func isInvaildType(_type string) error {
	for _, t := range TYPES {
		if t == _type {
			return nil
		}
	}
	return fmt.Errorf("unknow type %s", _type)
}

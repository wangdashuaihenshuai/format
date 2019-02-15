package format

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func toString(value interface{}) (string, error) {
	fmt.Println(value, reflect.TypeOf(value))
	switch v := value.(type) {
	case string:
		return v, nil
	case int:
		return strconv.FormatInt(int64(v), 10), nil
	case int8:
		return strconv.FormatInt(int64(v), 10), nil
	case int16:
		return strconv.FormatInt(int64(v), 10), nil
	case int32:
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case float32:
		return strconv.FormatFloat(float64(v), 'g', -1, 64), nil
	case float64:
		return strconv.FormatFloat(float64(v), 'g', -1, 64), nil
	case bool:
		if v {
			return "true", nil
		}
		return "false", nil
	default:
		bs, err := json.Marshal(value)
		if err != nil {
			return "", err
		}
		return string(bs), nil
	}
	return "", errors.New(" toString unkonw type error")
}

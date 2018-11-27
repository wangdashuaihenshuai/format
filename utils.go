package format

import "fmt"

func GetMapValue(m interface{}, name string) (interface{}, error) {
	newM, ok := m.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf(" value is no map can't get map value by key %s", name)
	}
	if val, ok := newM[name]; ok {
		return val, nil
	}
	return nil, fmt.Errorf(" map key %s is not exist", name)
}

func GetMapValueByNames(m interface{}, names []string) (interface{}, error) {
	selectValue := m
	var err error
	for _, name := range names {
		selectValue, err = GetMapValue(m, name)
		if err != nil {
			return nil, err
		}
	}
	return selectValue, err
}

func GetTypeDefualtValue(_type string) interface{} {
	return nil
}

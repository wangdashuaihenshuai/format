package format

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewFomaterEmptyError(t *testing.T) {
	formatStr := `{}`
	_, err := NewFormater(formatStr)
	fmt.Println(err)
	if err == nil {
		t.Error("error")
	}
}

// type error
func TestNewFomaterTypeError(t *testing.T) {
	formatStr := `{"type": "test"}`
	_, err := NewFormater(formatStr)
	fmt.Println(err)
	if err == nil {
		t.Error("error")
	}
}

// formats must be select and match is not empty
func TestNewFomaterSelectError(t *testing.T) {
	formatStr := `{
		"type": "number",
		"formats": [
			{
				"match": {
					"method": "equal",
					"value": "test"
				},
				"type": "string"
			}
		]
	}`
	_, err := NewFormater(formatStr)
	fmt.Println(err)
	if err == nil {
		t.Error("error")
	}
}

// formats success
func TestNewFomaterSelectSucess(t *testing.T) {
	formatStr := `{
		"type": "map",
		"select": "data->test",
		"formats": [
			{
				"match": {
					"method": "equal",
					"value": "test"
				},
				"type": "map",
				"fields": {
					"n": {
						"type": "number"
					}
				}	
			}
		]
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data := map[string]interface{}{"n": true, "test": "test"}
	_, err = f.FormatData(data)
	fmt.Println(err)
	if err == nil {
		t.Error("error......")
	}
}

// formats success
func TestNewFomaterSelectEmpty(t *testing.T) {
	formatStr := `{
		"type": "array",
		"format": {
			"type": "map",
			"select": "data->test",
			"formats": [
				{
					"match": {
						"method": "equal",
						"value": "test"
					},
					"type": "map",
					"fields": {
						"n": {
							"type": "number"
						}
					}	
				}
			]
		}
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	m1 := map[string]interface{}{"n": 1, "test": "test"}
	data := []interface{}{m1, m1}
	_, err = f.FormatData(data)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestNewFomaterSelectMatchError(t *testing.T) {
	formatStr := `{
		"type": "array",
		"format": {
			"type": "map",
			"select": "data->test",
			"formats": [
				{
					"match": {
						"method": "equal",
						"value": "t"
					},
					"type": "map",
					"fields": {
						"n": {
							"type": "number"
						}
					}	
				}
			]
		}
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	m1 := map[string]interface{}{"n": 1, "test": "test"}
	data := []interface{}{m1, m1}
	_, err = f.FormatData(data)
	if err == nil {
		t.Error("match formater emtpy error")
	}
}

func TestNewFomaterSelectMatchFilterSuccess(t *testing.T) {
	formatStr := `{
		"type": "array",
		"isFilter": true,
		"format": {
			"type": "map",
			"select": "data->test",
			"formats": [
				{
					"match": {
						"method": "equal",
						"value": "t"
					},
					"type": "map",
					"fields": {
						"n": {
							"type": "number"
						}
					}	
				}
			]
		}
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	m1 := map[string]interface{}{"n": 1, "test": "test"}
	m2 := map[string]interface{}{"n": 2, "test": "t"}
	data := []interface{}{m1, m2}
	r, err := f.FormatData(data)
	if err != nil || len(r.([]interface{})) != 1 {
		t.Error("filter array error")
	}
}
func TestNewFomaterMatchError(t *testing.T) {
	formatStr := `{
		"type": "number",
		"formats": [
			{
				"type": "number"
			},
			{
				"type": "string"
			}
		]
	}`
	_, err := NewFormater(formatStr)
	fmt.Println(err)
	if err == nil {
		t.Error("error")
	}
}

// test basic type
func TestFomaterNumberError(t *testing.T) {
	formatStr := `{
		"type": "map",
		"fields": {
			"n": {
				"type": "number"
			}
		}
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data := map[string]interface{}{"n": "111"}
	_, err = f.FormatData(data)
	fmt.Println(err)
	if err == nil {
		t.Error("error")
	}
}
func TestFomaterStringError(t *testing.T) {
	formatStr := `{
		"type": "map",
		"fields": {
			"n": {
				"type": "string"
			}
		}
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data := map[string]interface{}{"n": true}
	_, err = f.FormatData(data)
	fmt.Println(err)
	if err == nil {
		t.Error("error")
	}
}
func TestFomaterBoolError(t *testing.T) {
	formatStr := `{
		"type": "map",
		"fields": {
			"n": {
				"type": "bool"
			}
		}
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data := map[string]interface{}{"n": "111"}
	_, err = f.FormatData(data)
	fmt.Println(err)
	if err == nil {
		t.Error("error")
	}
}
func TestFomaterBasicSuccess(t *testing.T) {
	formatStr := `{
		"type": "map",
		"fields": {
			"n": {
				"type": "number"
			},
			"b": {
				"type": "bool"
			},
			"s": {
				"type": "string"
			}
		}
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data := map[string]interface{}{"n": 111, "s": "test", "b": true}
	_, err = f.FormatData(data)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestFomaterArraySuccess(t *testing.T) {
	formatStr := `{
		"type": "array",
		"format": {
			"type": "map",
			"fields": {
				"n": {
					"type": "number"
				},
				"b": {
					"type": "bool"
				},
				"s": {
					"type": "string"
				}
			}
		}
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	d := map[string]interface{}{"n": 111, "s": "test", "b": true}
	data := []interface{}{d}
	_, err = f.FormatData(data)
	fmt.Println(err)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestNewFomaterDefaultValueError(t *testing.T) {
	formatStr := `{
		"type": "map",
		"fields": {
			"n": {
				"type": "number",
				"optional": true,
				"default": "11"
			},
			"b": {
				"type": "bool"
			},
			"s": {
				"type": "string"
			}
		}
	}`
	_, err := NewFormater(formatStr)
	fmt.Println(err)
	if err == nil {
		t.Error("error")
	}
}
func TestNewFomaterEnumError(t *testing.T) {
	formatStr := `{
		"type": "map",
		"fields": {
			"n": {
				"type": "string",
				"enum": ["test1", 11]
			}
		}
	}`
	_, err := NewFormater(formatStr)
	fmt.Println(err)
	if err == nil {
		t.Error("error")
	}
}
func TestFomaterEnumError(t *testing.T) {
	formatStr := `{
		"type": "map",
		"fields": {
			"n": {
				"type": "string",
				"enum": ["test1", "test2"]
			}
		}
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data := map[string]interface{}{"n": "test3"}
	_, err = f.FormatData(data)
	fmt.Println(err)
	if err == nil {
		t.Error("error")
	}
}

func TestFomaterEnumSuccess(t *testing.T) {
	formatStr := `{
		"type": "map",
		"fields": {
			"n": {
				"type": "string",
				"enum": ["test1", "test2"]
			}
		}
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data := map[string]interface{}{"n": "test2"}
	_, err = f.FormatData(data)
	fmt.Println(err)
	if err != nil {
		t.Error(err.Error())
	}
}
func TestFomaterRenameSuccess(t *testing.T) {
	formatStr := `{
		"type": "map",
		"fields": {
			"n": {
				"type": "number",
				"rename": "n2"
			},
			"b": {
				"type": "bool"
			},
			"s": {
				"type": "string"
			}
		}
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data := map[string]interface{}{"n": 111, "s": "test", "b": true}
	newData, err := f.FormatData(data)
	if err != nil {
		t.Error(err.Error())
		return
	}
	newMap, ok := newData.(map[string]interface{})
	if !ok {
		t.Error("return new Data is not map")
		return
	}
	_, ok = newMap["n2"]
	if !ok {
		t.Error("rename data n => n2 error")
		return
	}
}

func TestFomaterDefaultValueSuccess(t *testing.T) {
	formatStr := `{
		"type": "map",
		"fields": {
			"n": {
				"type": "number"
			},
			"b": {
				"type": "bool"
			},
			"s": {
				"type": "string",
				"default": "11"
			}
		}
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data := map[string]interface{}{"n": 11, "b": true}
	newData, err := f.FormatData(data)
	if err != nil {
		t.Error(err.Error())
		return
	}
	newMap, ok := newData.(map[string]interface{})
	if !ok {
		t.Error("return new Data is not map")
		return
	}
	r, ok := newMap["s"]
	if !ok {
		t.Error("get value s error")
		return
	}
	if r.(string) != "11" {
		t.Error("set default value error")
		return
	}
}

func TestFomaterOptionalSuccess(t *testing.T) {
	formatStr := `{
		"type": "map",
		"fields": {
			"n": {
				"type": "number"
			},
			"b": {
				"type": "bool"
			},
			"s": {
				"type": "string",
				"optional": true
			}
		}
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data := map[string]interface{}{"n": 11, "b": true}
	newData, err := f.FormatData(data)
	if err != nil {
		t.Error(err.Error())
		return
	}
	newMap, ok := newData.(map[string]interface{})
	if !ok {
		t.Error("return new Data is not map")
		return
	}
	_, ok = newMap["s"]
	if ok {
		t.Error("optional true error")
		return
	}
}

func TestFomaterIsFilterFalse(t *testing.T) {
	formatStr := `{
		"type": "map",
		"fields": {
			"n": {
				"type": "number"
			},
			"b": {
				"type": "bool"
			},
			"s": {
				"type": "string"
			}
		}
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data := map[string]interface{}{"n": 111, "s": "test", "b": true, "c": "hello"}
	newData, err := f.FormatData(data)
	if err != nil {
		t.Error(err.Error())
	}
	newMap, ok := newData.(map[string]interface{})
	if !ok {
		t.Error("return new Data is not map")
		return
	}
	_, ok = newMap["c"]
	if !ok {
		t.Error("get extrace key c error")
		return
	}
}

func TestFomaterIsFilterTrue(t *testing.T) {
	formatStr := `{
		"type": "map",
		"isFilter": true,
		"fields": {
			"n": {
				"type": "number"
			},
			"b": {
				"type": "bool"
			},
			"s": {
				"type": "string"
			}
		}
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data := map[string]interface{}{"n": 111, "s": "test", "b": true, "c": "hello"}
	newData, err := f.FormatData(data)
	if err != nil {
		t.Error(err.Error())
	}
	newMap, ok := newData.(map[string]interface{})
	if !ok {
		t.Error("return new Data is not map")
		return
	}
	_, ok = newMap["c"]
	if ok {
		t.Error("filter extrace key c error")
		return
	}
}

func TestFomaterMapGetFormatSuccess(t *testing.T) {
	formatStr := `{
		"type": "map",
		"fields": {
			"n": {
				"type": "number"
			},
			"b": {
				"type": "bool"
			},
			"s": {
				"type": "string"
			}
		}
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data := map[string]interface{}{"n": 111, "s": "test", "b": true}
	_, err = f.FormatData(data)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestFomaterToStringSuccess(t *testing.T) {
	formatStr := `{
		"type": "map",
		"fields": {
			"n": {
				"toString": true,
				"type": "number"
			},
			"b": {
				"toString": true,
				"type": "bool"
			},
			"s": {
				"toString": true,
				"type": "string"
			}
		}
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data := map[string]interface{}{"n": 111, "s": "test", "b": true}
	res, err := f.FormatData(data)
	if err != nil {
		t.Error(err.Error())
	}
	bs, err := json.Marshal(res)
	if err != nil {
		t.Error(err.Error())
	}
	if string(bs) != `{"b":"true","n":"111","s":"test"}` {
		t.Error("toString errror")
	}
	return
}

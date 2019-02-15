# format

# BENCH

```bash
go test -bench=.
```

#  TEST

```bash
go test
```

#TODO

- 添加 any 类型
- 添加类型转换函数

## 校验格式

### Map

类型 `map[string]interface{}`

- type "map"
- fileds map[string]Format 格式化配置 map
- ?isFilter `bool` 是否过滤 key default: false
- ?rename `string` 更换名字
- ?default `any` 默认参数
- ?optional `bool` 是否可选

**例子**

注意 default 的参数格式必须和 type 相同。

```json
{
  "n": 111,
  "b": true,
  "s": "test"
}
```

```json
{
	"type": "map",
  "fields": {
    "n": {
      "type": "number",
      "optional": true,
      "default": 11
    },
    "b": {
      "type": "bool"
    },
    "s": {
      "type": "string"
    }
  }
}
```

### Array

类型 `[]interface{}`

- type "array"
- format Format 格式化配置
- ?isFilter 过滤格式化错误数据

```json
[
  {
    "n": 111,
    "b": true,
    "s": "test"
  }
]
```
```json
{
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
}
```
### Bool

类型 `bool`

- type      "bool"
- ?toString bool 类型转换为string

```json
true
```

例子
```json
{
  "type": "bool"
}
```
### Number

类型 `float32, float64, int, int8, int16, int32, int64`

- type "number"
- ?enum 数字枚举
- ?toString bool 类型转换为string

```json
1
```

```json
{
  "type": "number",
  "enum": [1,2,3]
}
```
### String

类型 `string`

- type "string"
- ?enum 数字枚举

```json
"1"
```

```json
{
  "type": "string",
  "enum": ["1", "2", "3"]
}
```

# FAQ

1. YAML 转为 golang json map 类型为 map[interface{}]interface{} 会出现格式化错误问题，需要将 YAML 转为 JSON 格式，之后转为 golang json 才可以。
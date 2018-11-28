package format

import (
	"encoding/json"
	"testing"
)

type A struct {
	L []M `json:"l"`
}

type M struct {
	F T `json:"f"`
}

type T struct {
	N float64 `json:"n"`
	B bool    `json:"b"`
	S string  `json:"s"`
}

func CreateData(num int) (string, error) {
	t := T{N: 1, B: true, S: "dfasdfasfd"}
	m := M{F: t}
	l := make([]M, num)
	for index := 0; index < num; index++ {
		l = append(l, m)
	}
	a := A{L: l}
	rb, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(rb), nil
}

func Other(str string) error {
	a := A{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		return err
	}
	return nil
}

func InitFormater() Formater {
	var formatStr = `{
		"type": "map",
		"fields": {
			"l": {
				"type": "array",
				"format": {
					"type": "map",
					"fields": {
						"f": {
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
				}
			}
		}
	}`
	f, err := NewFormater(formatStr)
	if err != nil {
		panic(err)
	}
	return f
}

var fomater = InitFormater()

func Format(str string) error {
	var a interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		return err
	}
	_, err = fomater.FormatData(a)
	if err != nil {
		return err
	}
	return nil
}

func benchmarkFormat(b *testing.B, num int) {
	str, err := CreateData(num)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := Format(str)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func benchmarkOther(b *testing.B, num int) {
	str, err := CreateData(num)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := Other(str)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkOther1(b *testing.B) {
	benchmarkOther(b, 1)
}

func BenchmarkOther16(b *testing.B) {
	benchmarkOther(b, 16)
}

func BenchmarkOther32(b *testing.B) {
	benchmarkOther(b, 32)
}
func BenchmarkOther128(b *testing.B) {
	benchmarkOther(b, 128)
}

func BenchmarkFormat1(b *testing.B) {
	benchmarkFormat(b, 1)
}

func BenchmarkFormat16(b *testing.B) {
	benchmarkFormat(b, 16)
}

func BenchmarkFormat32(b *testing.B) {
	benchmarkFormat(b, 32)
}
func BenchmarkFormat128(b *testing.B) {
	benchmarkFormat(b, 128)
}

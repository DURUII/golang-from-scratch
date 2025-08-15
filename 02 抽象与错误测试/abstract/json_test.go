package abstract

import (
	"encoding/json"
	"fmt"
	"testing"
)

type response struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func TestStructUnmarshalling(t *testing.T) {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response{}
	// 反序列化时入参是地址
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		t.Fail()
	}
	fmt.Println(res)
	resB, _ := json.Marshal(res)
	fmt.Println(resB)
	fmt.Println(string(resB))
}

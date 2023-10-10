package utils

import (
	"encoding/json"
	"fmt"
)

func PrintJson[T any](item T) {
	val, _ := json.MarshalIndent(item, "", "    ")
	fmt.Println(string(val))
}

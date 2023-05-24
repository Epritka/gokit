package main

import (
	"encoding/json"
	"fmt"

	"github.com/Epritka/gokit/example"
)

func main() {
	err := example.Create()
	if err != nil {
		bytes, _ := json.Marshal(err)
		fmt.Println(string(bytes))
	}
}

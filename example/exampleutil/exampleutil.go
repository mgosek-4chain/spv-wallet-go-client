package exampleutil

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func Print(s string, a any) {
	fmt.Println(strings.Repeat("~", 100))
	fmt.Println(s)
	fmt.Println(strings.Repeat("~", 100))
	res, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res))
}

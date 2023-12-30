package main

import (
	"fmt"

	fstr "github.com/ZiadMansourM/fstr/Interpolate"
)

func main() {
	data := map[string]interface{}{
		"name": "Ziad Mansour",
		"age":  23,
	}
	result, err := fstr.Interpolate("My name is {name} and I am {age} years old", data)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

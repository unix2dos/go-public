package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	NameValue string `mapstructure:"name_value"`
	AgeValue  int    `mapstructure:"age_value"`
	JobValue  string `mapstructure:"job_value,omitempty"`
}

func main() {
	p := &Person{
		NameValue: "dj",
		AgeValue:  18,
	}

	var m map[string]interface{}
	mapstructure.Decode(p, &m)

	fmt.Println("map", m)

	data, _ := json.Marshal(m)
	fmt.Println("str", string(data))
}

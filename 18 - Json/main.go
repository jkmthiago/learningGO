package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
	Idade uint `json:"idade"`
}

func main() {
	c := cachorro{"Ralf", "Dashhound", 9}
	fmt.Println(c)

	cJSON, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cJSON)

	fmt.Println(bytes.NewBuffer(cJSON))

	c2 := map[string]string{
		"nome": "Celaena",
		"raca": "Pincher",
		"idade": "2",
	}

	c2JSON, err2 := json.Marshal(c2)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(c2JSON)

	fmt.Println(bytes.NewBuffer(c2JSON))

	c3JSON := `{"nome":"Ralf","raca":"Dashhound","idade":9}`

	var c3 cachorro

	if err3 := json.Unmarshal([]byte(c3JSON), &c3); err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println(c3)
}
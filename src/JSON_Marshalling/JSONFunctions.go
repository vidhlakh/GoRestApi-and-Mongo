package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Message struct {
	Name string `json : "names`
	Body string `json : "body-content"`
	Time int64  `json : "time"`
}

func main() {
	M := Message{"Vidhya", "How are you", 2}
	fmt.Println(M, "Message struct")
	//Marshalling
	json_content, err := json.Marshal(M)
	if err != nil {
		fmt.Println("error in marshalling ")
	}
	fmt.Println("Json content", json_content)

	//Unmarshalling message struct
	var unM Message
	err2 := json.Unmarshal(json_content, &unM)
	if err2 != nil {
		fmt.Println("error in Unmarshalling ")
	}
	fmt.Println("After Unmarshalling:", unM)

	//Usage of encoder and Decoder
	dec := json.NewDecoder("testfile.txt")
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			if k != "Name" {
				delete(v, k)
			}
		}
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}

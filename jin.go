package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	var j interface{}
	err := json.Unmarshal(bytes, &j)
	if err != nil {
		fmt.Println("Not valid JSON")
		return
	} else {
		res, err := json.MarshalIndent(j, " ", "    ")
		if err == nil {
			fmt.Printf("%s\n", string(res))
		}
	}
}

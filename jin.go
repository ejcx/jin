package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var file = flag.String("f", "", "-f=file")
	flag.Parse()
	var bytes []byte
	if len(*file) == 0 {
		bytes, _ = ioutil.ReadAll(os.Stdin)
	} else {
		bytes, _ = ioutil.ReadFile(*file)
	}
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

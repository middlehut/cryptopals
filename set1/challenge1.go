package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	hexBytes, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}
	result := base64.StdEncoding.EncodeToString(hexBytes)
	fmt.Println(string(result))
}


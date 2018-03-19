package main

import (
	"encoding/hex"
	"log"
	"fmt"
)

func main() {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"

	num1, err := hex.DecodeString(input1)
	if err != nil {
		log.Fatal(err)
	}

	num2, err := hex.DecodeString(input2)
	if err != nil {
		log.Fatal(err)
	}

	result := make([]byte, len(num1))
	for i := 0; i < len(num1); i++ {
		result[i] = num1[i] ^ num2[i]
	}

	resultHex := hex.EncodeToString(result)
	if resultHex != expected {
		log.Fatal(fmt.Printf("error: xored value is different from expected: xored result = %v, expected = %v\n", resultHex, expected))
	}

	fmt.Println(resultHex)
}
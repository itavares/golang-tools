/*
Author: Ighor Tavares (@sum_b0dy)
Date: 23 August 2022

This program is used to perform XOR operation on binary files (read byte array and perform xor given key)
*/

package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){

	inputFile := flag.String("in","","Input File. (ie. example.exe)")
	outputFile := flag.String("out","","Output File. (ie. output.exe)")
	flag.Parse()

	//XOR key
	key := 0xAA

	//read the file
	b, errR := os.ReadFile(*inputFile)
	if errR != nil{
		fmt.Fprintf(os.Stderr, "xored: %v\n", errR)
		os.Exit(1)
	}

	//type byte == uint8 (byte is alias to uint8)
	// fmt.Printf("%T\n",b)
	// fmt.Printf("%X\n", key)
	// fmt.Printf("%d\n",uint8(key))

	//make array of byte to store the XORed bytes
	var byteArray []uint8 

	for _, bt := range b {
		xorOp := bt ^ uint8(key)
		byteArray = append(byteArray, xorOp)
	}

	//Write the byteArray to the new binary with given output name from flag
	errW := os.WriteFile(*outputFile, byteArray, 0644)
	if errW != nil{
		fmt.Fprintf(os.Stderr, "xored: %v\n", errW)
		os.Exit(1)
	}


}
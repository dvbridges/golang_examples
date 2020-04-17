/* 
Package for converting betweel C and F temps

It mainly demonstrates:
	Packages, files and imports
	Type declarations
	Adding methods to types
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"tempConv/helpers"
)


func main () {
	// Get args from cmdline
	args := os.Args[1:]
	FtoC := tempConv.FtoC
	CtoF := tempConv.CtoF

	for _, t := range args {

		v, err := strconv.ParseFloat(t, 64)
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("%v C = %v\n", t, CtoF(v))
		fmt.Printf("%v F = %v\n", t, FtoC(v))
	}


}
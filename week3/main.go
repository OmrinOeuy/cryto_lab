package main

import (
	"fmt"
	"os"

	"hashingprogram/utils/crack"
)

func main() {

	targetHash := "6a85dfd77d9cb35770c9dc6728d73d3f"

	wordlistPath := "wordlists/nord_vpn.txt"

	output, err := os.Create("verbose.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer output.Close()

	fmt.Println("Starting MD5 password cracking...")

	err = crack.CrackMD5(wordlistPath, targetHash, output)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Cracking completed!")
	fmt.Println("Verbose output saved to verbose.txt")
}

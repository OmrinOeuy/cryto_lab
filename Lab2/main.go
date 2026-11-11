package main

import (
	"fmt"
	"lab3/utils/crack"
)

func main() {

	targetHash := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"

	wordlist := "wordlists/nord_vpn.txt"

	fmt.Println("================================")
	fmt.Println("       SHA1 PASSWORD CRACKER")
	fmt.Println("================================")

	fmt.Println("[*] Target Hash:", targetHash)
	fmt.Println("[*] Wordlist:", wordlist)
	fmt.Println("[*] Starting password cracking...")
	fmt.Println()

	password, err := crack.CrackSHA1(targetHash, wordlist)

	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	if password != "" {
		fmt.Println("\n[SUCCESS] Cracking completed!")
	} else {
		fmt.Println("\n[FAILED] No matching password found.")
	}
}

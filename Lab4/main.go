package main

import (
	"Lab4/utils/crack"
	"fmt"
)

func main() {

	targetHash := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"

	wordlist := "nord_vpn.txt"

	fmt.Println("================================")
	fmt.Println("      SHA512 PASSWORD CRACKER")
	fmt.Println("================================")

	fmt.Println("[*] Target Hash:", targetHash)
	fmt.Println("[*] Wordlist:", wordlist)
	fmt.Println("[*] Starting password cracking...")
	fmt.Println()

	password, err := crack.CrackSHA512(targetHash, wordlist)

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

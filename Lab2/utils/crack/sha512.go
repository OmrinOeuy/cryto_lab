package crack

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"os"
	"strings"
)

func CrackSHA1(targetHash, wordlistPath string) (string, error) {

	file, err := os.Open(wordlistPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {

		password := strings.TrimSpace(scanner.Text())

		if password == "" {
			continue
		}

		count++

		hash := sha1.Sum([]byte(password))

		hashString := fmt.Sprintf("%x", hash)

		fmt.Printf("[%d] Testing: %s | SHA1: %s\n",
			count, password, hashString)

		if strings.EqualFold(hashString, targetHash) {
			fmt.Println("\n[+] PASSWORD FOUND!")
			fmt.Println("[+] Password:", password)

			return password, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	fmt.Println("\n[-] Password not found in wordlist.")

	return "", nil
}

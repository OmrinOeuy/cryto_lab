package crack

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func CrackMD5(wordlistPath string, targetHash string, output *os.File) error {
	file, err := os.Open(wordlistPath)
	if err != nil {
		return fmt.Errorf("Cannot open Wordlist: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Buffer(make([]byte, 4096), 1024*1024)

	fmt.Fprintln(output, "========== MD5 PASSWORD CRACKER ==========")
	fmt.Fprintf(output, "Target Hash: %s\n\n", targetHash)

	count := 0

	for scanner.Scan() {

		password := strings.TrimSuffix(scanner.Text(), "\r")
		count++

		hash := md5.Sum([]byte(password))
		hashString := hex.EncodeToString(hash[:])

		fmt.Fprintf(output, "[%d] Testing: %s | MD5: %s\n", count, password, hashString)

		if hashString == targetHash {

			fmt.Fprintln(output, "\n========== PASSWORD FOUND ==========")
			fmt.Fprintf(output, "Password: %s\n", password)
			fmt.Fprintf(output, "MD5: %s\n", hashString)
			fmt.Fprintf(output, "Total attempts: %d\n", count)

			return nil
		}

	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading wordlist: %w", err)
	}

	fmt.Fprintln(output, "\nPassword not found in wordlist.")
	fmt.Fprintf(output, "Total attempts: %d\n", count)

	return nil
}

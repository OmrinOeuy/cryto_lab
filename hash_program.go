package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/sha3"
)

func calNCom(txt1, txt2 string) {
	// MD5
	md5A := md5.Sum([]byte(txt1))
	md5B := md5.Sum([]byte(txt2))

	// SHA1
	sha1A := sha1.Sum([]byte(txt1))
	sha1B := sha1.Sum([]byte(txt2))

	// SHA256
	sha256A := sha256.Sum256([]byte(txt1))
	sha256B := sha256.Sum256([]byte(txt2))

	// SHA512
	sha512A := sha512.Sum512([]byte(txt1))
	sha512B := sha512.Sum512([]byte(txt2))

	// SHA3-256
	sha3A := sha3.Sum256([]byte(txt1))
	sha3B := sha3.Sum256([]byte(txt2))

	fmt.Println("\n======== Hash Results ========")

	// Display MD5
	fmt.Println("\nHash (MD5):")
	fmt.Printf("Output A = %x\n", md5A)
	fmt.Printf("Output B = %x\n", md5B)
	fmt.Printf("=> %s\n", compare(md5A, md5B))

	// Display SHA1
	fmt.Println("\nHash (SHA1):")
	fmt.Printf("Output A = %x\n", sha1A)
	fmt.Printf("Output B = %x\n", sha1B)
	fmt.Printf("=> %s\n", compare(sha1A, sha1B))

	// Display SHA256
	fmt.Println("\nHash (SHA256):")
	fmt.Printf("Output A = %x\n", sha256A)
	fmt.Printf("Output B = %x\n", sha256B)
	fmt.Printf("=> %s\n", compare(sha256A, sha256B))

	// Display SHA512
	fmt.Println("\nHash (SHA512):")
	fmt.Printf("Output A = %x\n", sha512A)
	fmt.Printf("Output B = %x\n", sha512B)
	fmt.Printf("=> %s\n", compare(sha512A, sha512B))

	// Display SHA3-256
	fmt.Println("\nHash (SHA3-256):")
	fmt.Printf("Output A = %x\n", sha3A)
	fmt.Printf("Output B = %x\n", sha3B)
	fmt.Printf("=> %s\n", compare(sha3A, sha3B))
}

func compare(a, b interface{}) string {
	if fmt.Sprintf("%x", a) == fmt.Sprintf("%x", b) {
		return "Match!"
	}
	return "No Match!"
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("======= Name + Hashing Program =======")

	fmt.Print("Please input value 1: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)

	fmt.Print("Please input value 2: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)

	calNCom(input1, input2)

}

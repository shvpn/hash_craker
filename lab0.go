package main

// All hash program and compared the two hashes of two inputs
import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func main() {
	var input1, input2 string

	fmt.Print("Enter first string: ")
	fmt.Scanln(&input1)
	fmt.Print("Enter second string: ")
	fmt.Scanln(&input2)
	// Sample inputs
	a := input1
	b := input2

	// hasting using SHA256
	var hash1, hash2 string
	hash1 = s256(a)
	hash2 = s256(b)
	//compare the two hashes
	println("===================================")
	println("SHA256 Hash 1:", hash1)
	println("SHA256 Hash 2:", hash2)
	if hash1 == hash2 {
		fmt.Println("The hashes match.")
	} else {
		fmt.Println("The hashes do not match.")
	}
	// hasting using SHA512
	hash1 = s512(a)
	hash2 = s512(b)
	//compare the two hashes
	println("===================================")
	println("SHA512 Hash 1:", hash1)
	println("SHA512 Hash 2:", hash2)
	if hash1 == hash2 {
		fmt.Println("The hashes match.")
	} else {
		fmt.Println("The hashes do not match.")
	}
	// hasting using SHA1
	hash1 = s1(a)
	hash2 = s1(b)
	//compare the two hashes
	println("===================================")
	println("SHA1 Hash 1:", hash1)
	println("SHA1 Hash 2:", hash2)
	if hash1 == hash2 {
		fmt.Println("The hashes match.")
	} else {
		fmt.Println("The hashes do not match.")
	}

	// hasting using SHA3
	hash1 = s3(a)
	hash2 = s3(b)
	//compare the two hashes
	println("====================================")
	println("SHA3 Hash 1:", hash1)
	println("SHA3 Hash 2:", hash2)
	if hash1 == hash2 {
		fmt.Println("The hashes match.")
	} else {
		fmt.Println("The hashes do not match.")
	}
	// hasting using MD5
	hash1 = computeMD5(a)
	hash2 = computeMD5(b)
	//compare the two hashes
	println("====================================")
	println("MD5 Hash 1:", hash1)
	println("MD5 Hash 2:", hash2)
	if hash1 == hash2 {
		fmt.Println("The hashes match.")
	} else {
		fmt.Println("The hashes do not match.")
	}

}

// computeHash computes the SHA256 hash of a given input string.
func s256(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}
func s512(input string) string {
	hash := sha512.Sum512([]byte(input))
	return hex.EncodeToString(hash[:])
}
func s1(input string) string {
	hash := sha1.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func s3(input string) string {
	hash := sha3.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}
func computeMD5(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

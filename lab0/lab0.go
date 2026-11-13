package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	

	"golang.org/x/crypto/sha3"
)
func main() {
	var input1 string
	var input2 string



	fmt.Print("enter input1: ")
	fmt.Scanln(&input1)

	fmt.Print("enter Input2: ")
	fmt.Scan(&input2)

	fmt.Println("input1:", input1)
	fmt.Println("input2:", input2)
	
	hash1 := md5.Sum([]byte(input1))
	hash2 := md5.Sum([]byte(input2))
	
	fmt.Println("MD5 Input1:", hash1)
 	fmt.Println("MD5 Input2:", hash2)
	

	if hash1 == hash2 {
		fmt.Println("Match")
	}else {
		fmt.Println("no match ")
	}
	




	sha1Hash1 := sha1.Sum([]byte(input1))
	sha1Hash2 := sha1.Sum([]byte(input2))
	
	
	fmt.Println("sha1 input1:",sha1Hash1)
	fmt.Println("sha1 input2:",sha1Hash2)

	if sha1Hash1 == sha1Hash2 {
		fmt.Println("Match")
	}else {
		fmt.Println("no match")
	}
	sha256Hash1 := sha256.Sum256([]byte(input1))
	sha256Hash2 := sha256.Sum256([]byte(input2))

	fmt.Println("sha256 input1:",sha256Hash1)
	fmt.Println("sha256 input2:",sha256Hash2)
	if sha256Hash1 == sha256Hash2 {
		fmt.Println("Match")
	}else {
		fmt.Println("no match")
	}
	sha512Hash1 := sha512.Sum512([]byte(input1))
	sha512Hash2 := sha512.Sum512([]byte(input2))

	fmt.Println("sha512 input1:",sha512Hash1)
	fmt.Println("sha512 input2:",sha512Hash2)
	if sha512Hash1 == sha512Hash2 {
		fmt.Println("Match")
	}else {
		fmt.Println("no match")
	}
	sha3Hash1 := sha3.Sum256([]byte(input1))
	sha3Hash2 := sha3.Sum256([]byte(input2))

	fmt.Println("sha3 input1:", sha3Hash1)
	fmt.Println("sha3 input2:", sha3Hash2)

	if sha3Hash1 == sha3Hash2 {
    fmt.Println("Match")
} else {
    fmt.Println("no match")
}




}
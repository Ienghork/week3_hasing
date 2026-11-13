package main

import (
	"bufio"
	"fmt"
	
	"os"

	"lab2/utils/crack"

)



func main () {


   	targetHash := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"
	
	
	file, err := os.Open("wordlists/nord_vpn.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		password := scanner.Text()
		result :=crack.Sha1Hash(password)


		fmt.Println("trying password:", password)


		if result == targetHash {
			fmt.Println("passeword found:", password)

			break
		}
	}
}

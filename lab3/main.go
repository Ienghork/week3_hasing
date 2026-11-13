package main

import (
    "fmt"
    "lab3/utils/crack"
	"os"
	"bufio"
)

func main() {
   targetHash := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"

   file, err := os.Open("wordlists/nord_vpn.txt")



   if err !=nil {
	fmt.Println("error opening file:", err)
	return
   }
   defer file.Close()

   scanner := bufio.NewScanner(file)


   for scanner.Scan() {
	password := scanner.Text()
	
	result := crack.Sha512Hash(password)
	fmt.Println("Trying password:", password)
	if result == targetHash {
		fmt.Println("Password found:", password)
		break
	}
   }
		if err := scanner.Err(); err != nil {
    fmt.Println("Error reading file:", err)
	}
}
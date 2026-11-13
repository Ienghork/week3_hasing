
package main

import (
    "bufio"
    "fmt"
    "os"

    "lab1/utils/crack"
)

func main() {
    targetHash := "6a85dfd77d9cb35770c9dc6728d73d3f"

    // Open the wordlist file
    file, err := os.Open("wordlists/nord_vpn.txt")

    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }

    defer file.Close()

    // Create a scanner to read the file line by line
    scanner := bufio.NewScanner(file)

    found := false

    // Read every line in the wordlist
    for scanner.Scan() {
        password := scanner.Text()

        // Generate MD5 hash for the current password
        result := crack.MD5Hash(password)

        // Display the password being tested
        fmt.Println("Trying password:", password)

        // Compare generated hash with target hash
        if result == targetHash {
            fmt.Println("Password found:", password)
            found = true
            break
        }
    }

    // Check for errors while reading the file
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }

    if !found {
        fmt.Println("Password not found in wordlist.")
    }
}


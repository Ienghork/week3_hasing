package crack

import (
	"crypto/sha512"
	"fmt"
	
)



func Sha512Hash(password string) string {
	hash := sha512.Sum512([]byte(password))


	result := fmt.Sprintf("%x", hash)
	return result
}




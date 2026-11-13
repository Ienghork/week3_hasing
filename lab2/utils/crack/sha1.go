package crack

import (
	"crypto/sha1"
	"fmt"
	
)



func Sha1Hash(password string) string {
	hash := sha1.Sum([]byte(password))


	result := fmt.Sprintf("%x", hash)
	return result
}




package crack

import (
    "crypto/md5"
    "fmt"
)

func MD5Hash(password string) string {
hash := md5.Sum([]byte(password))
return fmt.Sprintf("%x", hash)
}
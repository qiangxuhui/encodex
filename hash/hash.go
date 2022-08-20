package hash

import "crypto/md5"
import "fmt"

func Md5(str string) string {
	return fmt.Sprintf("%x\n", md5.New().Sum([]byte(str)))
}

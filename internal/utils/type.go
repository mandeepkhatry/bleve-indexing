package utils

import "fmt"

//FindType returns type of interface
func FindType(data interface{}) string {
	return fmt.Sprintf("%T", data)
}

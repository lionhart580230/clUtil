package clFile

import (
	"fmt"
	"testing"
)

func TestGetFileMD5(t *testing.T) {

	md51 := GetFileMD5("123.png")
	md52 := GetFileMD5("321.png")

	fmt.Printf("md51: %v\n", md51)
	fmt.Printf("md52: %v\n", md52)
}
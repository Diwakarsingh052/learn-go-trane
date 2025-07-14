package stringops

import "fmt"

func TrimAndUpper(s string) string {
	fmt.Println(trim(s))
	fmt.Println(toUpper(s))
	return "success"
}

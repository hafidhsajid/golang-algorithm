package function

import "fmt"

func ReverseString(s string) string {
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])

	}
	fmt.Println(res)
	return res
}

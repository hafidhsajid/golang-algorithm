package function

import (
	"fmt"
	"strings"
)

func FindDuplicate(input string) {

	inputSlice := strings.Replace(input, " ", "", -1)
	inputSlice = strings.ToLower(inputSlice)
	var res string
	for i := 0; i < len(inputSlice); i++ {
		var count = 1
		for j := i + 1; j < len(inputSlice); j++ {
			if inputSlice[i] == inputSlice[j] {
				count += 1
				inputSlice = inputSlice[:j] + inputSlice[j+1:]
			}
		}
		if count > 1 {
			inputSlice = inputSlice[:i] + fmt.Sprint(count) + inputSlice[i:]
			res = string(inputSlice)
		}
	}
	fmt.Println(res)
}

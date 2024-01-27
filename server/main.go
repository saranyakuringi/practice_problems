package main

import (
	"fmt"
)

func main() {
	str1 := "ABCABC"
	str2 := "ABC"
	len1 := len
	p := fmt.Println
	p(len(str1))
	p(len(str2))
	var i, j = 0, 0
	var result string
	for i < len(str1) && j < len(str2) {
		if str1[i]%str2[j] == 0 {
			//result = string(str1[i] + str2[j])
			result = result + string(str1)
			i++
			result = result + string(str2)
			j++
			//p(string(str1[i]), string(str2[j]))
		}
	}

	p(result)

}

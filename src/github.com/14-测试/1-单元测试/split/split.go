package split

import (
	"fmt"
	"strings"
)

// Split 切割字符串
// example:
// abc , b => [a c]

func Split(s string, sep string) (result []string) {
	if sep == "" {
		for _,v := range s {
			result = append(result,string(v))
		}
		return
	}
	i := strings.Index(s, sep)
	fmt.Println(i)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+1:]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

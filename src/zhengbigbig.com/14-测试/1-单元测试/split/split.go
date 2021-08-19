package split

import (
	"strings"
)

// Split 切割字符串
// example:
// abc , b => [a c]

func Split(s string, sep string) (result []string) {
	result = make([]string, 0, strings.Count(s, sep)+1) // 提前分配内存，避免后续内存分配导致性能下降
	if sep == "" {
		for _, v := range s {
			result = append(result, string(v))
		}
		return
	}
	i := strings.Index(s, sep)
	for i >= 0 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

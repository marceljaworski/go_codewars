import (
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) string {
	result := ""
	s := strings.Split(s1+s2, "")
	sort.Strings(s)
	for _, v := range s {
		if !strings.Contains(result, v) {
			result += v
		}
	}
	return result
}
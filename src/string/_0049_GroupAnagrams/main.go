package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	result := [][]string{}
	m := map[string][]string{}
	for _, str := range strs {
		tmp := []string{}
		for _, r := range str {
			tmp = append(tmp, string(r))
		}
		sort.Strings(tmp)
		key := strings.Join(tmp, "")
		m[key] = append(m[key], str)
	}
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

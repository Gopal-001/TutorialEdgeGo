package main

import "fmt"

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []string {
	mp := make(map[string]bool)
	ans := []string{}
	for idx := range developers {
		if _, ok := mp[developers[idx].Name]; !ok {
			ans = append(ans, developers[idx].Name)
			mp[developers[idx].Name] = true
		}
	}
	return ans
}

func main() {
	fmt.Println("Filter Unique Challenge")
}

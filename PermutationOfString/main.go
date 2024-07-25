package main

import "fmt"

func swap(a *rune, b *rune) {
	*a, *b = *b, *a
}
func checkAllPerm(i int, n int, s []rune, t *string) bool {
	if i == n-1 {
		return string(s) == *t
	}
	ans := checkAllPerm(i+1, n, s, t)
	j := i + 1
	for j < n {
		swap(&s[i], &s[j])
		ans = ans || checkAllPerm(i+1, n, s, t)
		swap(&s[i], &s[j])
		j++
	}
	return ans
}

func CheckPermutations(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	runeStr := []rune(str1)
	isPerm := checkAllPerm(0, len(str1), runeStr, &str2)
	return isPerm
}

func main() {
	fmt.Println("Check Permutations Challenge")

	str1 := "adme"
	str2 := "meda"

	isPermutation := CheckPermutations(str1, str2)
	fmt.Println(isPermutation)

}

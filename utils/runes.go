package utils

//rune slice type and sorting functions for sort.Sort interface
type Runes []rune

func (s Runes) Len() int {
	return len(s)
}

func (s Runes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Runes) Less(i, j int) bool {
	return s[i] < s[j]
}
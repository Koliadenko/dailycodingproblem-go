package day57

import "strings"

// BreakByLength takes a long string and then breaks it into full-word
// strings that are no longer than k characters.
func BreakByLength(long string, k int) []string {
	var result []string
	words := strings.Split(long, " ")
	var line []string
	ll := 0
	for i := 0; i < len(words); i++ {
		if ll == 0 && len(words[i]) <= k {
			line = append(line, words[i])
			ll = len(words[i])
		} else if ll+len(words[i])+1 <= k {
			line = append(line, words[i])
			ll += len(words[i]) + 1
		} else if ll > 0 {
			result = append(result, strings.Join(line, " "))
			line = line[:0]
			ll = 0
			i--
		} else {
			return nil
		}
	}
	if ll > 0 {
		result = append(result, strings.Join(line, " "))
	}
	return result
}

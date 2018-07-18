package daily

// This problem was asked by Microsoft.
// Given a dictionary of words and a string made up of those words (no spaces), return the original sentence in a list.
// If there is more than one possible reconstruction, return any of them. If there is no possible reconstruction, then return null.
// Case A: For example, given the set of words 'quick', 'brown', 'the', 'fox', and the string "thequickbrownfox", you should return ['the', 'quick', 'brown', 'fox'].
// Case B: Given the set of words 'bed', 'bath', 'bedbath', 'and', 'beyond', and the string "bedbathandbeyond", return either ['bed', 'bath', 'and', 'beyond] or ['bedbath', 'and', 'beyond'].

// Case C: set of words 'bed', 'bedbath', 'and', 'beyond', and string "bedbathandbeyond", return ['bedbath', 'and', 'beyond']

// toList returns the list of words from the input string which exist in the provided map of words.
// This solution does not solve Case C.
//
// runtime: O(n)
// space:   O(n) because of the output array
func toList(in string, words map[string]bool) []string {
	p1 := 0
	out := []string{}

	n := len(in)
	if n == 0 || len(words) == 0 {
		return nil
	}

	for p2 := 1; p2 <= n; p2++ {
		word := in[p1:p2]
		if words[word] {
			out = append(out, word)
			p1 = p2 + 1
		}
	}

	return out
}

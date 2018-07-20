package daily

import "fmt"

// This problem was asked by Facebook.

// Implement regular expression matching with the following special characters:

// . (period) which matches any single character
// * (asterisk) which matches zero or more of the preceding element
// That is, implement a function that takes in a string and a valid regular expression and returns whether or not the string matches the regular expression.

// For example, given the regular expression "ra." and the string "ray", your function should return true.
// The same regular expression on the string "raymond" should return false.

// Given the regular expression ".*at" and the string "chat", your function should return true.
// The same regular expression on the string "chats" should return false.

func matchPattern(pattern string, s string) (bool, error) {
	m, n := len(pattern), len(s)

	p2 := 0
	for p1 := 0; p1 < m; p1++ {
		// reached the end of both pattern and s
		if p1 == m && p2 == n {
			return true, nil
		}

		// reached the end of either pattern or s and not the other
		if p1 == m && p2 != n || p1 != m && p2 == n {
			return false, nil
		}

		if pattern[p1] == s[p2] {
			p2++
			continue
		}

		if p1+1 < m && pattern[p1+1] == '*' {
			if pattern[p1] == '*' {
				return false, fmt.Errorf("error: can't have '**' ")
			}

			if pattern[p1] == '.' {
				// .* is last characters in pattern, therefore match all chars in rest of string
				if p1+2 >= m {
					return true, nil
				}

				// .* is followed by another char, therefore move over until we find a match
				for s[p2] != pattern[p1+2] {
					p2++
				}

				p1++
				continue
			}

			// x* so match x until there is no more matches
			for s[p2] == pattern[p1] {
				p2++
			}

			p1++
			continue
		}

		if pattern[p1] == '.' {
			// match any char
			p2++
		}
	}

	// both are empty
	return true, nil
}

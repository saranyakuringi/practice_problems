//Merge string Alternately

//You are given two strings word1 and word2.
//Merge the strings by adding letters in alternating order, starting with word1.
//If a string is longer than the other, append the additional letters onto the end of the merged string.
//Return the merged string.

// Input: word1 = "abc", word2 = "pqr"
// Output: "apbqcr"
// Explanation: The merged string will be merged as so:
// word1:  a   b   c
// word2:    p   q   r
// merged: a p b q c r

package exc13

func MergeStrings(word1, word2 string) string {
	var result string
	var i, j = 0, 0
	for i < len(word1) && j < len(word2) {
		result = result + string(word1[i])
		i++
		result = result + string(word2[j])
		j++
	}
	result = result + word1[i:]
	result = result + word2[j:]
	return result
}

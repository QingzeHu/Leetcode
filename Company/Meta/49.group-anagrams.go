/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	anagram := map[string][]int{}
	for i := 0; i < len(strs); i++ {
		letters := make([]byte, 256)
		for j := 0; j < len(strs[i]); j++ {
			letters[strs[i][j]] += 1
		}
		s := string(letters)
		anagram[s] = append(anagram[s], i)
	}
	result := make([][]string, 0, len(anagram))
	for _, indexes := range anagram {
		str := make([]string, 0, len(indexes))
		for _, id := range indexes {
			str = append(str, strs[id])
		}
		result = append(result, str)
	}
	return result
}

// @lc code=end


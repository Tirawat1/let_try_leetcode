func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	ans := []rune{}
	for i := 0; i < len(strs[0]); i++ {
		chkFirstWord := strs[0][i]
		same := false
		// loop check another word
		for j := 1; j < len(strs); j++ {
			// check ตำแหน่งของ char ใน ทุก word
			if i > len(strs[j])-1 || chkFirstWord != strs[j][i] {
				same = false
				break
			}
			same = true
		}
		if same {
			ans = append(ans, rune(chkFirstWord))
		} else {
			break
		}
	}
	return string(ans)
}
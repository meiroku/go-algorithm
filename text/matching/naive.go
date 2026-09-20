package matching

// 素朴に文字を照合
func Naive(s, w string) int {
	// バイト単位ではなくUnicodeコードポイント単位で扱う
	rs, rw := []rune(s), []rune(w)
	lenS, lenW := len(rs), len(rw)

	if lenW == 0 {
		return 0
	}

	for i := 0; i <= lenS-lenW; i++ {
		j := 0

		// 一文字ずつ一致を確認
		for j < lenW && rs[i+j] == rw[j] {
			j++
		}
		if j == lenW {
			return i // 開始位置のインデックス
		}
	}

	return -1
}
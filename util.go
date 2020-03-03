package tony

func Substring(str string, indexs ...int) string {
	if len(indexs) == 0 {
		return str
	}

	strLen := len(str)
	if len(indexs) == 1 {
		if indexs[0] > strLen {
			return ""
		}
		return string([]rune(str)[indexs[0]:])
	}

	if indexs[0] >= indexs[1] {
		return ""
	}

	if indexs[1] > strLen {
		indexs[1] = strLen
	}
	return string([]rune(str)[indexs[0]:indexs[1]])
}

package service

func FindMaxSubstring(text string) string {
	// two pointers algorithm
	if len(text) <= 1 {
		return text
	}
	left, right := 0, 0
	maxSubString := ""
	charIndexMap := make(map[byte]int)

	for right < len(text) {
		// Если текущий символ уже встречался в подстроке, обновляем левый указатель
		if index, ok := charIndexMap[text[right]]; ok && index >= left {
			left = index + 1
		}
		// Обновляем максимальную подстроку при необходимости
		subString := text[left : right+1]
		if len(subString) > len(maxSubString) {
			maxSubString = subString
		}
		// Обновляем правый указатель и map
		right++
		charIndexMap[text[right-1]] = right - 1
	}
	return maxSubString
}

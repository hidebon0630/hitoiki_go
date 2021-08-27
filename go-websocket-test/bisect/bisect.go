package bisect

// 仮実装、あっているかどうかわからない
func BisectLeft(array []int, key int) int {
	var left, right, middle int
	left = 0
	right = len(array)
	for left < right {
		middle = ((right - left) / 2) + left
		if key <= array[middle] {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}

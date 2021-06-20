package tree

import (
	"errors"
	"sort"
)

func BinarySearch(array sort.IntSlice, target int) (index int, err error) {
	// 排序
	array.Sort()

	left := 0
	right := array.Len() - 1
	mid := right / 2
	for left <= right {
		// 目标值与当前值相等，退出循环
		if target == array[mid] {
			index = mid
			break
		}

		if target > array[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = (left + right) / 2
	}
	if right < left {
		err = errors.New("error: target value can not be found")
	}
	return
}

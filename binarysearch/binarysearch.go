package binarysearch

//查找对应target位置（返回对应位置，没找到则返回-1)
func BinarySearch1(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

//查找第一个大于或等于target的位置（返回对应位置，没找到则返回-1）
func BinarySearch2(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if l <= len(nums)-1 {
		return l
	}
	return -1
}

//查找第一个大于target的位置（返回对应位置，没找到则返回-1）
func BinarySearch3(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if l <= len(nums)-1 {
		return l
	}
	return -1
}

//查找最后一个小于等于target的位置（返回对应位置，没找到则返回-1）
func BinarySearch4(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if r >= 0 {
		return r
	}
	return -1
}

//查找最后一个小于target的位置（返回对应位置，没找到则返回-1）
func BinarySearch5(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if r >= 0 {
		return r
	}
	return -1
}

//查找第一个等于target的位置（返回对应位置，没找到则返回-1）
func BinarySearch6(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l <= len(nums) && nums[l] == target {
		return l
	}
	return -1
}

//查找最后一个等于target的位置（返回对应位置，没找到则返回-1）
func BinarySearch7(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if r >= 0 && nums[r] == target {
		return r
	}
	return -1
}

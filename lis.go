package goplay

func LIS(a []int) int {
	n := len(a)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	ans := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if a[j] < a[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func LISBinarySearch(a []int) int {
	n := len(a)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = a[0]

	for i := 1; i < n; i++ {
		if a[i] > dp[len(dp)-1] {
			// only compare to the last one and append if it's larger
			dp = append(dp, a[i])
		} else {
			idx := binarySearch(dp, a[i])
			if idx == 0 || idx == len(dp) { // if at the ends, replace the ends
				dp[idx] = a[i]
				continue
			} else {
				dp[idx-1] = a[i] // if in the middle, replace the one before
			}
		}
	}

	return len(dp)
}

// if target does not exist, it will return the index where it should be inserted
func binarySearch(a []int, target int) int {
	l, r := 0, len(a)-1
	for l < r {
		mid := (l + r) / 2
		if a[mid] == target {
			return mid
		}

		if a[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

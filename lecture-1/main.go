package lecture_1

/*
	Problen:
	Find the sum of the first N numbers

    Objective Function:
    f(i) is the sum of the first i elements

    Recurrence relation:
	f(n) = f(n-1) + n
*/
// n = 0
func nSum(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + i
	}
	return dp[n]
}

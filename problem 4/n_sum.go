func sum_to_n_a(n int) int {
	// AP sum formula
	// Time complexity: O(1), Fixed number of computations, not dependent on n.
    return (n * (n + 1)) / 2
}

func sum_to_n_b(n int) int {
	// Recursion
	// Time complexity: O(n), n recursive calls
	if n == 0 {
		return 0
	}

	return n + sum_to_n_b(n - 1)

}

func sum_to_n_c(n int) int {
	// For loop
	// Time complexity: O(n), n loops
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

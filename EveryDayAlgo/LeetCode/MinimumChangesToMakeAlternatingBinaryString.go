package main

func minOperationsBinaryString(s string) int {
	n := len(s)
	costA, costB := 0, 0

	for i := 0; i < n; i++ {
		expectedA := byte('0')
		expectedB := byte('1')

		if i%2 == 1 {
			expectedA = '1' // Target A: "0101..."
			expectedB = '0' // Target B: "1010..."
		}

		if s[i] != expectedA {
			costA++
		}
		if s[i] != expectedB {
			costB++
		}
	}

	if costA < costB {
		return costA
	}
	return costB
}

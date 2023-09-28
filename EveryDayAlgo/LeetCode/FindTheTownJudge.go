package main

func findJudge(n int, trust [][]int) int {
	if len(trust) < n-1 {
		return -1
	}
	indegree := make([]int, n+1)
	outdegree := make([]int, n+1)

	for i := 0; i < len(trust); i++ {
		outdegree[trust[i][0]] += 1
		indegree[trust[i][1]] += 1
	}

	for i := 1; i <= n; i++ {
		if outdegree[i] == 0 && indegree[i] == n-1 {
			return i
		}
	}

	return -1
}

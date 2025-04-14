package main

func countStudents(students []int, sandwiches []int) int {
	cnt := 0
	for len(students) != 0 {
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
			cnt = 0
		} else {
			cnt++
			students = append(students[1:], students[0])
		}
		if cnt == len(sandwiches) {
			return len(sandwiches)
		}
	}
	return 0
}

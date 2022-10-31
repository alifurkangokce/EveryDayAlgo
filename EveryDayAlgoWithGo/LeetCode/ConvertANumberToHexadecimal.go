package main

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num = 4294967295 + num + 1
	}
	m := map[int]string{
		0:  "0",
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "a",
		11: "b",
		12: "c",
		13: "d",
		14: "e",
		15: "f"}
	var a string
	r := 0
	for {
		if num < 16 {
			a = m[num] + a
			break
		}
		r = num % 16
		a = m[r] + a
		num = (num - r) / 16
	}
	return a
}

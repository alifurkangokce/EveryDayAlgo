package main

var months = map[int]int{
	0:  0,
	1:  31,
	2:  59,
	3:  90,
	4:  120,
	5:  151,
	6:  181,
	7:  212,
	8:  243,
	9:  273,
	10: 304,
	11: 334,
}

var weekDays = map[int]string{
	0: "Monday",
	1: "Tuesday",
	2: "Wednesday",
	3: "Thursday",
	4: "Friday",
	5: "Saturday",
	6: "Sunday",
}

// 1st January 0001 - Saturday
var firstDay = 5

func dayOfTheWeek(day int, month int, year int) string {
	checkLeapYear := year
	// do not inlude this year
	// if we don't pass February
	if month <= 2 {
		checkLeapYear -= 1
	}

	leapYears := (checkLeapYear / 4) - (checkLeapYear / 100) + (checkLeapYear / 400)

	monthDays := months[month-1]

	// we count years from 1st not from 0th
	days := 365 * (year - 1)
	days += monthDays
	days += day
	days += leapYears
	days += 1

	return weekDays[(days+firstDay)%7]
}

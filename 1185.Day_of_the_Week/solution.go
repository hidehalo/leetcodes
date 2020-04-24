package main

func dayOfTheWeek(day int, month int, year int) string {
	pastDays := 0
	daysOfMonth := []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	// 1971/01/01 is "Friday"
	daysText := []string{"Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday"}
	for i := 1971; i < year; i++ {
		pastDays += isLeapYear(i) + 365
	}
	pastDays += daysOfMonth[month-1] + day
	if month > 2 {
		pastDays += isLeapYear(year)
	}

	return daysText[(pastDays-1)%7]
}

func isLeapYear(year int) int {
	if (year%4 == 0 && year%100 != 0) || (year%100 == 0 && year%400 == 0) {
		return 1
	}

	return 0
}

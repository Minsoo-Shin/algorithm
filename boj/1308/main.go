package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var nowY, nowM, nowD int
	fmt.Fscanln(reader, &nowY, &nowM, &nowD)

	var y, m, d int
	fmt.Fscanln(reader, &y, &m, &d)
	fmt.Printf("%v\n", solution(nowY, nowM, nowD, y, m, d))
	return
}

func solution(nowY, nowM, nowD, y, m, d int) string {
	if y-nowY > 1000 ||
		(y-nowY == 1000 && m > nowM) ||
		(y-nowY == 1000 && nowM == m && d >= nowD) {
		return "gg"
	}
	days := CountDays(y, m, d) - CountDays(nowY, nowM, nowD)
	if nowY != y {
		for ; nowY < y; nowY++ {
			days += GetDaysOfYear(nowY)
		}
	}
	return fmt.Sprintf("D-%d", days)
}

func IsLeapYear(year int) bool {
	var isLeapYear bool
	if year%4 == 0 {
		isLeapYear = true
	}

	if year%100 == 0 {
		isLeapYear = false
	}

	if year%400 == 0 {
		isLeapYear = true
	}

	return isLeapYear
}

func CountDays(year, month, day int) int {
	days := 0

	for cur := 1; cur < month; cur++ {
		days += GetDaysOfMonth(year, cur)
	}
	days += day

	return days
}

func GetDaysOfMonth(year, month int) int {
	switch month {
	case 2:
		if IsLeapYear(year) {
			return 29
		}
		return 28
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	default:
		return 30
	}
}

func GetDaysOfYear(year int) int {
	if IsLeapYear(year) {
		return 366
	}
	return 365
}

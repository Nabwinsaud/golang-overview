package main

import "strconv"

func Split(sum int) (x string, y int) {
	sum = sum * 2 / 10
	x = strconv.Itoa(sum)
	y = sum - 10
	return
}

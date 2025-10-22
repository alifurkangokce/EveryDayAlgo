package main

import "strings"

func squareIsWhite(coordinates string) bool {
	var str = "aceg"
	if strings.Contains(str, string(coordinates[0])) {
		if int(coordinates[1])%2 == 0 {
			return true
		} else {
			return false
		}
	} else {
		if int(coordinates[1])%2 == 0 {
			return false
		} else {
			return true
		}
	}
}

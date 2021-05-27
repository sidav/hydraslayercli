package main

var Reset  = "\033[0m "
var Red    = "\033[31m"
var Green  = "\033[32m"
var Yellow = "\033[33m"
var Blue   = "\033[34m"
var Purple = "\033[35m"
var Cyan   = "\033[36m"
var Gray   = "\033[37m"
var White  = "\033[97m"

func colorizeString(color string, str string) string {
	return color + str + Reset
}

func colorizeStringByArray(colors []string, str string, singleColorLength int) string {
	totalColors := len(colors)
	if totalColors == 1 {
		return colorizeString(colors[0], str)
	}

	if singleColorLength == 0 {
		singleColorLength = len(str)/totalColors
	}
	coloredString := ""
	currColor := 0
	for i := 0; i < len(str); i+= singleColorLength {
		for j := i; j < i+singleColorLength; j++ {
			if j >= len(str) {
				break
			}
			if str[j] == "\033"[0] {
				i+=4
				continue
			}
		}
		if i +singleColorLength >= len(str) {
			coloredString += colors[currColor % len(colors)] + str[i:]
			break
		}
		coloredString += colors[currColor % len(colors)] + str[i : i+singleColorLength]
		currColor++
	}
	return coloredString + Reset
}

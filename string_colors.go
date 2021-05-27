package main

var Reset  = "\033[0m"
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

func colorizeStringByArray(colors []string, str string, colorizeCharwise bool) string {
	totalColors := len(colors)
	if totalColors == 1 {
		return colorizeString(colors[0], str)
	}
	singleColorChars := len(str)/totalColors
	coloredString := ""
	for i := 0; i < totalColors; i++ {
		if i < totalColors-1 {
			coloredString += colorizeString(colors[i], str[i*singleColorChars : (i+1)*singleColorChars])
		} else {
			coloredString += colorizeString(colors[i], str[i*singleColorChars : len(str)])
		}
	}
	return coloredString
}

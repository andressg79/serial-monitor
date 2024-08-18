package console

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

var Blood = ";1m"
var Underline = "\033[4m"
var Strike = "\033[9m"

func RedText(text string) string {
	return Red + text + Reset
}

func GreenText(text string) string {
	return Green + text + Reset
}

func YellowText(text string) string {
	return Yellow + text + Reset
}

func BlueText(text string) string {
	return Blue + text + Reset
}

func MagentaText(text string) string {
	return Magenta + text + Reset
}

func CyanText(text string) string {
	return Cyan + text + Reset
}

func GrayText(text string) string {
	return Gray + text + Reset
}

func WhiteText(text string) string {
	return White + text + Reset
}

func BloodText(text string) string {
	return Blood + text + Reset
}

func UnderlineText(text string) string {
	return Underline + text + Reset
}

func StrikeText(text string) string {
	return Strike + text + Reset
}

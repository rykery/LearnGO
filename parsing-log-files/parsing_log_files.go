package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	reg := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return reg.MatchString(text)
}

func SplitLogLine(text string) []string {
	reg := regexp.MustCompile(`<(~|\*|=|-)*>`)
	return reg.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	reg := regexp.MustCompile(`"`)
	cnt := 0
	for _, l := range lines {
		if reg.MatchString(l) {
			cnt++
		}
	}
	return cnt
}

func RemoveEndOfLineText(text string) string {
	panic("Please implement the RemoveEndOfLineText function")
}

func TagWithUserName(lines []string) []string {
	panic("Please implement the TagWithUserName function")
}

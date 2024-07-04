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
	reg := regexp.MustCompile(`(?i)"(.*)password(.*)"`)
	cnt := 0
	for _, l := range lines {
		if reg.MatchString(l) {
			cnt++
		}
	}
	return cnt
}

func RemoveEndOfLineText(text string) string {
	reg := regexp.MustCompile(`end-of-line([0-9]*)`)
	return reg.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	reg := regexp.MustCompile(`User\s+(\S+)`)
	ret := make([]string, len(lines))
	for i, line := range lines {
		if reg.MatchString(line) {
			usrName := reg.FindStringSubmatch(line)[1]
			ret[i] = "[USR] " + usrName + " " + line
		} else {
			ret[i] = line
		}
	}
	return ret
}

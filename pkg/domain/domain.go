package domain

import "regexp"

var regex = regexp.MustCompile("^\\b[\\w+-]+[\\w.-]+\\.[a-zA-Z]{2,6}\\b")

func IsDomain(name string) bool {
	var isDomain = false
	result := regex.FindAllStringSubmatch(name, -1)
	if result != nil {
		isDomain = true
	}
	return isDomain
}
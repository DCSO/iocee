// DCSO IOCee IOC Extractor
// Copyright (c) 2017, DCSO GmbH

package data_types

import (
	"regexp"
)

var EMailRegex = regexp.MustCompile("(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$)")

func ParseEMail(line string) []string {
	result := EMailRegex.FindAllString(line, 100)
	if len(result) > 0 {
		return result
	}
	return []string{}
}

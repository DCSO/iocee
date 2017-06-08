// DCSO IOCee IOC Extractor
// Copyright (c) 2017, DCSO GmbH

package data_types

import (
	"regexp"
)

var URLRegex = regexp.MustCompile("(?i)(https?|ftp)://(-\\.)?([^\\s/?\\.#\"']+\\.?)+(/[^\\s\\\"']*)?")

func ParseURL(line string) []string {
	result := URLRegex.FindAllString(line, 100)
	if len(result) > 0 {
		return result
	}
	return []string{}
}

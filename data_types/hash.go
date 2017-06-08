// DCSO IOCee IOC Extractor
// Copyright (c) 2017, DCSO GmbH

package data_types

import "regexp"
import "fmt"

var HashRegex = regexp.MustCompile("(?i)(?:^|[^\\w])([\\w]{32}|[\\w]{40}|[\\w]{64})(?:[^\\w]|$)")

func ParseHash(line string) []string {
	result := HashRegex.FindAllString(line, 100)
	if len(result) > 0 {
		for _, s := range result {
			switch len(s) {
			case 32:
				result = append(result, fmt.Sprintf("md5:%s", s))
			case 40:
				result = append(result, fmt.Sprintf("sha1:%s", s))
			case 64:
				result = append(result, fmt.Sprintf("sha256:%s", s))
			}
		}
		return result
	}
	return []string{}
}

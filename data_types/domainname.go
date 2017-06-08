// DCSO IOCee IOC Extractor
// Copyright (c) 2017, DCSO GmbH

package data_types

import "strings"
import "regexp"
import "fmt"

var DomainNameRegex = regexp.MustCompile("(?i)([[:alpha:]\\-][\\w\\-]+\\.)+([[:alpha:]\\-][\\w-]+)")

func ParseDomainName(line string) []string {
	result := DomainNameRegex.FindAllString(line, 100)
	if len(result) > 0 {
		fullResult := make([]string, len(result), len(result))
		for _, s := range result {
			parts := strings.Split(s, ".")
			if len(parts) > 2 {
				for i := 1; i < len(parts)-1; i++ {
					joinedString := strings.Join(parts[i:len(parts)], ".")
					fullResult = append(fullResult, joinedString)
					fullResult = append(fullResult, fmt.Sprintf("*.%s", joinedString))
				}
			}
		}
		return fullResult
	}
	return []string{}
}

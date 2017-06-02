package data_types

import "strings"
import "regexp"

var DomainNameRegex = regexp.MustCompile("(?i)([[:alpha:]\\-][\\w\\-]+\\.)+([[:alpha:]\\-][\\w-]+)")

func ParseDomainName(line string) []string {
	result := DomainNameRegex.FindAllString(line, 100)
	if len(result) > 0 {
		for _, s := range result {
			parts := strings.Split(s, ".")
			if len(parts) > 2 {
				for i := 1; i < len(parts)-1; i++ {
					result = append(result, strings.Join(parts[i:len(parts)], "."))
				}
			}
		}
		return result
	}
	return []string{}
}

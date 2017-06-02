package data_types

import "fmt"
import "regexp"
import "strconv"

var IPv4Regex = regexp.MustCompile("(?i)([\\d]{1,3})\\.([\\d]{1,3})\\.([\\d]{1,3})\\.([\\d]{1,3})")

func ParseIPv4(line string) []string {
	result := IPv4Regex.FindAllString(line, 100)
	validatedResult := make([]string, 0, len(result))
Validate:
	for _, s := range result {
		matchAndNumbers := IPv4Regex.FindStringSubmatch(s)
		for i := 0; i < 4; i++ {
			n, _ := strconv.Atoi(matchAndNumbers[i+1])
			if n > 255 {
				continue Validate
			}
		}
		validatedResult = append(validatedResult, s)
		validatedResult = append(validatedResult, fmt.Sprintf("%s/32", s))
	}
	return validatedResult
}

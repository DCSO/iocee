// DCSO IOCee IOC Extractor
// Copyright (c) 2017, DCSO GmbH

package iocee

import "github.com/DCSO/iocee/data_types"

//A parser takes an input string and produces a list of output strings (0 or more)
type Parser func(string) []string

var parsers = map[string]Parser{
	"domainname": data_types.ParseDomainName,
	"ipv4":       data_types.ParseIPv4,
	"url":        data_types.ParseURL,
	"hash":       data_types.ParseHash,
	"email":      data_types.ParseEMail,
}

func Parse(line string) []string {
	results := make([]string, 0, 10)
	for _, parser := range parsers {
		result := parser(line)
		if len(result) > 0 {
			results = append(results, result...)
		}
	}
	return results
}

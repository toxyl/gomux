package main

import (
	"regexp"
)

var (
	reToken = regexp.MustCompile(`\{[^\}]+\}|\[[^\]]+\]|<[^>]+>`)
	cmdReg  = []*command{}
)

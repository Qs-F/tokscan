package tokscan

import "fmt"

var usages = []string{
	"text",
	"bg",
	"border",
	"fill",
}

var intentions = []string{
	"shade",
	"primary",
	"secondary",
	"positive",
	"negative",
	"notice",
	"informative",
}

var levels = []string{
	"dark",
	"medium",
	"light",
	"white",
}

var states = []string{
	"default",
	"hover",
	"active",
	"disabled",
}

func AllTokens() []string {
	ret := make([]string, 0, len(usages)*len(intentions)*len(levels)*len(states))
	for _, usage := range usages {
		for _, intention := range intentions {
			for _, level := range levels {
				for _, state := range states {
					ret = append(ret, fmt.Sprintf("%s-%s-%s-%s", usage, intention, level, state))
				}
			}
		}
	}
	return ret
}

func IsValidToken(token string) bool {
	for _, validToken := range AllTokens() {
		if token == validToken {
			return true
		}
	}
	return false
}

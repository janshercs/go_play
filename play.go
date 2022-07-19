package goplay

import "strings"

func bruteForceBinsRules(rules, bins []string) []string {
	subset := make([]string, 0, len(bins))
	for _, bin := range bins {
		for _, rule := range rules {
			if strings.HasPrefix(bin, rule) {
				subset = append(subset, bin)
			}
		}
	}
	return subset
}

var (
	mapSet  = false
	ruleSet = map[string]bool{}
)

// technically an O(n) + O(m) operation
func mapBinsRules(rules, bins []string) []string {
	subset := make([]string, 0, len(bins))
	initRules(rules)
	for _, bin := range bins {
		if fulfilsRules(bin) {
			subset = append(subset, bin)
		}
	}
	return subset
}

// O(m) operation
func initRules(rules []string) {
	if mapSet {
		return
	}
	for _, rule := range rules {
		ruleSet[rule] = true
	}
	mapSet = true
}

// Technically an O(1) operation
func fulfilsRules(bin string) bool {
	max := 6
	if len(bin) < 6 {
		max = len(bin) // setting length to be min(len(bin), 6)
	}

	var ok bool
	for i := 2; i < max+1; i++ { // O(5)
		if ok = ruleSet[bin[:i]]; !ok { // O(1)
			continue
		}
		return true
	}
	return false
}

type foo struct {
	Bar string `json:"bar"`
	Baz int    `json:"baz"`
}

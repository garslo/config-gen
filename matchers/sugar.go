package matchers

import "github.com/onsi/gomega/types"

// HaveSameLen will test if `item' has the same length as the match
// parameter.
func HaveSameLen(item interface{}) types.GomegaMatcher {
	return HaveSameLenMatcher{
		Item: item,
	}
}

package matchers

import "github.com/onsi/gomega/types"

func HaveSameLen(item interface{}) types.GomegaMatcher {
	return HaveSameLenMatcher{
		Item: item,
	}
}

package matchers

import (
	"fmt"

	"github.com/onsi/gomega/format"
)

type HaveSameLenMatcher struct {
	Item interface{}
}

func (matcher HaveSameLenMatcher) Match(actual interface{}) (success bool, err error) {
	length, ok := lengthOf(actual)
	if !ok {
		return false, fmt.Errorf("HaveSameLen matcher expects a string/array/map/channel/slice.  Got:\n%s", format.Object(actual, 1))
	}
	itemLength, ok := lengthOf(matcher.Item)
	if !ok {
		return false, fmt.Errorf("HaveSameLen matcher expects a string/array/map/channel/slice.  Got:\n%s", format.Object(actual, 1))
	}
	return length == itemLength, nil
}

func (matcher HaveSameLenMatcher) FailureMessage(actual interface{}) (message string) {
	itemLength, ok := lengthOf(matcher.Item)
	if !ok {
		return fmt.Sprintf("HaveSameLen matcher expects a string/array/map/channel/slice.  Got:\n%s", format.Object(actual, 1))
	}
	return fmt.Sprintf("Expected\n%s\nto have length %d", format.Object(actual, 1), itemLength)
}

func (matcher HaveSameLenMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	itemLength, ok := lengthOf(matcher.Item)
	if !ok {
		return fmt.Sprintf("HaveSameLen matcher expects a string/array/map/channel/slice.  Got:\n%s", format.Object(actual, 1))
	}
	return fmt.Sprintf("Expected\n%s\nnot to have length %d", format.Object(actual, 1), itemLength)
}

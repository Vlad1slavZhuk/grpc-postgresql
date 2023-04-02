package domain

import "fmt"

const (
	Done = iota
	New
	Pending
	Invalid
)

var (
	statusInt = map[int32]string{
		Done:    "done",
		New:     "new",
		Pending: "pending",
		Invalid: "invalid",
	}

	statusStr = map[string]int32{
		"done":    Done,
		"new":     New,
		"pending": Pending,
		"invalid": Invalid,
	}
)

func ValidStrStatus(s string) (int32, error) {
	result, ok := statusStr[s]
	if !ok {
		return 0, fmt.Errorf("invalid category: %v", s)
	}
	return result, nil
}

func ValidIntStatus(i int32) string {
	result, ok := statusInt[i]
	if !ok {
		return ""
	}
	return result
}

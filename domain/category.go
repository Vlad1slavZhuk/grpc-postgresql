package domain

import "fmt"

const (
	Phone = iota + 1
	Tablet
	Notebook
	TV
	Other
)

var (
	categoryInt = map[int32]string{
		Phone:    "phone",
		Tablet:   "tablet",
		Notebook: "notebook",
		TV:       "tv",
		Other:    "other",
	}

	categoryStr = map[string]int32{
		"phone":    Phone,
		"tablet":   Tablet,
		"notebook": Notebook,
		"tv":       TV,
		"other":    Other,
	}
)

func ValidStrCategory(s string) (int32, error) {
	result, ok := categoryStr[s]
	if !ok {
		return 0, fmt.Errorf("invalid category: %v", s)
	}
	return result, nil
}

func ValidIntCategory(i int32) string {
	result, ok := categoryInt[i]
	if !ok {
		return ""
	}
	return result
}

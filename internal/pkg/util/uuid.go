package util

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateUUIDs(n int, upper bool) []string {
	result := make([]string, 0)
	for i := 0; i < n; i++ {
		result = append(result, GenerateUUID(upper))
	}
	return result
}

func GenerateUUID(u bool) string {
	if u {
		return strings.ToUpper(uuid.NewString())
	} else {
		return uuid.NewString()
	}
}

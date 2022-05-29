package functions

import "github.com/google/uuid"

func UUID() string {
	return uuid.New().String()
}

func seq(n int) []int {
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, i)
	}
	return res
}

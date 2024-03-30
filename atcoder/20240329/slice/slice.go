package slice

import "strconv"

func Atoi(s []string) ([]int, error) {
	r := make([]int, len(s))
	for i, v := range s {
		c, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		r[i] = c
	}
	return r, nil
}

package helpers

import (
	"fmt"
	"strconv"
)

func CastStringArrayToInt64Array(values []string) ([]int64, error) {
	result := make([]int64, len(values))
	for i, v := range values {
		val, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("helper.CastStringArrayToInt64Array: error converting %s to int: %w", v, err)
		}

		result[i] = val
	}

	return result, nil
}

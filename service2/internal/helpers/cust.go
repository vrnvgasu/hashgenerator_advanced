package helpers

import (
	"context"
	"fmt"
	"service2/internal/lg"
	"strconv"
)

func CastStringArrayToInt64Array(ctx context.Context, values []string) ([]int64, error) {
	result := make([]int64, len(values))
	for i, v := range values {
		val, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			em := fmt.Errorf("helper.CastStringArrayToInt64Array: error converting %s to int: %w", v, err)
			lg.Error(ctx, "helper.CastStringArrayToInt64Array", "helpers", err)
			return nil, em
		}

		result[i] = val
	}

	return result, nil
}

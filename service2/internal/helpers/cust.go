package helpers

import (
	"fmt"
	"service2/internal/lg"
	"strconv"

	"github.com/vrnvgasu/logwrapper"
)

func CastStringArrayToInt64Array(values []string) ([]int64, error) {
	result := make([]int64, len(values))
	for i, v := range values {
		val, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			em := fmt.Errorf("helper.CastStringArrayToInt64Array: error converting %s to int: %w", v, err)
			lg.Logger.Payload(logwrapper.NewPayload().Op("helper.CastStringArrayToInt64Array").Package("main")).Error(em)
			return nil, em
		}

		result[i] = val
	}

	return result, nil
}

package hashhandler

import (
	"service2/models"
)

func FilterWhereNotIn(list []string, hashes []*models.Hash) []string {
	existedHashMap := make(map[string]any, len(hashes))
	for _, hash := range hashes {
		if hash.Hash == nil {
			continue
		}

		existedHashMap[*hash.Hash] = struct{}{}
	}

	result := make([]string, 0)
	for _, s := range list {
		if _, ok := existedHashMap[s]; !ok {
			result = append(result, s)
		}
	}

	return result
}

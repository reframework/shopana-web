package loaderV2Utils

import (
	"strings"

	"github.com/graph-gophers/dataloader"
)

func CollectFields(loaderKeys []dataloader.Key /** <key>.<gql-field> */) ([]string, []string) {
	keys := []string{}
	fields := []string{}

	keysSet := map[string]int8{}
	fieldSet := map[string]int8{}

	for _, lk := range loaderKeys {
		parts := strings.SplitN(lk.String(), ".", 2)

		if len(parts) == 2 {
			k, f := parts[0], parts[1]

			if _, ok := keysSet[k]; !ok {
				keysSet[parts[0]] = 1
				keys = append(keys, k)
			}

			if _, ok := fieldSet[f]; !ok {
				fieldSet[parts[1]] = 1
				fields = append(fields, f)
			}
		}
	}

	return keys, fields
}

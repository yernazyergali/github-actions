package internal

import "strings"

func StringToArrayConversion(names string) ([]string, error) {
	list := strings.Split(names, ",")
	data := make([]string, 0, len(list))

	for _, name := range list {
		data = append(data, name)
	}

	return data, nil
}

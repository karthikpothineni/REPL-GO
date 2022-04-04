package command

// copyMap makes a copy of the input map
func copyMap(input map[string]string) map[string]string {
	output := make(map[string]string, len(input))

	for key, val := range input {
		output[key] = val
	}

	return output
}

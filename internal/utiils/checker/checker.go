package checker

func IsContains[T comparable](container []T, target []T) ([]T, bool) {
	containerMap := map[T]bool{}

	output := []T{}

	for _, v := range container {
		containerMap[v] = false
	}

	for _, v := range target {
		if _, has := containerMap[v]; has {
			output = append(output, v)
		}
	}

	return output, len(output) > 0
}

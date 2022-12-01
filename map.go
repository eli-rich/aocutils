package main

func MapMirror[K comparable, V comparable](m map[K]V) map[V]K {
	// Returns a mirrored map
	result := make(map[V]K)
	for k, v := range m {
		result[v] = k
	}
	return result
}

func MapKeys[K comparable, V comparable](m map[K]V) []K {
	// Returns an array of map keys
	var result []K
	for k := range m {
		result = append(result, k)
	}
	return result
}

func MapValues[K comparable, V comparable](m map[K]V) []V {
	// Returns an array of map values
	var result []V
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func MapEntries[K comparable, V comparable](m map[K]V) ([]K, []V) {
	// Converts a map to an array of keys and values
	keys := MapKeys(m)
	values := MapValues(m)
	return keys, values
}

func MapFromEntries[K comparable, V comparable](keys []K, values []V) map[K]V {
	// Creates a map from an array of keys and values
	result := make(map[K]V)
	for index, key := range keys {
		result[key] = values[index]
	}
	return result
}

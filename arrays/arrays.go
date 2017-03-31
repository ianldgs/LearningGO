package arrays

func StringUnique(arr []string) (arrUnique []string) {
	items := make(map[string]bool)

	for _, item := range arr  {
		items[item] = true
	}

	for item := range items {
		arrUnique = append(arrUnique, item)
	}

	return
}

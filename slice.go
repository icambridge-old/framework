package framework

func SliceHasString(wantedString string, stringSlice []string) bool {

	for _, found := range stringSlice {
		if found == wantedString {
			return true
		}
	}

	return false
}

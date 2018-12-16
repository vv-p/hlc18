func getAccount(testString string) string {

	startIndex := -1
	endIndex := -1
	depth := 0
	nextStartIndex := -1


	if ! strings.ContainsAny(testString, "{}") {
		return testString  // we're in the middle of somewhere, so we need to return all
	} 
	
	startIndex = strings.Index(testString, "{")
	if startIndex > -1 { 
		depth += 1 // start "{" was found, it's our normal issue
	}

	endIndex = strings.Index(testString[startIndex:], "}")
	if endIndex == -1 {
		// we found start "{" without end "}"
		return testString[startIndex:]
	}

	// more interesing issue, we have "}" in our testString
	// so we need to check another level of depth here looking for the next "{"
	// between start "{" and end "}"
	nextStartIndex = strings.Index(testString[startIndex + 1:endIndex], "{")
	if nextStartIndex == -1 {  // no more levels in this depth
		depth -= 1
		if depth == 0 {
			// we're on surface now
			//toJson()
		}
		return testString[startIndex:endIndex + 1]
	}

	// most interesting issue, we have next "{" that means next level of our json
	depth += 1

	return getAccount(testString[nextStartIndex:])
}
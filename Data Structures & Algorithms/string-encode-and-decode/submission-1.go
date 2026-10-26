type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder

	// Transform into format <length>;<string> like this:
	// 3;abc6;qwerty
	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteRune(';')
		sb.WriteString(str)
	}

	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}

	i := 0
	for i < len(encoded) {
		j := i+1
		for encoded[j] != ';' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j]) // get length of the following string
		i = j+1 // skip delimiter
		result = append(result, encoded[i:i+length]) // get the string itself
		i += length // go to the next encoded substring
	}

	return result
}

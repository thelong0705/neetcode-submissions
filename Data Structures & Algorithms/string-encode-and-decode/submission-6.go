type Solution struct{}

const sep = "あ"

func (s *Solution) Encode(strs []string) string {
	res := ""
	for _, str := range strs {
		res += str
		res += sep
	}

	return res
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}
	temp := []rune{}

	var sepRune rune

	for _, c := range sep {
		sepRune = c
	}
	
	for _, c := range encoded {
		if c != sepRune {
			temp = append(temp, c)
		} else {
			res = append(res, string(temp))
			temp = []rune{}
		}
	}

	return res
}

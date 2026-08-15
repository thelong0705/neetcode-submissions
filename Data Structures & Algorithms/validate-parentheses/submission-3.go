func isValid(s string) bool {
	st := stack{
		sl: []byte{},
	}

	m := map[byte]byte{
		')' : '(',
		'}' : '{',
		']' : '[',
	}

	for i:=0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			st.sl = append(st.sl, s[i])
			fmt.Println(st.sl)
		} else {
			c, ok := st.pop()
			fmt.Println(c)
			if !ok {
				return false
			}

			if m[s[i]] != c {
				return false
			}
		}
	}

	return len(st.sl) == 0
}

type stack struct{
	sl []byte
}

func(st *stack) pop() (byte, bool) {
	if len(st.sl) == 0 {
		return 0, false
	}

	x := st.sl[len(st.sl) - 1]
	st.sl = st.sl[:len(st.sl) - 1]

	return x, true
}
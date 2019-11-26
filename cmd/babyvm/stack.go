package main

type stack []uint8

func (s stack) Push(v uint8) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, uint8) {
	var l = len(s)

	if l == 0 {
		return s[:], 0
	} else {
		return s[:l-1], s[l-1]
	}
}

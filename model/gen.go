package model

type Gen []bool

func NewGen(length uint) Gen {
	gen := make(Gen, length)
	halfLength := length >> 1
	gen[halfLength] = true
	return gen
}

func (g Gen) Next(rule Rule) Gen {
	next := make(Gen, len(g))
	for i := 1; i < len(g)-1; i++ {
		first := g[i-1]
		second := g[i]
		third := g[i+1]

		next[i] = rule[first][second][third]
	}

	return next
}

var (
	byteOne  = "#"[0]
	byteZero = " "[0]
)

func (g Gen) String() string {
	s := make([]byte, len(g))

	for i, cell := range g {
		if cell {
			s[i] = byteOne
		} else {
			s[i] = byteZero
		}
	}

	return string(s)
}

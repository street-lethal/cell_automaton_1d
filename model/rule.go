package model

type Rule map[bool]map[bool]map[bool]bool

func NewRule(ruleNumber uint) *Rule {
	rule := make(Rule)

	for i := 0; i <= 7; i++ {
		bit := (ruleNumber>>i)&1 == 1

		first := i&4 == 4
		second := i&2 == 2
		third := i&1 == 1

		if _, ok := rule[first]; !ok {
			rule[first] = make(map[bool]map[bool]bool)
		}
		if _, ok := rule[first][second]; !ok {
			rule[first][second] = make(map[bool]bool)
		}

		rule[first][second][third] = bit
	}

	return &rule
}

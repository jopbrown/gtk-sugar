package sugar

type CandyGiver func(c Candy)

var candyGiverReg = make(map[string]CandyGiver, 10)

func RegisterCandyGiver(name string, giver CandyGiver) {
	candyGiverReg[name] = giver
}

func GiveCandy(c Candy) {
	for _, giver := range candyGiverReg {
		giver(c)
	}
}

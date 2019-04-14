package sugar

type CandyGiver func(c Candy)

var globalCandyGiverReg = make(map[string]CandyGiver, 10)

func RegisterGlobalCandyGiver(name string, giver CandyGiver) {
	globalCandyGiverReg[name] = giver
}

func GiveCandyToEveryone(c Candy) {
	for _, giver := range globalCandyGiverReg {
		giver(c)
	}
}

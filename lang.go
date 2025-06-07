package main

type Caption int

const (
	Play Caption = iota
	Restart
	Options
	Exit
	Volume
	Chance
	Odd
	Stake
	Balance
)

var (
	EN = map[Caption]string{
		Play:    "play",
		Restart: "restart",
		Options: "options",
		Exit:    "exit",
		Volume:  "Music Volume",
		Chance:  "chance",
		Odd:     "odd",
		Stake:   "stake",
		Balance: "balance",
	}
	PL = map[Caption]string{
		Play:    "graj",
		Restart: "restart",
		Options: "opcje",
		Exit:    "wyjdź",
		Volume:  "Głośność Muzyki",
		Chance:  "szansa",
		Odd:     "kurs",
		Stake:   "stawka",
		Balance: "Bilans",
	}
	Lang = PL
)

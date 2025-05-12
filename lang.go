package main

type Caption int

const (
	Start Caption = iota
	Restart
	Options
	Exit
	Volume
)

var (
	EN = map[Caption]string{
		Start:   "start",
		Restart: "restart",
		Options: "options",
		Exit:    "exit",
		Volume:  "Music Volume",
	}
	PL = map[Caption]string{
		Start:   "start",
		Restart: "restart",
		Options: "opcje",
		Exit:    "wyjdź",
		Volume:  "Głośność Muzyki",
	}
	Lang = PL
)

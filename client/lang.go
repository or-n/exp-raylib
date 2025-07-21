package main

type Caption int

const (
	Start Caption = iota
	Restart
	Options
	Exit
	ConfigureKeys
	Volume
)

var (
	EN = map[Caption]string{
		Start:         "start",
		Restart:       "restart",
		Options:       "options",
		Exit:          "exit",
		ConfigureKeys: "Configure Keys (Press to Change)",
		Volume:        "Music Volume",
	}
	PL = map[Caption]string{
		Start:         "start",
		Restart:       "restart",
		Options:       "opcje",
		Exit:          "wyjdź",
		ConfigureKeys: "Konfiguruj Klawisze (Naciśnij by Zmienić)",
		Volume:        "Głośność Muzyki",
	}
	Lang = PL
)

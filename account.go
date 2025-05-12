package main

import (
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"strconv"
)

type Account struct {
	Balance f64
	decided bool
	input   string
}

type Possibility struct {
	chance f64
	odd    f64
}

var (
	AccountFile = "asset/account.gob"
	MainAccount Account
	event       []Possibility
	n           = i32(1)
	keyToChar   = map[int32]rune{
		KeyZero:   '0',
		KeyOne:    '1',
		KeyTwo:    '2',
		KeyThree:  '3',
		KeyFour:   '4',
		KeyFive:   '5',
		KeySix:    '6',
		KeySeven:  '7',
		KeyEight:  '8',
		KeyNine:   '9',
		KeyPeriod: '.',
	}
)

func AccountInit() {
	if err := Load(AccountFile, &MainAccount); err != nil {
		fmt.Println("Error loading account:", err)
		MainAccount.Balance = 1000
		if err := Save(AccountFile, &MainAccount); err != nil {
			fmt.Println("Failed to save account:", err)
		}
	}
}

func mix(a, b, t f64) f64 {
	return a + (b-a)*t
}

func EventNew() {
	var sum f64
	v := make([]f64, n)
	for i := range n {
		v[i] = rand.Float64()
		sum += v[i]
	}
	event = make([]Possibility, n+1)
	margin := mix(0.7, 1.1, rand.Float64())
	for i := range n {
		p := v[i]
		event[i] = Possibility{chance: p, odd: margin / p}
	}
	event[n] = Possibility{chance: 1 - event[0].chance, odd: 0}
	MainAccount.decided = false
}

func EventDraw() {
	for i := range n {
		chance := fmt.Sprintf("chance: %.2f", event[i].chance)
		odd := fmt.Sprintf("odd: %.2f", event[i].odd)
		DrawText(chance, 200, 50*(i+1), 20, White)
		DrawText(odd, 400, 50*(i+1), 20, White)
	}
}

func PossibilityId(x f64) i32 {
	for i := range n {
		if x > event[i].chance {
			x -= event[i].chance
		} else {
			return i
		}
	}
	return n
}

func AccountUpdate(account *Account) {
	if account.input == "0" {
		account.input = ""
	}
	for key, char := range keyToChar {
		if IsKeyPressed(key) {
			account.input += string(char)
		}
	}
	if IsKeyPressed(KeyBackspace) {
		account.input = ""
	}
	if account.input == "" {
		account.input = "0"
	}
	if IsKeyPressed(KeyEnter) {
		stake, err := strconv.ParseFloat(account.input, 32)
		if err == nil && stake <= account.Balance {
			i := PossibilityId(rand.Float64())
			account.Balance += stake * (event[i].odd - 1)
			EventNew()
		}
	}
}

func AccountDraw(account *Account) {
	balance := fmt.Sprintf("balance: %.2f", account.Balance)
	DrawText(balance, 200, 500, 20, White)
	position := NewVector2(200, 600)
	padding := NewVector2(20, 20)
	text := "stake: " + account.input
	size := MeasureTextEx(GetFontDefault(), text, 20, 2)
	color := NewColor(0, 0, 0, 127)
	DrawRectangleV(position, Vector2Add(size, Vector2Scale(padding, 2)), color)
	text_position := Vector2Add(position, padding)
	DrawText(text, i32(text_position.X), i32(text_position.Y), 20, White)
}

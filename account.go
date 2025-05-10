package main

import (
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"strconv"
)

type Account struct {
	balance f64
	decided bool
	input   string
}

type Possibility struct {
	chance f64
	odd    f64
}

var (
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

func AccountInit(account *Account) {
	account.balance = 1000
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
	for key, char := range keyToChar {
		if IsKeyPressed(key) {
			account.input += string(char)
		}
	}
	length := len(account.input)
	if IsKeyPressed(KeyBackspace) && length > 0 {
		account.input = account.input[:length-1]
	}
	if IsKeyPressed(KeyEnter) {
		stake, err := strconv.ParseFloat(account.input, 32)
		if err == nil && stake <= account.balance {
			i := PossibilityId(rand.Float64())
			account.balance += stake * (event[i].odd - 1)
			EventNew()
		}
	}
}

func AccountDraw(account *Account) {
	balance := fmt.Sprintf("balance: %.2f", account.balance)
	DrawText(balance, 200, 500, 20, White)
	DrawText(account.input, 200, 600, 20, White)
}

package main

import (
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
	. "github.com/or-n/util-go"
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
	n           = i32(5)
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
	ShowOdd0 = false
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
		v[i] = mix(0.1, 1, rand.Float64())
		sum += v[i]
	}
	chance0 := rand.Float64()
	sum /= 1 - chance0
	event = make([]Possibility, n+1)
	margin := mix(0.7, 1.1, rand.Float64())
	x := margin / f64(n)
	for i := range n {
		p := v[i] / sum
		event[i] = Possibility{chance: p, odd: x / p}
	}
	event[n] = Possibility{chance: chance0, odd: 0}
	MainAccount.decided = false
}

func EventDraw() {
	line_x := i32(90)
	line_y := i32(40)
	col_x := line_x + 10*3
	w := 2*line_x + 10*5
	x := (i32(WindowSize.X) - w) / 2
	y := i32(100)
	count := n + 1
	if !ShowOdd0 {
		count -= 1
	}
	DrawRectangle(x, y, line_x+10*2, (count+1)*line_y+10*3-5, LightGray)
	DrawText(Lang[Chance], x+10, y+10, 20, Black)
	DrawRectangle(x+col_x, y, line_x+10*2, (count+1)*line_y+10*3-5, LightGray)
	DrawText(Lang[Odd], x+10+col_x, y+10, 20, Black)
	y += line_y + 10*2
	for i := range count {
		DrawRectangle(x+10, y, line_x, line_y-5, SkyBlue)
		chance := fmt.Sprintf("%.2f", event[i].chance)
		DrawText(chance, x+20, y+5, 20, Black)
		DrawRectangle(x+10+col_x, y, line_x, line_y-5, SkyBlue)
		odd := fmt.Sprintf("%.2f", event[i].odd)
		DrawText(odd, x+20+col_x, y+5, 20, Black)
		y += line_y
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
	balance := fmt.Sprintf("%s: %.2f", Lang[Balance], account.Balance)
	s := MeasureTextEx(GetFontDefault(), balance, 20, 2)
	p := Vector2Scale(Vector2Subtract(WindowSize, s), 0.5)
	DrawText(balance, i32(p.X), i32(p.Y)+200, 20, White)
	padding := NewVector2(20, 20)
	text := Lang[Stake] + ": " + account.input
	size := MeasureTextEx(GetFontDefault(), text, 20, 2)
	color := NewColor(0, 0, 0, 127)
	size_with_padding := Vector2Add(size, Vector2Scale(padding, 2))
	position := Vector2Scale(Vector2Subtract(WindowSize, size_with_padding), 0.5)
	DrawRectangleV(position, size_with_padding, color)
	text_position := Vector2Add(position, padding)
	DrawText(text, i32(text_position.X), i32(text_position.Y), 20, White)
}

package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

const DieSides int = 6

var (
	app         = cli.App("passphrase", "Generate random word based passphrases")
	roles       = app.IntOpt("r roles", 4, "Number of roles (words to use)")
	camelCase   = app.BoolOpt("c camelCase", true, "Format passphrase as camelCase")
	shortList   = app.BoolOpt("s shortlist", false, "Use words from the \"short list\"")
	specialChar = app.BoolOpt("sp specialChar", false, "Include a special character")
	digit       = app.BoolOpt("d, digit", false, "Include a random digit")
	verbose     = app.BoolOpt("v, verbose", false, "Show verbose information")
)

type Dice struct {
	Die    []*Die
	Result int64
}

func NewDice(count int) *Dice {
	dice := &Dice{
		Die: make([]*Die, count),
	}
	for i := 0; i < count; i++ {
		dice.Die[i] = &Die{}
	}
	return dice
}

func (dice *Dice) Roll() {
	padding := func(n int64) int64 { // Pad numbers to "join" them together.
		var p int64 = 1
		for p < n {
			p *= 10
		}
		if p < 10 {
			p = 10
		}
		return p
	}
	var result int64
	for i, die := range dice.Die {
		die.Roll()
		if *verbose {
			fmt.Printf("Die[%d/%d] rolled: %d\n", i, len(dice.Die), die.Result)
		}
		result = result*padding(die.Result) + die.Result
	}
	dice.Result = result

}

type Die struct {
	Result int64
}

func (die *Die) Roll() {
	for {
		die.Result = int64(rand.Intn(DieSides))
		if die.Result > 0 {
			break
		}
	}
}

func generate() {
	var (
		wordList   map[int64]string
		passPhrase string
	)
	if *shortList {
		wordList = ShortList
	} else {
		wordList = LargeList
	}
	for i := 0; i < *roles; i++ {
		var (
			dieCount int
			word     string
		)
		if *shortList {
			dieCount = 4
		} else {
			dieCount = 5
		}
		dice := NewDice(dieCount)
		dice.Roll()

		word = wordList[dice.Result]
		if *verbose {
			fmt.Printf("Dice rolled: %d (%s)\n", dice.Result, word)
		}
		if *camelCase {
			if i >= 1 {
				word = fmt.Sprintf("%s%s", strings.ToUpper(string(word[0])), string(word[1:]))
			}
		}
		passPhrase += word
	}
	if *specialChar {
		passPhrase += SpecialCharacters[rand.Intn(len(SpecialCharacters))]
	}
	if *digit {
		passPhrase += Digits[rand.Intn(len(Digits))]
	}
	if *verbose {
		fmt.Println("Passphrase entropy: ", math.Log2(float64(len(wordList)))*float64(*roles))
	}
	fmt.Println(passPhrase)
}

func main() {
	app.Action = generate
	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards    string
	strength int
	bid      int
}

type Hands []*Hand

const HIGH_CARD = 0
const ONE_PAIR = 1
const TWO_PAIR = 2
const THREE_OF_A_KIND = 3
const FULL_HOUSE = 4
const FOUR_OF_A_KIND = 5
const FIVE_OF_A_KIND = 6

var strength = map[byte]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

func (hand *Hand) greaterThan(other *Hand) bool {
	if hand.strength > other.strength {
		return true
	} else if hand.strength < other.strength {
		return false
	}
	for i := range hand.cards {
		ourStrength := strength[hand.cards[i]]
		otherStrength := strength[other.cards[i]]
		if ourStrength > otherStrength {
			return true
		} else if ourStrength < otherStrength {
			return false
		}
	}
	return false
}

func (hands Hands) sort() {
	for i := len(hands) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if hands[j].greaterThan(hands[j+1]) {
				tmp := hands[j+1]
				hands[j+1] = hands[j]
				hands[j] = tmp
			}
		}
	}
}

func readHand(cards string, bid int) *Hand {
	var strength int
	seen := []byte{}
	counts := ""
	for _, card := range []byte(cards) {
		if !slices.Contains(seen, card) {
			seen = append(seen, card)
			counts += fmt.Sprintf("%d", strings.Count(cards, string(card)))
		}
	}

	if strings.Contains(counts, "5") {
		strength = FIVE_OF_A_KIND
	} else if strings.Contains(counts, "4") {
		strength = FOUR_OF_A_KIND
	} else if strings.Contains(counts, "3") {
		if strings.Contains(counts, "2") {
			strength = FULL_HOUSE
		} else {
			strength = THREE_OF_A_KIND
		}
	} else {
		pairCount := strings.Count(counts, "2")
		if pairCount == 2 {
			strength = TWO_PAIR
		} else if pairCount == 1 {
			strength = ONE_PAIR
		} else {
			strength = HIGH_CARD
		}
	}
	return &Hand{cards: cards, strength: strength, bid: bid}
}

func stringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("Could not convert string %s to int", str)
	}
	return num
}

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	hands := Hands{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		hands = append(hands, readHand(parts[0], stringToInt(parts[1])))
	}
	hands.sort()

	total := 0
	for i, hand := range hands {
		total += (i + 1) * hand.bid
	}
	fmt.Printf("Total winnings: %d\n", total)
}

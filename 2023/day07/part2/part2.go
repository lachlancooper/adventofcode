package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(winnings(scanner))
}

type (
	handStrength int
	cardLabel    int
)

const (
	highCard handStrength = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func (s handStrength) String() string {
	switch s {
	case highCard:
		return "high card"
	case onePair:
		return "one pair"
	case twoPair:
		return "two pair"
	case threeOfAKind:
		return "three of a kind"
	case fullHouse:
		return "full house"
	case fourOfAKind:
		return "four of a kind"
	case fiveOfAKind:
		return "five of a kind"
	}

	return ""
}

const (
	joker cardLabel = iota + 1
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	queen cardLabel = iota + 2
	king
	ace
)

// hand represents a hand of five cards
type hand struct {
	// cards is the list of cards in the hand
	cards string
	// bid is the size of the bid for the hand
	bid int
	// strength is the ranking of the hand's type
	strength handStrength
}

// winnings calculates the total winnings of a set of hands, by adding up the result of multiplying each hand's bid with its rank.
func winnings(scanner *bufio.Scanner) int {
	hands := []hand{}

	// read all cards into slice, calculating their strength
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		cards := line[0]
		bid, _ := strconv.Atoi(line[1])

		hands = append(hands, hand{
			cards:    cards,
			bid:      bid,
			strength: strength(cards),
		})
	}

	// sort slice from lowest to highest
	slices.SortFunc(hands, cmpHands)

	// read cards from lowest to highest rank, adding each card's winnings to the total
	total := 0
	for i, hand := range hands {
		total += (i + 1) * hand.bid
	}

	return total
}

// strength calculates the strength of a hand of cards.
func strength(cards string) handStrength {
	cardCounts := make(map[cardLabel]int)

	for i := range cards {
		cardCounts[label(cards[i])]++
	}

	max := 0
	for _, v := range cardCounts {
		if v > max {
			max = v
		}
	}

	// identify hands based on high or low card counts
	switch max {
	case 5:
		return fiveOfAKind
	case 4:
		if cardCounts[joker] >= 1 {
			return fiveOfAKind
		} else {
			return fourOfAKind
		}
	case 1:
		if cardCounts[joker] >= 1 {
			return onePair
		} else {
			return highCard
		}
	}

	seenTwo := false
	seenThree := false

	// max is either 2 or 3, now go through remaining cards
	for k, v := range cardCounts {
		switch v {
		case 3:
			if k != joker {
				if cardCounts[joker] >= 2 { // 222JJ
					return fiveOfAKind
				} else if cardCounts[joker] >= 1 { // 2223J
					return fourOfAKind
				} else if seenTwo { // 22233
					return fullHouse
				} // 22234, fall through
			} // JJJ23 or JJJ22, fall through
			seenThree = true
		case 2:
			if k != joker {
				if cardCounts[joker] >= 3 { // 22JJJ
					return fiveOfAKind
				} else if cardCounts[joker] >= 2 { // 22JJ3
					return fourOfAKind
				} else if seenThree { // 22333
					return fullHouse
				} else if seenTwo && cardCounts[joker] >= 1 { // 2233J
					return fullHouse
				} else if seenTwo { // 22334
					return twoPair
				} // 22345, fall through
			} // JJ234 or JJ224, fall through
			seenTwo = true
		}
	}

	switch {
	case seenTwo:
		if cardCounts[joker] >= 1 {
			return threeOfAKind
		} else {
			return onePair
		}
	case seenThree:
		if cardCounts[joker] >= 1 {
			return fourOfAKind
		} else {
			return threeOfAKind
		}
	}

	return highCard
}

// cmpHands returns a negative number when a < b, or a positive number when a > b.
func cmpHands(a, b hand) int {
	// primary ordering: compare strength
	s := cmpStrength(a, b)
	if s != 0 {
		return s
	}

	// secondary ordering: compare labels
	return cmpLabels(a, b)
}

// cmpStrength compares hands by their strength.
// Returns a negative number when a < b, a positive number when a > b, or zero when a == b.
func cmpStrength(a, b hand) int {
	return int(a.strength - b.strength)
}

// cmpLabels compares hands by their card labels.
// Returns a negative number when a < b, a positive number when a > b, or zero when a == b.
func cmpLabels(a, b hand) int {
	for i := range a.cards {
		if a.cards[i] == b.cards[i] {
			continue
		}

		return int(label(a.cards[i])) - int(label(b.cards[i]))
	}

	return 0
}

// label gets the value of a card label.
func label(c byte) cardLabel {
	val, err := strconv.Atoi(string(c))
	if err == nil {
		return cardLabel(val)
	}

	switch c {
	case 'T':
		return ten
	case 'J':
		return joker
	case 'Q':
		return queen
	case 'K':
		return king
	case 'A':
		return ace
	}

	return 0
}

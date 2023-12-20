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

const (
	two cardLabel = iota + 2
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
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
	cardCounts := make(map[rune]int)
	for _, c := range cards {
		cardCounts[c]++
	}

	counts := []int{}
	for _, val := range cardCounts {
		counts = append(counts, val)
	}

	slices.Sort(counts)
	slices.Reverse(counts)

	switch counts[0] {
	case 5:
		return fiveOfAKind
	case 4:
		return fourOfAKind
	case 3:
		if counts[1] == 2 {
			return fullHouse
		}
		return threeOfAKind
	case 2:
		if counts[1] == 2 {
			return twoPair
		}
		return onePair
	default:
		return highCard
	}
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
		return jack
	case 'Q':
		return queen
	case 'K':
		return king
	case 'A':
		return ace
	}

	return 0
}

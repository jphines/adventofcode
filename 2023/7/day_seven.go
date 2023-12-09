package main

import (
	"fmt"
	"reflect"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var strengthOne = "AKQJT98765432"
var strengthTwo = "AKQT98765432J"

type CamelHand struct {
	Hand string
	Bid  int
}

func (c *CamelHand) String() string {
	return fmt.Sprintf("(hand:%v bid:%v)", c.Hand, c.Bid)
}

type CH struct {
	Strength   string
	JokersWild bool
	Hands      []*CamelHand
}

func (c CH) Len() int      { return len(c.Hands) }
func (c CH) Swap(i, j int) { c.Hands[i], c.Hands[j] = c.Hands[j], c.Hands[i] }
func (c CH) Less(i, j int) bool {
	iScore := c.ScoreHand(c.Hands[i].Hand)
	jScore := c.ScoreHand(c.Hands[j].Hand)
	if iScore != jScore {
		return iScore < jScore
	} else {
		return c.StrengthCompare(c.Hands[i].Hand, c.Hands[j].Hand)
	}
}

func (c CH) StrengthCompare(hand1, hand2 string) bool {
	for i := 0; i < len(hand1); i++ {
		if hand1[i] == hand2[i] {
			continue
		}
		for j := 0; j < len(c.Strength); j++ {
			if hand1[i] == c.Strength[j] {
				return false
			} else if hand2[i] == c.Strength[j] {
				return true
			}
		}
	}
	return false
}

func (c CH) ScoreHand(hand string) int {
	occurrences := make(map[string]int)
	for _, char := range hand {
		occurrences[string(char)]++
	}

	if c.JokersWild {
		keys := make([]string, 0, len(occurrences))
		for key, _ := range occurrences {
			keys = append(keys, key)
		}
		sort.SliceStable(keys, func(i, j int) bool {
			return occurrences[keys[i]] > occurrences[keys[j]]
		})
		if v, ok := occurrences["J"]; ok {
			for _, key := range keys {
				if key == "J" {
					continue
				}
				occurrences[key] += v
				break
			}
			delete(occurrences, "J")
		}
	}

	values := make([]int, 0, len(occurrences))
	for _, value := range occurrences {
		values = append(values, value)
	}
	slices.Sort(values)
	if reflect.DeepEqual(values, []int{1, 1, 1, 1, 1}) {
		return 1
	} else if reflect.DeepEqual(values, []int{1, 1, 1, 2}) {
		return 2
	} else if reflect.DeepEqual(values, []int{1, 2, 2}) {
		return 3
	} else if reflect.DeepEqual(values, []int{1, 1, 3}) {
		return 4
	} else if reflect.DeepEqual(values, []int{2, 3}) {
		return 5
	} else if reflect.DeepEqual(values, []int{1, 4}) {
		return 6
	} else {
		return 7
	}
}

// parse camel card hands into a slice of camel hands
func parseInput(input string, strength string, jokers bool) CH {
	strHands := strings.Split(input, "\n")
	hands := make([]*CamelHand, len(strHands))
	for i, strHand := range strHands {
		splitHands := strings.Split(strHand, " ")
		bid, _ := strconv.Atoi(splitHands[1])
		hand := &CamelHand{

			Hand: splitHands[0],
			Bid:  bid,
		}
		hands[i] = hand
	}
	return CH{
		Strength:   strength,
		JokersWild: jokers,
		Hands:      hands,
	}
}

func daySeven(input string, strength string, jokers bool) {
	c := parseInput(input, strength, jokers)
	sort.Sort(c)
	winnings := 0
	for i := 0; i < len(c.Hands); i++ {
		winnings += c.Hands[i].Bid * (i + 1)
	}
	fmt.Println(winnings)
}

package main

import (
	"fmt"
	"github.com/WulffTL/advent-of-code-2020/utils"
	"strconv"
	"strings"
)

// Bag ...
type Bag struct {
	color     string
	adjective string
}

// BagContent ...
type BagContent struct {
	bag    Bag
	amount int
}

// Rule ...
type Rule struct {
	bag         Bag
	bagContents []BagContent
}

func main() {
	inputPath := "day7/puzzle-input.txt"
	input := utils.GetArrayOfStringsFromFile(inputPath)

	answer := GetTotalBagCountWithinGold(input)

	fmt.Print(answer)
}

// GetTotalBagCountWhichCanContainGold ...
func GetTotalBagCountWhichCanContainGold(input []string) int {
	shinyGoldBag := Bag{
		adjective: "shiny",
		color:     "gold",
	}

	var rules []Rule
	for i := 0; i < len(input); i++ {
		rule := getRule(input[i])
		rules = append(rules, rule)
	}

	var bags []Bag
	bags = append(bags, shinyGoldBag)
	for i := 0; i < len(rules); i++ {
		rule := rules[i]

		for j := 0; j < len(bags); j++ {
			contains := containsBag(bags[j], rule)
			if contains && !isInBagsList(rule.bag, bags) {
				bags = append(bags, rule.bag)
				i = -1
				break
			}
		}
	}

	return len(bags) - 1
}

// GetTotalBagCountWithinGold ...
func GetTotalBagCountWithinGold(input []string) int {
	shinyGoldBag := Bag{
		adjective: "shiny",
		color:     "gold",
	}

	rules := map[Bag]Rule{}
	for i := 0; i < len(input); i++ {
		rule := getRule(input[i])
		rules[rule.bag] = rule
	}

	return getBagCountWithin(Bag{}, shinyGoldBag, rules) - 1

}

func getBagCountWithin(parentBag Bag, bag Bag, rules map[Bag]Rule) int {
	rule := rules[bag]
	bagCount := 1
	bagContents := rule.bagContents
	bagQuantity := 1
	parentRule := rules[parentBag]
	parentBagContents := parentRule.bagContents
	for i := 0; i < len(parentBagContents); i++ {
		if parentBagContents[i].bag == bag {
			bagQuantity = parentBagContents[i].amount
		}
	}

	if len(bagContents) == 0 {
		return bagQuantity
	}

	for i := 0; i < len(bagContents); i++ {
		innerBag := bagContents[i].bag
		bagCount += getBagCountWithin(bag, innerBag, rules)
	}

	return bagCount * bagQuantity
}

func isInBagsList(bag Bag, list []Bag) bool {
	for i := 0; i < len(list); i++ {
		if bag == list[i] {
			return true
		}
	}
	return false
}

func containsBag(bag Bag, rule Rule) bool {
	contents := rule.bagContents
	for i := 0; i < len(contents); i++ {
		if contents[i].bag == bag {
			return true
		}
	}
	return false
}

func getRule(ruleDesc string) Rule {
	parts := strings.Split(ruleDesc, "contain")
	bagDescription := strings.Split(parts[0], " ")
	bag := Bag{
		adjective: bagDescription[0],
		color:     bagDescription[1],
	}

	bagContentsDesc := strings.Split(parts[1], ",")
	bagContentsDesc[0] = strings.TrimSpace(bagContentsDesc[0])

	var rule Rule

	if bagContentsDesc[0] == "no other bags." {
		rule = Rule{
			bag:         bag,
			bagContents: []BagContent{},
		}
	} else {
		var bagContents []BagContent
		for i := 0; i < len(bagContentsDesc); i++ {
			bagContent := getBagContent(bagContentsDesc[i])
			bagContents = append(bagContents, bagContent)
		}
		rule = Rule{
			bag:         bag,
			bagContents: bagContents,
		}
	}

	return rule
}

func getBagContent(description string) BagContent {
	description = strings.TrimSpace(description)
	words := strings.Split(description, " ")

	amount, _ := strconv.ParseInt(words[0], 10, 64)
	adj := words[1]
	color := words[2]

	return BagContent{
		bag: Bag{
			color:     color,
			adjective: adj,
		},
		amount: int(amount),
	}
}

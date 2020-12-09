package day7

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"dfortier.org/advent2020/pkg/util"
)

type bag struct {
	color   string
	subbags []subbag
	visited bool
}

type subbag struct {
	nbBags int
	bag    *bag
}

func newSubbag(color string, nbBags int) subbag {
	return subbag{
		nbBags: nbBags,
		bag:    newBag(color),
	}
}

func newSubbagFromExisting(bag *bag, nbBags int) subbag {
	return subbag{
		nbBags: nbBags,
		bag:    bag,
	}
}

func newBag(newColor string) *bag {
	return &bag{
		color:   newColor,
		subbags: make([]subbag, 0),
		visited: false,
	}
}

type rules struct {
	bags map[string]*bag
}

func readData() rules {
	var f *os.File
	var err error
	var subbag subbag
	var result = rules{
		bags: make(map[string]*bag, 0),
	}
	if f, err = os.Open("bagrules.txt"); err != nil {
		panic("Unable to read file")
	}

	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()
		bagAndChild := strings.Split(line, "contain")
		parentBagColor := extractColorName(bagAndChild[0])
		var parentBag *bag
		var ok bool
		if parentBag, ok = result.bags[parentBagColor]; !ok {
			parentBag = newBag(parentBagColor)
			result.bags[parentBag.color] = parentBag
		}

		if bagAndChild[1] != " no other bags." {
			childrenBagsString := strings.Split(bagAndChild[1], ",")
			for _, childBagString := range childrenBagsString {
				qte, bagColor := extractSubbag(childBagString)
				// Find if we know this one
				if childBag, ok := result.bags[bagColor]; ok {
					subbag = newSubbagFromExisting(childBag, qte)
				} else {
					// Else create it
					subbag = newSubbag(bagColor, qte)
					result.bags[subbag.bag.color] = subbag.bag
				}
				parentBag.subbags = append(parentBag.subbags, subbag)
			}
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return result
}

func extractSubbag(oneSubbag string) (int, string) {
	childBagString := strings.TrimSpace(oneSubbag)
	qteString := childBagString[:1]
	qte := util.Convert(qteString)
	bagColor := childBagString[1:]

	return qte, extractColorName(bagColor)
}

func extractColorName(color string) string {
	color = strings.ReplaceAll(color, "bags.", "")
	color = strings.ReplaceAll(color, "bags", "")
	color = strings.ReplaceAll(color, "bag.", "")
	color = strings.ReplaceAll(color, "bag", "")

	return strings.TrimSpace(color)
}

func Day1VisitPart1() {
	rules := readData()

	count := 0
	for k, v := range rules.bags {
		println(fmt.Sprintf("Processing bag %s", k))
		if visitBagTree(v) {
			count++
		}
	}
	println(fmt.Sprintf("Found %d that can hold shiny gold bag", count))
}

func visitBagTree(oneBag *bag) bool {
	println(fmt.Sprintf("Visiting %s", oneBag.color))
	for _, child := range oneBag.subbags {
		if child.bag.color == "shiny gold" {
			return true
		}
		mayReachGold := visitBagTree(child.bag)
		if mayReachGold {
			return true
		}
	}

	return false
}

func Day1VisitPart2() {
	rules := readData()

	goldBag := rules.bags["shiny gold"]

	count := visitBag(goldBag)
	println(fmt.Sprintf("Found %d that can hold shiny gold bag", count))
}

func visitBag(oneBag *bag) int {
	println(fmt.Sprintf("Visiting %s", oneBag.color))

	totalChildCount := 0
	for _, child := range oneBag.subbags {
		childCount := visitBag(child.bag)
		childCount = child.nbBags * childCount
		totalChildCount += childCount
		println(fmt.Sprintf("Total child for %s %d", oneBag.color, totalChildCount))
	}
	if oneBag.color != "shiny gold" {
		totalChildCount = totalChildCount + 1
	}
	return totalChildCount
}

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type operateFn func(wr int) int
type Monkey struct {
	items                  []int
	divisor                int
	numberOfInspectedItems int
	operate                operateFn
	throwToIndex           [2]int
}

// monkey
// items its holding
// inspects items -> does operation &
// Worry level / 3 (floor)
// checks result agains condition to decide who to throw item to
// repeat for all items

// repeat for all monkeys = round

var rounds int = 10000

var monkeys []Monkey

func main() {

	file, err := os.ReadFile("./day11/test.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	fileText := string(file)
	loadMonkeys(fileText)

	tests := make([]int, len(monkeys))

	for i, mk := range monkeys {
		tests[i] = mk.divisor
	}
	divisor := calculateLCMOfSlice(tests)

	for r := 0; r < rounds; r++ {

		for i := 0; i < len(monkeys); i++ {
			monkeys[i].takeTurn(divisor)
		}

	}

	var inspectedItemsAr []int

	for i := 0; i < len(monkeys); i++ {
		inspectedItemsAr = append(inspectedItemsAr, monkeys[i].numberOfInspectedItems)
	}
	sort.Slice(inspectedItemsAr, func(i int, j int) bool {
		return inspectedItemsAr[i] > inspectedItemsAr[j]
	})

	fmt.Println(inspectedItemsAr)

	fmt.Println(inspectedItemsAr[0] * inspectedItemsAr[1])

	// monkey business = top two inspected items * each other
}

func (m *Monkey) takeTurn(divisor int) {

	for len(m.items) != 0 {
		itemWorryLevel := m.inspect(0)

		// lcm := LCM(itemWorryLevel, m.divisor)
		// if lcm%m.divisor == itemWorryLevel%m.divisor {
		// itemWorryLevel = lcm
		// }

		throwTo := &monkeys[m.throwToIndex[m.test(itemWorryLevel)]]
		// if divisor != 0 {
		itemWorryLevel = itemWorryLevel % divisor
		// }
		throwTo.enqueu(itemWorryLevel)
		m.items = m.items[1:]
	}
}

func (m *Monkey) test(wr int) int {
	if wr%m.divisor == 0 {
		return 1
	}
	return 0
}

func (m *Monkey) inspect(i int) int {

	// take first item
	m.items[i] = m.operate(m.items[i])

	// m.items[i] = int(math.Floor(float64(m.items[i]) / 3.0))

	m.numberOfInspectedItems++

	return m.items[i]
}

func (m *Monkey) enqueu(item int) {
	m.items = append(m.items, item)
}

func multiplyFactory(n int) operateFn {
	return func(wr int) int {
		return wr * n
	}
}

func square(val int) int {
	// return int(math.Pow(float64(val), 2))
	return val
}

func addFactory(n int) operateFn {
	return func(wr int) int {
		return wr + n
	}
}

func loadMonkeys(txt string) {
	parts := strings.Split(txt, "\n\n")

	for i, mk := range parts {
		rows := strings.Split(mk, "\n")

		items := strings.Split(strings.Split(rows[1], ": ")[1], ",")

		monkeys = append(monkeys, Monkey{
			items: convertToIntList(&items),
		},
		)

		// operating
		operation := strings.Split(rows[2], "= ")[1]

		if strings.Contains(operation, "*") {
			if operation == "old * old" {
				monkeys[i].operate = square
			} else {
				sn := strings.TrimSpace(strings.Split(operation, " * ")[1])
				n, _ := strconv.Atoi(sn)
				monkeys[i].operate = multiplyFactory(n)
			}
		} else {
			sn := strings.TrimSpace(strings.Split(operation, " + ")[1])
			n, _ := strconv.Atoi(sn)
			monkeys[i].operate = addFactory(n)
		}

		// test
		testLine := strings.Split(rows[3], " ")
		n, _ := strconv.Atoi(testLine[len(testLine)-1])
		monkeys[i].divisor = n

		truMonkey, _ := strconv.Atoi(string(rows[4][len(rows[4])-1]))
		falsMonkey, _ := strconv.Atoi(string(rows[5][len(rows[5])-1]))
		monkeys[i].throwToIndex = [2]int{falsMonkey, truMonkey}
	}

	// for idx, mk := range monkeys {
	// 	fmt.Println("MONKEY INDEX: ", idx)
	// 	fmt.Println(mk.items)
	// 	fmt.Println(mk.operate(1))
	// 	fmt.Println(mk.test(10))
	// 	fmt.Println(mk.throwToIndex)
	// }
}

func convertToIntList(sl *[]string) []int {
	var intList []int
	for _, str := range *sl {
		num, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			fmt.Printf("Error converting string '%s' to int: %v\n", str, err)
			continue // Skip this element if it can't be converted
		}

		intList = append(intList, num)
	}
	return intList
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func calcLCM(integers ...int) int {
	return LCM(integers[0], integers[1], integers[2:]...)
}
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
func calculateLCM(a, b int) int {
	// Find the greatest common divisor (GCD) using Euclidean algorithm
	gcd := calculateGCD(a, b)

	// Calculate LCM using the formula: LCM(a, b) = (a * b) / GCD(a, b)
	return (a * b) / gcd
}

// Calculate the GCD of two integers using Euclidean algorithm
func calculateGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Calculate the LCM of a slice of integers
func calculateLCMOfSlice(numbers []int) int {
	if len(numbers) < 2 {
		// LCM is not defined for fewer than two numbers
		return 0
	}

	lcm := numbers[0]
	for i := 1; i < len(numbers); i++ {
		lcm = calculateLCM(lcm, numbers[i])
	}
	return lcm
}

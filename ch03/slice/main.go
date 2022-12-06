package main

import (
	"fmt"
	"sort"

	"github.com/study-onboard/master-go/common/ui"
)

func main() {
	simplyUse()
	subSlice()
	capAndSize()
	copySlice()
	sortSlice()
	appendArrayToSlice()
}

func simplyUse() {
	ui.PrintTitle("Simply Use")
	defer ui.PrintEnd()

	messages := make([]string, 20)
	fmt.Printf("message size: %d\n", len(messages))
	for index, message := range messages {
		fmt.Printf("%d - `%s`\n", index, message)
	}

	var emptySlice []string
	fmt.Printf("empty slice size: %d\n", len(emptySlice))
	if emptySlice == nil {
		fmt.Println("empty slice is nil")
	}

	emptySlice = append(emptySlice, "show me the money")
	fmt.Printf("empty slice size: %d\n", len(emptySlice))

	emptySlice = nil
	fmt.Printf("empty slice size: %d\n", len(emptySlice))

	emptySlice = append(emptySlice, "how do you turn this on")
	emptySlice = append(emptySlice, "make love, no war")
	fmt.Printf("empty slice size: %d\n", len(emptySlice))
}

func subSlice() {
	ui.PrintTitle("Sub Slice")
	defer ui.PrintEnd()

	printMessages := func(title string, messages []string) {
		fmt.Printf("%s: %s\n", title, messages)
	}

	messages := []string{
		"show me the money",
		"how do you turn this on",
		"good good study, day day up",
	}
	printMessages("messages", messages)

	slogans := messages[1:]
	slogans[0] = "make love, no war"
	printMessages("slogans", slogans)
	printMessages("messages", messages)
}

func capAndSize() {
	ui.PrintTitle("Cap and Size")
	defer ui.PrintEnd()

	printSliceInfo := func(messages []string) {
		fmt.Printf("cap: %d, size: %d\n", cap(messages), len(messages))
	}

	var messages []string
	printSliceInfo(messages)

	messages = append(messages, "show me the money")
	printSliceInfo(messages)

	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	messages = append(messages, "how do you turn this on")
	printSliceInfo(messages)
}

func copySlice() {
	ui.PrintTitle("Copy Slice")
	defer ui.PrintEnd()

	printStrings := func(name string, strings []string) {
		fmt.Println(name)
		fmt.Println("...................................")
		for _, item := range strings {
			fmt.Println(item)
		}
		fmt.Println()
		fmt.Println()
	}

	messages := []string{
		"show me the money",
		"how do you turn this on",
		"good good study, day day up",
		"love you love you go to dead",
	}
	printStrings("messages", messages)

	slogans := []string{
		"make love, no war",
		"long time no see",
	}
	printStrings("slogans", slogans)

	copy(messages, slogans)
	printStrings("messages", messages)
	printStrings("slogans", slogans)

	var target []string
	copy(target, messages)
	printStrings("before make, target", target)

	target = make([]string, 3)
	copy(target, messages)
	printStrings("after make, target", target)

	target[0] = "hahahahahahhahahaha"
	printStrings("after item update, target", target)
	printStrings("after item update, messages", messages)

	copy(target[1:], messages[3:])
	printStrings("after sub item copy, target", target)
}

func sortSlice() {
	ui.PrintTitle("Sort Slice")
	defer ui.PrintEnd()

	type person struct {
		name   string
		width  int
		height int
	}

	printPersons := func(description string, persons []person) {
		fmt.Println(description)
		fmt.Println(".......................................")

		for i, p := range persons {
			fmt.Printf("%d. %v\n", i, p)
		}

		fmt.Println()
	}

	persons := []person{
		{"kut", 4, 5},
		{"lana", 2, 8},
		{"qing", 9, 1},
	}
	printPersons("After initialize persons", persons)

	sort.Slice(persons, func(leftIndex, rightIndex int) bool {
		left := persons[leftIndex]
		right := persons[rightIndex]
		return left.width < right.width
	})
	printPersons("After persons sort by width", persons)

	sort.Slice(persons, func(leftIndex, rightIndex int) bool {
		left := persons[leftIndex]
		right := persons[rightIndex]
		return left.height < right.height
	})
	printPersons("After persons sort by height", persons)
}

func appendArrayToSlice() {
	ui.PrintTitle("Append array to slice")
	defer ui.PrintEnd()

	var messages []string

	var items [3]string

	items = [3]string{
		"111111",
		"111111",
		"111111",
	}
	messages = append(messages, items[:]...)

	items = [3]string{
		"211111",
		"211111",
		"211111",
	}
	messages = append(messages, items[:]...)

	items = [3]string{
		"311111",
		"311111",
		"311111",
	}
	messages = append(messages, items[:]...)

	for _, message := range messages {
		fmt.Println(message)
	}
}

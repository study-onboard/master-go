package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	englishMessage := "aB34e"
	chineseMessage := "一二三四五"
	printLine()

	// print content
	fmt.Printf("English message: %s\n", englishMessage)
	fmt.Printf("Chinese message: %s\n", chineseMessage)
	printLine()

	// test len
	fmt.Printf("English message len: %d\n", len(englishMessage))
	fmt.Printf("Chinese message len: %d\n", len(chineseMessage))
	printLine()

	// test count characters
	fmt.Printf("English message characters: %d\n", utf8.RuneCountInString(englishMessage))
	fmt.Printf("Chinese message characters: %d\n", utf8.RuneCountInString(chineseMessage))
	printLine()

	// upper and lower
	fmt.Printf("English message upper: %s\n", strings.ToUpper(englishMessage))
	fmt.Printf("Chinese message upper: %s\n", strings.ToUpper(chineseMessage))
	fmt.Printf("English message lower: %s\n", strings.ToLower(englishMessage))
	fmt.Printf("Chinese message lower: %s\n", strings.ToLower(chineseMessage))
	printLine()

	// check prefix and subfix
	if strings.HasPrefix(englishMessage, "aB") {
		fmt.Printf("English message has prefix `aB`\n")
	} else {
		fmt.Printf("English message has not prefix `aB`\n")
	}
	if strings.HasPrefix(chineseMessage, "一二") {
		fmt.Printf("Chinese message has prefix `一二`\n")
	} else {
		fmt.Printf("Chinese message has not prefix `一二`\n")
	}
	printLine()

	// index of text
	fmt.Printf("Index of `34` for English message: %d\n", strings.Index(englishMessage, "34"))
	fmt.Printf("Index of `abc` for English message: %d\n", strings.Index(englishMessage, "abc"))
	printLine()

	// will return bytes position
	fmt.Printf("Index of `三四` for Chinese message: %d\n", strings.Index(chineseMessage, "三四"))
	fmt.Printf(
		"Substring for Chinese message: %s\n",
		chineseMessage[strings.Index(chineseMessage, "三四"):],
	)
	printLine()

	// crazy substring style....
	subStringIndex := strings.Index(chineseMessage, "三四")
	subString := chineseMessage[subStringIndex : subStringIndex+len("三四")]
	fmt.Printf("Substring for `三四`: %s\n", subString)
	fmt.Printf("Index of `六七` for Chinese message: %d\n", strings.Index(chineseMessage, "六七"))
	printLine()

	// trim
	text := "  show me the money \t\n  "
	fmt.Printf("text: %q\n", text)
	fmt.Printf("trimed text: %q\n", strings.Trim(text, "\t\n "))
	fmt.Printf("trimed space text: %q\n", strings.TrimSpace(text))
	fmt.Printf("left trimed text: %q\n", strings.TrimLeft(text, "\t\n "))
	fmt.Printf("right trimed text: %q\n", strings.TrimRight(text, "\t\n "))
	printLine()

	// fields
	text = "KutZhang\tItManager\tABC"
	fields := strings.Fields(text)
	for i, field := range fields {
		fmt.Printf("%d. %s\n", i, field)
	}
	printLine()

	// compare
	text1 := "abc你我他"
	text2 := "你我他abc"
	if result := strings.Compare(text1, text2); result == 0 {
		fmt.Printf("%q equals %q\n", text1, text2)
	} else {
		fmt.Printf("%q not equals %q -> %d\n", text1, text2, result)
	}
	printLine()

	// case insensitive comparison
	text1 = "abcd"
	text2 = "ABCD"
	if strings.EqualFold(text1, text2) {
		fmt.Printf("%q equals %q\n", text1, text2)
	} else {
		fmt.Printf("%q not equals %q\n", text1, text2)
	}

	// replace
	text = "你是好人"
	fmt.Printf("text: %q\n", text)
	text = strings.Replace(text, "好人", "王八蛋", -1)
	fmt.Printf("after replace, text: %q\n", text)
	printLine()

	// split
	text = "人才 畜生 美国人"
	parts := strings.Split(text, " ")
	for i, part := range parts {
		fmt.Printf("%d. %s\n", i, part)
	}
}

// print line
func printLine() {
	fmt.Println("-------------------------------------------------------")
}

package parser

import (
	"bufio"
	"os"
	"regexp"

	"github.com/mattn/go-runewidth"
	"github.com/maxmellon/tmng-parser/ast"
	"github.com/maxmellon/tmng-parser/category"
)

func match(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}

// ReadFile read file
func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReaderSize(file, 4096)
	var result []string
	for line := ""; err == nil; line, err = reader.ReadString('\n') {
		if line == "" {
			continue
		}
		result = append(result, line)
	}
	return result
}

// Parse try parse tmng file
func Parse(data []string) []ast.Tml {
	var result []ast.Tml
	for i, l := range data {
		if match(`^■`, l) {
			node := ast.New(i, category.Header, l, 0, runewidth.StringWidth(l))
			result = append(result, node)
		} else if match(`^●`, l) {
			node := ast.New(i, category.Title, l, 0, runewidth.StringWidth(l))
			result = append(result, node)
		} else if match(`^◎`, l) {
			node := ast.New(i, category.SubTitle, l, 0, runewidth.StringWidth(l))
			result = append(result, node)
		} else if match(`^○`, l) {
			node := ast.New(i, category.ListItem, l, 0, runewidth.StringWidth(l))
			result = append(result, node)
		} else if match(`^\n|^\r`, l) {
			node := ast.New(i, category.Empty, l, 0, runewidth.StringWidth(l))
			result = append(result, node)
		} else if match(`^\s+○`, l) || match(`^\s+●`, l) || match(`^\s+◎`, l) {
			node := ast.New(i, category.Error, l, 0, runewidth.StringWidth(l))
			result = append(result, node)
		} else {
			node := ast.New(i, category.PlaneText, l, 0, runewidth.StringWidth(l))
			result = append(result, node)
		}
	}
	return result
}

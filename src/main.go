package main

import (
	"util"
	"lexer"
	"fmt"
	"token"
	"parser"
	"io"
	"os"
	"strings"
)

func main() {

	util.Read("er.puml", handleLine);
}

var block string = ""
var inBlock bool
var startLineN int

func handleLine(line string, lineN int) {
	if startOfBlock(line) {
		block += line
		inBlock = true
		startLineN = lineN
		return
	}
	if endOfBlock(line) {
		block += line
		inBlock = false
	}

	l := lexer.New(block, startLineN)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(os.Stdout, p.Errors())
		return
	}

	io.WriteString(os.Stdout, program.String())
	io.WriteString(os.Stdout, "\n")

	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}

//忽略该行
func startOfBlock(line string) bool {
	return strings.HasPrefix(strings.TrimLeft(line, " "), "entity");
}
func endOfBlock(line string) bool {
	return strings.TrimSpace(line) == "}";
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some  business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

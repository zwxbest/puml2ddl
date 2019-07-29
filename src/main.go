package main

import (
	"fmt"
	"io"
	"lexer"
	"os"
	"parser"
	"strings"
	"util"
)

func main() {

	util.Read("er.puml", handleLine);
}

var block  = ""
var startLineN int
var inBlock = false;

func handleLine(line string, lineN int) bool{
	if startOfBlock(line) {
		block += line
		inBlock = true
		startLineN = lineN
		return true
	}
	if endOfBlock(line) {
		block += line
		//开始处理
		l := lexer.New(block, startLineN)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(os.Stdout, p.Errors())
			return false
		}
		block = ""
		fmt.Printf("%+v\n", program)
		inBlock = false
	}
	if inBlock == true {
		block += line
		return true
	}
	return true
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

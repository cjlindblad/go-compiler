package repl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		// new stuff
		p := parser.New(l)
		program := p.ParseProgram()
		printJSON(program)

		// old stuff
		/*
			fmt.Fprintf(out, "%+v\n", program)
				for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
					fmt.Fprintf(out, "%+v\n", tok)
				}
		*/
	}
}

func printJSON(data interface{}) {
	b, _ := json.MarshalIndent(data, "", "  ")

	fmt.Println(string(b))
}

package main

import {
	"fmt"
	"errors"
}

type Dictionary map[string]interface{}

func last(stack) {
	stack[len(stack)-1]
}

func parser(re) {
	i := 0
	stack := [][]Dictionary{}
	dict := Dictionary{"type" : "wildcard", "quantifier" : "exactlyOne"}
	for (i < re.length) {
		const next = re[i]
		switch next {
			case '.':
				stack = append(stack, dict)
				i++
				continue
			case '?':
				lastElement := last(last(stack))
				if (lastElement && lastElement["quantifier"] != "exactlyOne" ){
					err := errors.New("Quantifier must follow an unquantified element or group")
				}
				lastElement["quantifier"] = "zeroOrOne"
				i++
				continue
			
			case '*':
				lastElement := last(last(stack))
				if (lastElement && lastElement["quantifier"] != "exactlyOne" ){
					err := errors.New("Quantifier must follow an unquantified element or group")
				}
				lastElement["quantifier"] = "zeroOrMore"
				i++
				continue
			case '+':
				lastElement := last(last(stack))
				if (lastElement && lastElement["quantifier"] != "exactlyOne" ){
					err := errors.New("Quantifier must follow an unquantified element or group")
				}
				zeroOrMoreCopy := {"type" : "wildcard", "quantifier" : "zeroOrMore"}
				i++
				continue
	


		}

	}
}

func main() {
	fmt.Println("hello world")
}

package main

import "fmt"

func checkBrackets(src string) bool {
	stack := NewStack()
	var symb = map[string]string{}

	symb["("] = ")"
	symb["{"] = "}"
	symb["["] = "]"

	for _, el := range src {
		str := string(el)

		if str == "{" || str == "(" || str == "[" {
			stack.Push(str)
		}

		if str == "}" || str == ")" || str == "]" {
			popped, ok := stack.Pop()

			if !ok {
				fmt.Println("not a pair 1")
				return false
			}

			if str != symb[popped.(string)] {
				fmt.Println(popped, str, "not a pair 2")
				return false
			} else {
				fmt.Println(popped, str, "pair 2")
			}
		}
	}

	if stack.GetSize() > 0 {
		return false
	}

	return true
}

func main() {
	/*
		()
		()(){}[]
		[][{}
		(
		{[][{}]}
	*/

	source := "{[][{}]}"
	fmt.Printf("string '%s' is balanced: %t", source, checkBrackets(source))
}

package main

import (
	"fmt"
)

func main() {
    testes := []string{
        "()",        // true
        "(())",      // true
        "(()",       // false
        "())",       // false
        "()()()",    // true
        "((())())",  // true
        "(()))(()",  // false
    }

    for _, t := range testes {
        fmt.Printf("%s -> %v\n", t, balparenteses(t))
    }
}

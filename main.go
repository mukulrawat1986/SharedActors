package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ActorNames []string

type stringReader interface {
	ReadString(byte) (string, error)
}

func AskForNames(r stringReader) {
	asking := true

	for asking {
		if len(ActorNames) < 2 {
			AskForName(r)
		} else {
			fmt.Println("Would you like to add another name?(y/n)")
			answer, _ := r.ReadString('\n')
			answer = strings.ToLower(strings.TrimSpace(answer))
			switch answer {
			case "y":
				AskForName(r)
			case "n":
				asking = false
			}
		}
	}
}

func AskForName(r stringReader) {
	fmt.Println("Please enter an actor's name")
	name, _ := r.ReadString('\n')
	name = strings.TrimSpace(name)
	ActorNames = append(ActorNames, name)
}

func main() {
	AskForNames(bufio.NewReader(os.Stdin))

	fmt.Printf("You selected the following %d actors:\n", len(ActorNames))
	for _, v := range ActorNames {
		fmt.Println(v)
	}
}

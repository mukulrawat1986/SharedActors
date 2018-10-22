package main_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/mukulrawat1986/SharedActors"
)

func setup() {
	main.ActorNames = []string{}
}

func TestAskForName(t *testing.T) {
	setup()

	a := assert.New(t)
	b := []byte("Mark\n")

	r := bytes.NewBuffer(b)

	main.AskForName(r)

	a.Equal(len(main.ActorNames), 1)
	a.Equal(main.ActorNames[0], "Mark")
}

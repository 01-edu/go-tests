package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
)

type node struct {
	flags          []string
	flagsShorthand []string
	randArgFlag    []string
	randArg        []string
}

func main() {
	s := []string{"--insert=", "--order"}
	strShorthand := []string{"-i=", "-o"}
	var randflag []string
	var randflagarg []string
	for i := 0; i < 2; i++ {
		randflagarg = append(randflagarg, rand.RandWords())
		randflag = append(randflag, rand.RandWords())
	}

	node := &node{
		flags:          s,
		flagsShorthand: strShorthand,
		randArgFlag:    randflagarg,
		randArg:        randflag,
	}

	node.randArg = append(node.randArg, "")

	challenge.Program("flags", node.flagsShorthand[0]+"v2", "v1")
	challenge.Program("flags", node.flagsShorthand[1], "v1")
	challenge.Program("flags", "-h")
	challenge.Program("flags", "--help")
	challenge.Program("flags")

	for _, v2 := range node.randArgFlag {
		for _, v1 := range node.randArg {
			challenge.Program("flags", node.flags[0]+v2, node.flags[1], v1)
		}
	}
	for _, v2 := range node.randArgFlag {
		for _, v1 := range node.randArg {
			challenge.Program("flags", node.flagsShorthand[0]+v2, node.flagsShorthand[1], v1)
		}
	}
}

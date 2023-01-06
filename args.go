package tokscan

import (
	"flag"
	"fmt"
	"os"
)

type Args struct {
	Dirs   []string
	Tokens []string
}

type ArrayArg []string

func (aa *ArrayArg) String() string {
	return fmt.Sprintf("%#v", *aa)
}

func (aa *ArrayArg) Set(v string) error {
	*aa = append(*aa, v)
	return nil
}

type Tokens []string

func Parse(args []string) (*Args, error) {
	flagset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	var dirs ArrayArg
	var tokens ArrayArg
	flagset.Var(&dirs, "d", "Specifying directories to be scanned")
	flagset.Var(&tokens, "t", "Specifying tokens to be scanned")

	if err := flagset.Parse(args); err != nil {
		return nil, err
	}

	return &Args{
		Dirs:   dirs,
		Tokens: tokens,
	}, nil
}

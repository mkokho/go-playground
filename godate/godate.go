package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"strconv"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "help" {
		fmt.Print("Examples: \n")
		fmt.Fprint(os.Stdout, strings.Join(examples(), "\n"))
		os.Exit(1)
	}

	res, err := run(os.Args[1:])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprint(os.Stdout, res)
	os.Exit(0)
}

func run(args []string) (string, error) {
	if len(args) == 0 {
		return time.Now().String(), nil
	} else if len(args) == 1 {
		nanos, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			loc, err := time.LoadLocation(args[0])
			if err != nil {
				return "", err
			}
			return time.Now().In(loc).String(), nil
		}
		return time.Unix(0, nanos).String(), nil
	}
	return "not yet implemetnted", nil
}

func examples() []string{
	supportedArgs := [][]string{
		[]string{},
		[]string{"US/Central"},
		[]string{"UTC"},
		[]string{"1481687107832160800"},
	}

	res := []string{}
	for _, args := range supportedArgs {
		x, err := run(args)
		if err != nil {
			x = err.Error()
		}
		res = append(res, fmt.Sprintf("> godate %s\n%s", strings.Join(args, " "), x))
	}
	return res
}

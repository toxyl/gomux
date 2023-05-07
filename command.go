package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/toxyl/glog"
)

type command struct {
	Name    string
	Usage   string
	Desc    string
	numArgs int
	run     func(arg ...string) error
}

func (ac *command) helpTextUsage() string {
	u := reToken.ReplaceAllStringFunc(ac.Usage, func(s string) string {
		return glog.Auto(s)
	})
	return glog.Bold() + APP_NAME + glog.Reset() + " " + glog.Underline() + glog.Auto(ac.Name) + glog.Reset() + " " + u
}

func (ac *command) helpTextUsagePadded(maxLen, maxLenUsage int) string {
	u := reToken.ReplaceAllStringFunc(ac.Usage, func(s string) string {
		return glog.Auto(s)
	})
	return glog.Bold() + APP_NAME + glog.Reset() + " " + glog.Underline() + glog.PadRight(glog.Auto(ac.Name), maxLen, ' ') + glog.Reset() + " " + glog.PadRight(u, maxLenUsage+1, ' ') + glog.WrapDarkGreen("// "+glog.StripANSI(ac.Desc))
}

func (ac *command) helpText() string {
	return fmt.Sprintf(
		"%s\n\n",
		glog.Bold()+"Usage: "+glog.Reset()+ac.helpTextUsage(),
	)
}

func (ac *command) is(op string) bool {
	return ac.Name == op
}

func (ac *command) Run() error {
	numArgsExpected := ac.numArgs
	numArgs := len(os.Args) - 2
	hasTooFewArgs := numArgsExpected > -1 && numArgs < numArgsExpected
	hasTooManyArgs := numArgsExpected > -1 && numArgs > numArgsExpected

	if hasTooFewArgs {
		fmt.Print(ac.helpText())
		fmt.Println(glog.WrapOrange("Not enough arguments!"))
		os.Exit(EXIT_MISSING_ARGS)
	}

	if hasTooManyArgs {
		fmt.Print(ac.helpText())
		fmt.Println(glog.WrapOrange("Too many arguments!"))
		os.Exit(EXIT_TOO_MANY_ARGS)
	}

	args := os.Args[2:]

	return ac.run(args...)
}

type ArgList []string

func (ail *ArgList) String() string {
	res := []string{}
	for _, ai := range *ail {
		res = append(res, fmt.Sprintf("[%s]", ai))
	}
	return strings.Join(res, " ")
}

func newCommand(
	name string,
	desc string,
	args ArgList,
	run func(arg ...string) error,
) *command {
	ac := &command{
		Name:    name,
		Usage:   args.String(),
		Desc:    desc,
		numArgs: len(args),
		run:     run,
	}

	return ac
}

func RegisterCommand(
	name string,
	desc string,
	args ArgList,
	run func(arg ...string) error,
) {
	cmdReg = append(cmdReg, newCommand(
		name,
		desc,
		args,
		run,
	))
}

func FindCommand(op string) *command {
	for _, cmd := range cmdReg {
		if cmd.is(op) {
			return cmd
		}
	}
	return nil
}

func ListCommands() {
	maxLenName := getMaxCommandNameLength()
	maxLenUsage := getMaxCommandUsageLength()
	for _, cmd := range cmdReg {
		fmt.Println("  " + cmd.helpTextUsagePadded(maxLenName, maxLenUsage))
	}
}

func getMaxCommandNameLength() int {
	maxLen := 0
	for _, cmd := range cmdReg {
		maxLen = glog.Max(maxLen, len(cmd.Name))
	}

	return maxLen
}

func getMaxCommandUsageLength() int {
	maxLen := 0
	for _, cmd := range cmdReg {
		maxLen = glog.Max(maxLen, len(cmd.Usage))
	}

	return maxLen
}

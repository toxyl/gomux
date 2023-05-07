package main

import (
	"fmt"
	"os"

	"github.com/toxyl/glog"
)

func main() {
	glog.LoggerConfig.ShowSubsystem = false
	glog.LoggerConfig.ShowDateTime = false
	glog.LoggerConfig.ShowRuntimeMilliseconds = false
	glog.LoggerConfig.ShowRuntimeSeconds = false

	RegisterCommand(
		COMMAND_START,
		"Starts a tmux session using the given config file.",
		ArgList{"config"},
		Start,
	)

	RegisterCommand(
		COMMAND_DAEMON,
		"Starts a tmux session in the background using the given config file.",
		ArgList{"config"},
		Daemon,
	)

	RegisterCommand(
		COMMAND_DETACH,
		"Detaches all clients connected to the session started with the given config file.",
		ArgList{"config"},
		Detach,
	)

	RegisterCommand(
		COMMAND_LIST,
		"Lists all active sessions.",
		ArgList{},
		List,
	)

	if len(os.Args) < 2 {
		fmt.Println("\n" +
			glog.Bold() + glog.WrapYellow("Welcome to "+APP_NAME+"!") + glog.Reset() +
			"\n" +
			"\n" +
			glog.Underline() + "Available Commands" + glog.Reset() +
			"\n")
		ListCommands()
		fmt.Println()
		os.Exit(EXIT_OK)
	}
	cmd := FindCommand(os.Args[1])
	if cmd != nil {
		err := cmd.Run()
		if err != nil {
			fmt.Println("\n" +
				glog.Bold() + glog.WrapDarkRed("ERROR") + glog.Reset() +
				"\n" +
				glog.Error(err) +
				"\n",
			)
		}
	}
}

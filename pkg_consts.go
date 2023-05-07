package main

const (
	EXIT_OK = iota
	EXIT_MISSING_ARGS
	EXIT_TOO_MANY_ARGS
)

const (
	APP_NAME = "gomux"

	COMMAND_START  = "start"
	COMMAND_DAEMON = "daemon"
	COMMAND_DETACH = "detach"
	COMMAND_LIST   = "list"
)

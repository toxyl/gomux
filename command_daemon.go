package main

import (
	"fmt"
)

func Daemon(arg ...string) error {
	file := arg[0]

	w, err := LoadConfig(file)
	if err != nil {
		return err
	}
	panes := make([]string, len(w.Panes))
	for i, pane := range w.Panes {
		arg := ""
		if pane.Delay > 0 {
			arg = "sleep " + fmt.Sprint(pane.Delay) + " ; "
		}
		panes[i] = arg + pane.Command + " ; /bin/bash" // to keep the window open
	}

	return SpawnWorkspace(w.Name, false, panes...)
}

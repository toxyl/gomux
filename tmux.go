package main

type TmuxWorkspace struct {
	created bool
	name    string
}

func (tw *TmuxWorkspace) add(cmd string) error {
	if !tw.created {
		if err := ExecNoOut("new-session", "-s", tw.name, "-d", cmd); err != nil {
			return err
		}
		tw.created = true
		return nil
	}
	return ExecNoOut("split-window", "-v", cmd)
}

func (tw *TmuxWorkspace) spawn(attach bool) error {
	if err := ExecNoOut("select-layout", "tiled"); err != nil {
		return err
	}
	if err := ExecNoOut("set-option", "-g", "mouse", "on"); err != nil {
		return err
	}
	if attach {
		if err := ExecNoOut("attach-session"); err != nil {
			return err
		}
	}
	return nil
}

func newTmuxWorkspace(name string) *TmuxWorkspace {
	tw := &TmuxWorkspace{
		name:    APP_NAME + "-" + name,
		created: false,
	}
	return tw
}

func SpawnWorkspace(name string, attach bool, commands ...string) error {
	if HasSession(name) {
		// we already have a session
		if !attach {
			// but we don't want to attach, let's quietly ignore this
			return nil
		}
		// let's attach
		return ExecNoOut("attach-session", "-t", APP_NAME+"-"+name)
	}

	tw := newTmuxWorkspace(name)

	for _, cmd := range commands {
		if err := tw.add(cmd); err != nil {
			return err
		}
	}

	return tw.spawn(attach)
}

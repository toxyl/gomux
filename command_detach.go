package main

func Detach(arg ...string) error {
	w, err := LoadConfig(arg[0])
	if err != nil {
		return err
	}
	return DetachClientsFromSession(w.Name)
}

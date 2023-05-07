package main

import (
	"fmt"
	"strings"

	"github.com/toxyl/glog"
)

func List(arg ...string) error {
	clientsPerSession := map[string][]string{}
	tmuxSessions, _ := GetSessions()
	tmuxClients, _ := GetClients()
	for _, session := range tmuxSessions {
		session = strings.ReplaceAll(session, APP_NAME+"-", "")
		clientsPerSession[session] = []string{}
	}
	for session, clients := range tmuxClients {
		session = strings.ReplaceAll(session, APP_NAME+"-", "")
		clientsPerSession[session] = clients
	}

	for session, clients := range clientsPerSession {
		fmt.Printf("%s (%s)\n", glog.Auto(session), glog.IntAmount(len(clients), "client", "clients"))
		for _, client := range clients {
			fmt.Printf("- %s\n", glog.Highlight(client))
		}
	}
	return nil
}

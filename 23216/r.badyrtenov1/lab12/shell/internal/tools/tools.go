package tools

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

type Job struct {
	Pid    int
	Status string
	Cmd    string
}

const (
	OUTPIP = 01
	INPIP  = 02
)

func RemoveByte(slice []byte, beginId int, endId int) []byte {
	newSlice := make([]byte, len(slice)+beginId-endId)
	if beginId != endId {
		slice = append(slice[:beginId], slice[endId:]...)
	}
	copy(newSlice, slice)
	return newSlice
}

func RemoveJob(slice []Job, beginId int, endId int) []Job {
	newSlice := make([]Job, len(slice)+beginId-endId)
	if beginId != endId {
		slice = append(slice[:beginId], slice[endId:]...)
	}
	copy(newSlice, slice)
	return newSlice
}

func Promptline() error {
	customer, err := user.Current()
	if err != nil {
		customer = &user.User{Username: "username"}
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "hostname"
	}

	cwd, err := os.Getwd()
	if err != nil {
		cwd = "~"
	} else {
		homeDir := os.Getenv("HOME")
		if strings.HasPrefix(cwd, homeDir) {
			cwd = strings.Replace(cwd, homeDir, "~", 1)
		}
	}
	_, err = fmt.Printf("%s@%s:%s$ ", customer.Username, hostname, cwd)
	return err
}

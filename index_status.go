package main

import "sync"

type (
	StatusDict struct {
		sync.RWMutex
		S []string // status value
	}
)

func MakeStatusDict() *StatusDict {
	return &StatusDict{S: []string{}}
}

func (sd *StatusDict) GetId(status string) int {
	for i, s := range sd.S {
		if s == status {
			return i
		}
	}
	return -1
}

func (sd *StatusDict) Add(status string) int {

	if i := sd.GetId(status); i > -1 {
		return i
	}

	sd.S = append(sd.S, status)
	return len(sd.S) - 1
}

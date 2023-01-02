package test

import (
	"errors"
	"github.com/nskforward/fix"
	"os"
)

func session() (*fix.Session, error) {
	addr := os.Getenv("FIX_ADDR")
	if addr == "" {
		return nil, errors.New("env FIX_ADDR is empty")
	}

	sender := os.Getenv("FIX_SENDER")
	if sender == "" {
		return nil, errors.New("env FIX_SENDER is empty")
	}

	target := os.Getenv("FIX_TARGET")
	if target == "" {
		return nil, errors.New("env FIX_TARGET is empty")
	}

	pass := os.Getenv("FIX_PASS")
	if pass == "" {
		return nil, errors.New("env FIX_PASS is empty")
	}

	return fix.NewSession(addr, sender, target, pass)
}

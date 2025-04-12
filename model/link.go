package model

import (
	"errors"
	"strings"
)

type Link struct {
	Data string `json:"data"`
}

func (l *Link) Validate() error {
	if l.Data == "" {
		return errors.New("data is required")
	}

	if !strings.HasPrefix(l.Data, "https://") && !strings.HasPrefix(l.Data, "http://") {
		return errors.New("data must start with https:// or http://")
	}

	if len(l.Data) > 20480 {
		return errors.New("data must be less than 20480 characters")
	}

	return nil
}

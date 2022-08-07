package model

import (
	"encoding/json"
	"fmt"
	"io"
)

type Polling []Game

func NewGamePolling(rdr io.Reader) (Polling, error) {
	var polling Polling

	err := json.NewDecoder(rdr).Decode(&polling)
	if err != nil {
		err = fmt.Errorf("problem parsing polling: %v", err)
	}

	return polling, err
}

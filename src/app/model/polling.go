package model

import (
	"encoding/json"
	"fmt"
	"io"
)

func NewGamePolling(rdr io.Reader) ([]Game, error) {
	var polling []Game

	err := json.NewDecoder(rdr).Decode(&polling)
	if err != nil {
		err = fmt.Errorf("problem parsing polling: %v", err)
	}

	return polling, err
}

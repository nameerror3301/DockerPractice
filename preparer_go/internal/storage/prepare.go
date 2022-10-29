package storage

import (
	"fmt"
	"sync"
)

type storage interface {
	SendDataOnPrep(data []byte) error
}

var (
	store storage
	call  sync.Once = sync.Once{}
)

func SetPrep(set storage) {
	call.Do(func() {
		store = set
	})
}

// Buisnes logic
func SendPrepData(data []byte) error {
	if err := store.SendDataOnPrep(data); err != nil {
		return fmt.Errorf("SendPrepData error: %v", err)
	}

	return nil
}

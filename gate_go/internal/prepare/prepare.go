package prepare

import (
	"github.com/sirupsen/logrus"
)

type preparer interface {
	SendDataOnPrep(data []byte) error
}

var prep preparer

func SetPrep(set preparer) {
	prep = set
}

// Buisnes logic
func SendPrepData(data []byte) {
	if err := prep.SendDataOnPrep(data); err != nil {
		logrus.Error("Write error: %v", err)
	}
}

package cmd

import (
	// external
	"github.com/sirupsen/logrus"
)

const (
	VERSION = "1.1.0"
)

func runVersion() {
	logrus.Printf("Glod-cli's version %s", VERSION)
}

// Package to configure options that should be acessible througout all commands
package lsh

import (
	"log"

	sdk "github.com/latitudesh/latitudesh-go"
	"github.com/spf13/viper"
)

// Dry run flag
var DryRun bool

// Debug flag indicating that cli should output debug logs
var Debug bool

// LogDebugf writes debug log to stdout
func LogDebugf(format string, v ...interface{}) {
	if !Debug {
		return
	}
	log.Printf(format, v...)
}

func NewClient() *sdk.Client {
	AuthorizationKey := viper.GetString("Authorization")

	c := sdk.NewClientWithAuth("latitudesh", " ", nil)

	if AuthorizationKey != "" {
		c = sdk.NewClientWithAuth("latitudesh", AuthorizationKey, nil)
	}

	return c
}

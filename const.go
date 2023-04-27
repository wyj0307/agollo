package agollo

import (
	"time"
)

const (
	defaultConfName  = "agollo.json" // by xingdonghai edited, ori: app.properties
	defaultNamespace = "application"

	longPoolInterval      = time.Second * 2
	longPoolTimeout       = time.Second * 90
	queryTimeout          = time.Second * 2
	defaultNotificationID = -1

	errMissENV = "environment variable not set" // by xingdonghai
	errMissCLI = "cli arguments not set" // by xingdonghai
	defaultTagName = "config" // by xingdonghai
)

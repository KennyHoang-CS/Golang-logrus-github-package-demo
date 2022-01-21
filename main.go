package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetFormatter(&log.JSONFormatter{}) // formats the log as JSON
	//	log.SetReportCaller(true)              // determines which method the log is calling from.
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel) // only log when the severity is info or above.

	// A common pattern is to re-use fields between logging statements by
	// reusing the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "this is other",
	})

	contextLogger.Info("I'll be logged with common and other fields")
	contextLogger.Info("Same")

	log.WithFields(log.Fields{
		"animal": "cat",
	}).Info("A cat appears.")

	log.WithFields(log.Fields{
		"animal": "penguins",
	}).Warn("Penguins can be deadly if provoked.")

	log.WithFields(log.Fields{
		"animal": "bear",
	}).Fatal("Avoid Bears!!!")

}

package main

import (
	"flag"
	"io/ioutil"
	"strconv"

	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
)

var version = "undefined"

func main() {

	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	logrus.SetFormatter(formatter)

	a := app{}
	flag.IntVar(&a.Port, "port", 8000, "Application port")
	var config = flag.String("config", "config.yaml", "Config file path")
	flag.Parse()

	file, err := ioutil.ReadFile(*config)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := yaml.Unmarshal(file, &a); err != nil {
		logrus.Fatal(err)
	}

	if err := a.Initialize(); err != nil {
		logrus.Fatal(err)
	}

	logrus.WithFields(logrus.Fields{
		"redis":   a.Redis.Hostname + ":" + strconv.Itoa(a.Redis.Port),
		"version": version,
	}).Info("Countme is running on :", a.Port)

	if err := a.Run(); err != nil {
		logrus.Fatal(err)
	}
}

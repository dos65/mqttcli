package main

import (
	"bufio"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func publish(c *cli.Context) {
	setDebugLevel(c)

	opts, err := NewOption(c)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	client, err := connect(c, opts, map[string]byte{})
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	qos := c.Int("q")
	topic := c.String("t")
	if topic == "" {
		log.Errorf("Please specify topic")
		os.Exit(1)
	}
	log.Infof("Topic: %s", topic)

	retain := c.Bool("r")

	if c.Bool("s") {
		// Read from Stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			err = client.Publish(topic, []byte(scanner.Text()), qos, retain, true)
			if err != nil {
				log.Error(err)
			}

		}
	} else {
		payload := c.String("m")
		err = client.Publish(topic, []byte(payload), qos, retain, true)
		if err != nil {
			log.Error(err)
		}

	}
	log.Info("Published")
	err = client.Disconnect()
	if err != nil {
		log.Errorf("disconnect error: %s", err)
	}

}

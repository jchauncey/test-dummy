package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config is used to parse out runtime config using env vars
type Config struct {
	FailureRate int           `envconfig:"FAILURE_RATE" default:"50"`
	Sleep       time.Duration `envconfig:"SLEEP" default:"60s"`
}

func main() {
	var config Config
	err := envconfig.Process("test-dummy", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Config:%+v", config)

	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	for {
		n := rnd.Intn(100)
		if n > config.FailureRate {
			log.Printf("Exceeded failure rate exiting (%v)...", n)
			os.Exit(1)
		}
		time.Sleep(config.Sleep)
	}

}

package xxx

import (
	"log"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	if err := Run(); err != nil {
		log.Printf("ERROR : %s", err.Error())
	}

	ticker := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-ticker.C:
			if err := Run(); err != nil {
				log.Printf("ERROR : %s", err.Error())
			}
		}
	}
}

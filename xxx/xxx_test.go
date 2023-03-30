package xxx

import (
	"log"
	"testing"
	"time"
)

func TestXXX(t *testing.T) {
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

package proxydp

import (
	"log"
	"testing"
)

func TestAPI(t *testing.T) {
	err := (&AssembleGit{}).API()
	if err != nil {
		log.Printf("New Testflight Success!")
	} else {
		log.Printf("New Testflight Failed!")
	}
}

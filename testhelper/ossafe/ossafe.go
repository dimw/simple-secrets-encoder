package ossafe

import (
	"log"
	"os"
)

func Remove(filename string) {
	err := os.Remove(filename)
	if err != nil {
		log.Printf(`File "%v" could not be deleted.`, filename)
	}
}

func RemoveAll(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		log.Printf(`Path "%v" could not be deleted.`, path)
	}
}

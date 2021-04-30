package variables

import (
	"encoding/json"
	"log"
	"os"
)

var Config Configuration

func init() {
	log.Println("Reading config file...")

	file, err := os.Open("conf.json")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)

	if err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"fmt"
	"math/rand"
	"os"
	"text/template"
	"time"

	"github.com/google/uuid"
)

const configTemplate = `
export DOMINION_ID="{{.ID}}"
export DOMINION_PORT="{{.Port}}"
`

type MinimumConfig struct {
	ID   uuid.UUID
	Port int
}

func main() {
	var fileName string
	if len(os.Args) < 2 {
		panic(fmt.Errorf("$DOMINION_CONFIG not specified"))
	}
	fileName = os.Args[1]

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tpl, err := template.New("config").Parse(configTemplate)
	if err != nil {
		panic(err)
	}

	mc := MinimumConfig{
		ID:   uuid.New(),
		Port: getRandomPort(),
	}

	tpl.Execute(file, mc)
}

// getRandomPort returns a random port based on a seed
func getRandomPort() int {
	rand.Seed(time.Now().UnixNano())
	uint16Max := (1 << 16) - 1
	i := rand.Intn(uint16Max)
	if i < 1024 || 49151 < i { // stay in registered port numbers
		// hacky, I know. It's script code, what can I say :shrug:
		return getRandomPort()
	}
	return i
}

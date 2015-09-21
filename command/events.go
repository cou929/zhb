package command

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
)

func CmdEvents(c *cli.Context) {
	client := NewZenhubClient(c)

	events, err := client.GetEvents(c.Int("page"))
	DieIfError(err)

	eventsInBytes, err := json.Marshal(events)
	DieIfError(err)

	fmt.Println(string(eventsInBytes))
}

package command

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
)

func CmdPipelines(c *cli.Context) {
	client := NewZenhubClient(c)

	pipelines, err := client.GetPipelines(0)
	DieIfError(err)

	result := make([]map[string]interface{}, len(pipelines))
	for i, v := range pipelines {
		result[i] = map[string]interface{}{
			"Id":   v.Id,
			"Name": v.Name,
		}
	}

	marshaled, err := json.Marshal(result)
	DieIfError(err)

	fmt.Println(string(marshaled))
}

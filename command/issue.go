package command

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	zenhub "github.com/cou929/zenhub-client"
	"strconv"
)

type Estimate struct {
	Value int
}

type Pipeline struct {
	Id   string
	Name string
}

type Plus struct {
	Id           string
	Comment      int
	ZenhubUserId string
	GithubUserId int
	UserName     string
	CreatedAt    string
}

func CmdIssue(c *cli.Context) {
	issueNumber, err := strconv.Atoi(c.Args().Get(0))
	if err != nil {
		cli.ShowCommandHelp(c, "issue")
	}

	client := NewZenhubClient(c)

	estimate, err := findEstimate(client, issueNumber)
	if err != nil {
		Log(err.Error())
	}

	pipeline, err := findPipeline(client, issueNumber)
	if err != nil {
		Log(err.Error())
	}

	pluses := make([]Plus, 0)
	repoId, err := strconv.Atoi(c.GlobalString("repo_id"))
	if err == nil {
		res, err := findPluses(client, issueNumber, repoId)
		if err != nil {
			Log(err.Error())
		}
		pluses = *res
	}

	result := map[string]interface{}{
		"issueNumber":   issueNumber,
		"pipelineId":    pipeline.Id,
		"pipelineName":  pipeline.Name,
		"estimateValue": estimate.Value,
		"pluses":        pluses,
	}

	marshaled, err := json.Marshal(result)
	DieIfError(err)

	fmt.Println(string(marshaled))
}

func findEstimate(client *zenhub.Client, issueNumber int) (*Estimate, error) {
	estimates, err := client.GetEstimates(issueNumber)
	if err != nil {
		return nil, err
	}

	for _, v := range estimates {
		if v.Selected {
			return &Estimate{Value: v.Value}, nil
		}
	}

	return nil, nil
}

func findPipeline(client *zenhub.Client, issueNumber int) (*Pipeline, error) {
	pipelines, err := client.GetPipelines(issueNumber)
	if err != nil {
		return nil, err
	}

	for _, v := range pipelines {
		if v.IsIn {
			return &Pipeline{Id: v.Id, Name: v.Name}, nil
		}
	}

	return nil, nil
}

func findPluses(client *zenhub.Client, issueNmber int, repoId int) (*[]Plus, error) {
	pluses, err := client.GetPluses(issueNmber, repoId)
	if err != nil {
		return nil, err
	}

	res := make([]Plus, len(pluses), len(pluses))
	for i, v := range pluses {
		p := Plus{
			Id:           v.Id,
			Comment:      v.Comment,
			ZenhubUserId: v.User.Id,
			GithubUserId: v.User.GithubUserInfo.Id,
			UserName:     v.User.GithubUserInfo.Name,
			CreatedAt:    v.CreatedAt,
		}
		res[i] = p
	}

	return &res, nil
}

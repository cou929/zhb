# zhb

[![wercker status](https://app.wercker.com/status/4c192feadf7d5c5b6208f7d7fad70407/s/master "wercker status")](https://app.wercker.com/project/bykey/4c192feadf7d5c5b6208f7d7fad70407)

Unofficial cli tool for ZenHub.io

## Description

`zhb` is a command-line interface tool for [ZenHub.io](http://zenhub.io/). With `zhb`, you can get, update or clear the parameters such as pipelines, estimates or `+1` of the issue.

### Attention

ZenHub has no public API at the moment. `zhb` uses internal APIs of the ZenHub chrome extensions arbitrarily. So interface of the commands would be changed.

## Installation

```sh
$ go get github.com/cou929/zhb
$ go install github.com/cou929/zhb
```

## Usage

```
zhb issue <issue_number>
zhb transfer <issue_number> <pipeline_name>
zhb estimate <issue_number> <estimate_value>
zhb clearEstimate <issue_number>
zhb pipelines
zhb events [--page=<page_number>]
```

### EXAMPLES

#### `issue`

Show details of the issue. The example below gets information of issue `1` and uses [`jq`](https://stedolan.github.io/jq/) command to handle json response. `zhb` find and show estimates, pipelines and `+1` of the issue.

```sh
$ zhb issue 1 | jq '.'
{
  "estimateValue": 5,
  "issueNumber": 1,
  "pipelineId": "1234567890",
  "pipelineName": "New Issue",
  "pluses": [
    {
      "Id": "1234567890",
      "Comment": -1,       # -1 means the `+1` is for entire issue, not for specific comment
      "ZenhubUserId": "1234567890",
      "GithubUserId": 1234567890,
      "UserName": "user_name",
      "CreatedAt": "2015-09-01T00:00:00.000Z"
    }
  ]
}
```

#### `transfer`

Transfer the issue to the pipeline. The example below transfers issue `1` to `To Do` pipeline. `zhb` exits with code `0` if the execution finished successfully.

```sh
$ zhb transfer 1 "To Do"
```

### `estimate`

Set the estimate value to the issue. The example below set estimate value `5` to issue `1`. `zhb` exits with code `0` if the execution finished successfully.

```sh
$ zhb estimate 1 5
```

### `clearEstimate`

Clear estimate value of the issue. The example below clears estiamte value of the issue `1`. `zhb` exits with code `0` if the execution finished successfully.

```sh
$ zhb clearEstimate 1
```

### `pipelines`

Show all pipelines. The example below uses [`jq`](https://stedolan.github.io/jq/) command to handle json response.

```sh
$ zhb pipelines | jq '.'
[
  {
    "Id": "1234567890",
    "Name": "New Issues"
  },
  {
    "Id": "1234567890",
    "Name": "To Do"
  },
  {
    "Id": "1234567890",
    "Name": "Done"
  }
]
```

### `events`

Show events of ZenHub activities. The example below gets page `2` of events and uses [`jq`](https://stedolan.github.io/jq/) command to handle json response. `--page` option is optional and default value is `1`.

```sh
$ zhb events --page=2
[
  {
    "id": "1234567890",
    "actor": {
      "id": "1234567890",
      "github": {
        "id": 1234,
        "username": "user_name",
        "avatarUrl": "https://avatars.githubusercontent.com/u/000?v=3"
      }
    },
    "type": "transferIssue",
    "repoId": 1234,
    "organization": "org",
    "repository": "repo",
    "issue": 123,
    "comment": 0,
    "recipient": 0,
    "createdAt": "2015-01-01T00:00:00.000Z",
    "srcPipelineName": "New Issues",
    "destPipelineName": "To Do"
  },
  ...
]
```

## Environment Variables

```sh
export ZHB_AUTH_TOKEN=<Your zenhub auth token>
export ZHB_ORG=<Your github organization name>
export ZHB_REPO=<Your github repository name>
export ZHB_REPO_ID=<Your github repository name>
```

You can overwrite there variables through command line options.

```sh
zhb --org="another org name" [command]
```

### `ZHB_AUTH_TOKEN`

At the moment `zhb` has no way to authorize ZenHub account. So you should get your `auth token` from browser session which is authorized.

You can find out `auth token` by below steps:

1. Open the developer tool.
1. Open the `Network` section.
1. Add fileter `zenhub` to eliminate all other requests.
1. Open your repository root page on GitHub.
1. Open the `access` request.
1. You can see `x-authentication-token` request header. The value of it is `auth token`. Just copy it.

![](https://raw.githubusercontent.com/wiki/cou929/zhb/images/token-header.png)

### `ZHB_REPO_ID`

Repository Id of GitHub is needed by `+1` APIs of ZenHub. So you must set the variable to handle `+1`. You can omit this variable if you do not need `+1` information. `zhb` just ignore `+1` APIs if `ZHB_REPO_ID` is not presented.

You can find out `Repository Id` by [GitHub Repository API](https://developer.github.com/v3/repos/). Below is an example to get ids and names of repositories by using `curl` and [`jq`](https://stedolan.github.io/jq/) command.

```sh
curl -H "Authorization: token <your github auth token>" https://api.github.com/orgs/<your_org>/repos | jq '.[] | {id, name}'
```

## Contribution

1. Fork ([https://github.com/cou929/zhb/fork](https://github.com/cou929/zhb/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[Kosei Moriyama](https://github.com/cou929)

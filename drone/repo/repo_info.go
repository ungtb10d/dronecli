package repo

import (
	"os"
	"text/template"

	"github.com/drone/drone-cli/drone/internal"
	"github.com/drone/funcmap"
	"github.com/urfave/cli"
)

var repoInfoCmd = cli.Command{
	Name:      "info",
	Usage:     "show repository details",
	ArgsUsage: "<repo/name>",
	Action:    repoInfo,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "format",
			Usage: "format output",
			Value: tmplRepoInfo,
		},
	},
}

func repoInfo(c *cli.Context) error {
	arg := c.Args().First()
	owner, name, err := internal.ParseRepo(arg)
	if err != nil {
		return err
	}

	client, err := internal.NewClient(c)
	if err != nil {
		return err
	}

	repo, err := client.Repo(owner, name)
	if err != nil {
		return err
	}

	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.String("format"))
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, repo)
}

// template for repo information
var tmplRepoInfo = `Owner: {{ .Namespace }}
Repo: {{ .Name }}
Config: {{ .Config }}
Visibility: {{ .Visibility }}
Private: {{ .Private }}
Trusted: {{ .Trusted }}
Protected: {{ .Protected }}
Remote: {{ .HTTPURL }}
CancelRunning: {{ .CancelRunning }}
CancelPulls: {{ .CancelPulls }}
CancelPush: {{ .CancelPush }}
`

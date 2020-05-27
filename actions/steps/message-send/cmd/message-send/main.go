package main

import (
	"flag"

	"github.com/puppetlabs/nebula-sdk/pkg/log"
	"github.com/puppetlabs/nebula-sdk/pkg/taskutil"
	"github.com/slack-go/slack"
)

type ConnectionSpec struct {
	APIToken string `spec:"apiToken"`
}

type Spec struct {
	// New-form API Token.
	Connection *ConnectionSpec

	// Old-form API token
	APIToken string `spec:"apitoken"`

	Channel  string
	Message  string
	Username string
}

func (s Spec) ConfiguredAPIToken() (string, bool) {
	if s.Connection != nil {
		return s.Connection.APIToken, s.Connection.APIToken != ""
	}

	return s.APIToken, s.APIToken != ""
}

func main() {
	u, err := taskutil.MetadataSpecURL()
	if err != nil {
		log.FatalE(err)
	}
	specURL := flag.String("spec-url", u, "url to fetch the spec from")

	flag.Parse()

	planOpts := taskutil.DefaultPlanOptions{SpecURL: *specURL}

	var spec Spec
	if err := taskutil.PopulateSpecFromDefaultPlan(&spec, planOpts); err != nil {
		log.FatalE(err)
	}

	tok, found := spec.ConfiguredAPIToken()
	if !found {
		log.Fatal("specify the Slack connection to use")
	} else if spec.Message == "" {
		log.Fatal("specify the message to send to Slack")
	} else if spec.Channel == "" {
		log.Fatal("specify the channel to send the message to")
	}

	if spec.Username == "" {
		spec.Username = "Relay by Puppet"
	}

	api := slack.New(tok)
	_, _, err = api.PostMessage(spec.Channel, slack.MsgOptionText(spec.Message, false), slack.MsgOptionUsername(spec.Username))
	if err != nil {
		log.FatalE(err)
	}
}

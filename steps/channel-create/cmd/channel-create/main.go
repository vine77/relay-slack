package main

import (
	"flag"

	"github.com/puppetlabs/relay-sdk-go/pkg/log"
	"github.com/puppetlabs/relay-sdk-go/pkg/outputs"
	"github.com/puppetlabs/relay-sdk-go/pkg/taskutil"
	"github.com/slack-go/slack"
)

type ConnectionSpec struct {
	APIToken string `spec:"apiToken"`
}

type Spec struct {
	// New-form API Token.
	Connection *ConnectionSpec

	Channel  string
	Topic    string
	MemberID string
}

// This uses the new conversations api as the old channels api is deprecated
// https://api.slack.com/changelog/2020-01-deprecating-antecedents-to-the-conversations-api
func main() {
	var (
		specURL = flag.String("spec-url", mustGetDefaultMetadataSpecURL(), "url to fetch the spec from")
	)

	flag.Parse()

	spec := mustPopulateSpec(*specURL)

	if spec.Connection == nil {
		log.Fatal("specify the Slack connection to use")
	} else if spec.Connection.APIToken == "" {
		log.Fatal("the specified connection must be a Slack connection")
	} else if spec.Channel == "" {
		log.Fatal("specify the channel to create")
	}

	log.Info("connecting to slack api...")
	api := slack.New(spec.Connection.APIToken)
	log.Info("connected!")
	log.Info("creating channel...")
	conv, err := api.CreateConversation(spec.Channel, false)
	if err != nil {
		log.FatalE(err)
	}
	log.Info("channel was created!")
	if spec.Topic != "" {
		log.Info("setting topic...")
		_, err = api.SetTopicOfConversation(conv.ID, spec.Topic)
		if err != nil {
			log.FatalE(err)
		}
		log.Info("topic set!")
	}
	if spec.MemberID != "" {
		log.Info("inviting member...")
		_, err = api.InviteUsersToConversation(conv.ID, spec.MemberID)
		if err != nil {
			log.FatalE(err)
		}
		log.Info("member invited!")
	}
	if client, err := outputs.NewDefaultOutputsClientFromNebulaEnv(); err != nil {
		log.FatalE(err)
	} else {
		if err := client.SetOutput(context.Background(), "channelId", ch.ID); err != nil {
			log.FatalE(err)
		}
	}
}

// This just encapsulates some setup logic to clean up the main function a bit.
func mustGetDefaultMetadataSpecURL() string {
	if metadataSpecURL, err := taskutil.MetadataSpecURL(); err != nil {
		log.FatalE(err)

		// control should exit before we get to here (thanks to the fatal above).
		// this just makes the compiler shut up.
		panic(err)
	} else {
		return metadataSpecURL
	}
}

func mustPopulateSpec(specURL string) (spec Spec) {
	opts := taskutil.DefaultPlanOptions{SpecURL: specURL}

	if err := taskutil.PopulateSpecFromDefaultPlan(&spec, opts); err != nil {
		log.FatalE(err)
	}

	return
}

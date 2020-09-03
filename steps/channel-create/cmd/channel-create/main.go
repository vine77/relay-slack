package main

import (
	"flag"

	"github.com/puppetlabs/relay-sdk-go/pkg/log"
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
}

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

	api := slack.New(spec.Connection.APIToken)
	_, err := api.CreateChannel(spec.Channel)
	if err != nil {
		log.FatalE(err)
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

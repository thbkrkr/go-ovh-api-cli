package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/ovh/go-ovh/ovh"
)

const (
	HELP = `
 Usage: ovhapi METHOD PATH

 Call the OVH API.

 Methods: GET, POST, PUT, DELETE
 Path: see https://api.ovh.com/console/
`
)

// Config for the OVH client
type Config struct {
	Endpoint          string `envconfig:"endpoint" default:"ovh-eu"`
	ApplicationKey    string `envconfig:"ak" required:"true"`
	ApplicationSecret string `envconfig:"as" required:"true"`
	ConsumerKey       string `envconfig:"ck" required:"true"`
}

var (
	method string
	path   string

	config Config
)

func init() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, HELP)
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	// Parse arguments
	if flag.NArg() != 2 {
		exit(fmt.Sprint("ovhapi requires a minimum of 2 arguments."))
	} else {
		method = flag.Args()[0]
		path = flag.Args()[1]
	}

	// Try to read stdin
	stat, err := os.Stdin.Stat()
	if err != nil {
		exit(fmt.Sprintf("Fail to read stdin:\n %v\n", err))
	}
	var data interface{}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		reader := bufio.NewReader(os.Stdin)
		rawData, err := ioutil.ReadAll(reader)
		if err != nil {
			exit(fmt.Sprintf("Fail to read stdin:\n %v\n", err))
		}

		if len(rawData) > 0 {
			err = json.Unmarshal(rawData, &data)
			if err != nil {
				exit(fmt.Sprintf("Cannot parse JSON body:\n %v\n", err))
			}
		}
	}

	// Load config and create Ovh client
	err = envconfig.Process("OVH", &config)
	if err != nil {
		exit(fmt.Sprintf("Fail to process env config:\n %v\n", err))
	}
	ovhClient, err := ovh.NewClient(
		config.Endpoint,
		config.ApplicationKey,
		config.ApplicationSecret,
		config.ConsumerKey,
	)
	if err != nil {
		exit(fmt.Sprintf("Bad configuration:\n %v\n", err))
	}

	// Call OVH API
	var response interface{}
	err = ovhClient.CallAPI(method, path, data, &response, true)
	if err != nil {
		switch err.(type) {
		case *ovh.APIError:
			apiErr := err.(*ovh.APIError)
			fmt.Printf("Error: %d on %s %s\n%s\n", apiErr.Code, method, path, apiErr.Message)
		default:
			fmt.Printf("Error: on %s %s\n%s\n", method, path, err.Error())
		}
		os.Exit(1)
	}

	bytes, err := json.Marshal(response)
	fmt.Println(string(bytes))
}

func exit(msg string) {
	fmt.Fprintf(os.Stderr, msg)
	os.Exit(1)
}

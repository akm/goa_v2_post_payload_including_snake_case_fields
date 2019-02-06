// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// examples HTTP client CLI support package
//
// Command:
// $ goa gen
// github.com/akm/goa_v2_post_payload_including_snake_case_fields/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	userc "github.com/akm/goa_v2_post_payload_including_snake_case_fields/gen/http/user/client"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `user (create|update)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` user create --body '{
      "first_name": "Nihil magni vel dolores quia quo.",
      "last_name": "Sit officiis maiores."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		userFlags = flag.NewFlagSet("user", flag.ContinueOnError)

		userCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		userCreateBodyFlag = userCreateFlags.String("body", "REQUIRED", "")

		userUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		userUpdateBodyFlag = userUpdateFlags.String("body", "REQUIRED", "")
		userUpdateIDFlag   = userUpdateFlags.String("id", "REQUIRED", "")
	)
	userFlags.Usage = userUsage
	userCreateFlags.Usage = userCreateUsage
	userUpdateFlags.Usage = userUpdateUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if len(os.Args) < flag.NFlag()+3 {
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = os.Args[1+flag.NFlag()]
		switch svcn {
		case "user":
			svcf = userFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(os.Args[2+flag.NFlag():]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = os.Args[2+flag.NFlag()+svcf.NFlag()]
		switch svcn {
		case "user":
			switch epn {
			case "create":
				epf = userCreateFlags

			case "update":
				epf = userUpdateFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if len(os.Args) > 2+flag.NFlag()+svcf.NFlag() {
		if err := epf.Parse(os.Args[3+flag.NFlag()+svcf.NFlag():]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "user":
			c := userc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = userc.BuildCreatePayload(*userCreateBodyFlag)
			case "update":
				endpoint = c.Update()
				data, err = userc.BuildUpdatePayload(*userUpdateBodyFlag, *userUpdateIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// userUsage displays the usage of the user command and its subcommands.
func userUsage() {
	fmt.Fprintf(os.Stderr, `Service is the user service interface.
Usage:
    %s [globalflags] user COMMAND [flags]

COMMAND:
    create: Create implements create.
    update: Update implements update.

Additional help:
    %s user COMMAND --help
`, os.Args[0], os.Args[0])
}
func userCreateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user create -body JSON

Create implements create.
    -body JSON: 

Example:
    `+os.Args[0]+` user create --body '{
      "first_name": "Nihil magni vel dolores quia quo.",
      "last_name": "Sit officiis maiores."
   }'
`, os.Args[0])
}

func userUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user update -body JSON -id INT

Update implements update.
    -body JSON: 
    -id INT: 

Example:
    `+os.Args[0]+` user update --body '{
      "first_name": "Modi quod.",
      "last_name": "Vero omnis commodi inventore molestiae voluptas."
   }' --id 2294843666486776210
`, os.Args[0])
}

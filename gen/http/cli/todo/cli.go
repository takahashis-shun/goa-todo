// Code generated by goa v3.7.0, DO NOT EDIT.
//
// todo HTTP client CLI support package
//
// Command:
// $ goa gen github.com/takahashis-shun/todo-goa/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	todoc "github.com/takahashis-shun/todo-goa/gen/http/todo/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `todo (hello|show|create)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` todo hello --name "Blanditiis iure voluptas."` + "\n" +
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
		todoFlags = flag.NewFlagSet("todo", flag.ContinueOnError)

		todoHelloFlags    = flag.NewFlagSet("hello", flag.ExitOnError)
		todoHelloNameFlag = todoHelloFlags.String("name", "REQUIRED", "Name")

		todoShowFlags  = flag.NewFlagSet("show", flag.ExitOnError)
		todoShowIDFlag = todoShowFlags.String("id", "REQUIRED", "ID")

		todoCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		todoCreateBodyFlag = todoCreateFlags.String("body", "REQUIRED", "")
	)
	todoFlags.Usage = todoUsage
	todoHelloFlags.Usage = todoHelloUsage
	todoShowFlags.Usage = todoShowUsage
	todoCreateFlags.Usage = todoCreateUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "todo":
			svcf = todoFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "todo":
			switch epn {
			case "hello":
				epf = todoHelloFlags

			case "show":
				epf = todoShowFlags

			case "create":
				epf = todoCreateFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
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
		case "todo":
			c := todoc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "hello":
				endpoint = c.Hello()
				data, err = todoc.BuildHelloPayload(*todoHelloNameFlag)
			case "show":
				endpoint = c.Show()
				data, err = todoc.BuildShowPayload(*todoShowIDFlag)
			case "create":
				endpoint = c.Create()
				data, err = todoc.BuildCreatePayload(*todoCreateBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// todoUsage displays the usage of the todo command and its subcommands.
func todoUsage() {
	fmt.Fprintf(os.Stderr, `Service that manage todo.
Usage:
    %[1]s [globalflags] todo COMMAND [flags]

COMMAND:
    hello: Hello implements hello.
    show: Show implements show.
    create: Create implements create.

Additional help:
    %[1]s todo COMMAND --help
`, os.Args[0])
}
func todoHelloUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todo hello -name STRING

Hello implements hello.
    -name STRING: Name

Example:
    %[1]s todo hello --name "Blanditiis iure voluptas."
`, os.Args[0])
}

func todoShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todo show -id INT

Show implements show.
    -id INT: ID

Example:
    %[1]s todo show --id 151656195322834973
`, os.Args[0])
}

func todoCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todo create -body JSON

Create implements create.
    -body JSON: 

Example:
    %[1]s todo create --body '{
      "title": "Totam voluptatibus adipisci eos vel."
   }'
`, os.Args[0])
}

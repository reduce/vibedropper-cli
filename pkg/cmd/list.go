// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/reduce/vibedropper-cli/internal/apiquery"
	"github.com/reduce/vibedropper-cli/internal/requestflag"
	"github.com/reduce/vibedropper-go"
	"github.com/reduce/vibedropper-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var listsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a list",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "list-id",
			Required: true,
		},
	},
	Action:          handleListsRetrieve,
	HideHelpCommand: true,
}

var listsList = cli.Command{
	Name:    "list",
	Usage:   "List lists",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Default:   1,
			QueryPath: "page",
		},
	},
	Action:          handleListsList,
	HideHelpCommand: true,
}

func handleListsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("list-id") && len(unusedArgs) > 0 {
		cmd.Set("list-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Lists.Get(ctx, cmd.Value("list-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "lists retrieve", obj, format, transform)
}

func handleListsList(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := vibedropper.ListListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Lists.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "lists list", obj, format, transform)
}

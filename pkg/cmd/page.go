// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/vibedropper-cli/internal/apiquery"
	"github.com/stainless-sdks/vibedropper-cli/internal/requestflag"
	"github.com/stainless-sdks/vibedropper-go"
	"github.com/stainless-sdks/vibedropper-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var pagesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a page",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "page-id",
			Required: true,
		},
	},
	Action:          handlePagesRetrieve,
	HideHelpCommand: true,
}

var pagesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a page",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "page-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			BodyPath: "status",
		},
	},
	Action:          handlePagesUpdate,
	HideHelpCommand: true,
}

var pagesList = cli.Command{
	Name:    "list",
	Usage:   "List pages",
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
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     `Filter by status. Omit or use "all" to return all pages.`,
			QueryPath: "status",
		},
	},
	Action:          handlePagesList,
	HideHelpCommand: true,
}

var pagesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a page",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "page-id",
			Required: true,
		},
	},
	Action:          handlePagesDelete,
	HideHelpCommand: true,
}

func handlePagesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("page-id") && len(unusedArgs) > 0 {
		cmd.Set("page-id", unusedArgs[0])
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
	_, err = client.Pages.Get(ctx, cmd.Value("page-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "pages retrieve", obj, format, transform)
}

func handlePagesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("page-id") && len(unusedArgs) > 0 {
		cmd.Set("page-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := vibedropper.PageUpdateParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pages.Update(
		ctx,
		cmd.Value("page-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "pages update", obj, format, transform)
}

func handlePagesList(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := vibedropper.PageListParams{}

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
	_, err = client.Pages.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "pages list", obj, format, transform)
}

func handlePagesDelete(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("page-id") && len(unusedArgs) > 0 {
		cmd.Set("page-id", unusedArgs[0])
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
	_, err = client.Pages.Delete(ctx, cmd.Value("page-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "pages delete", obj, format, transform)
}

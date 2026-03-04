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

var formsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a form",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "form-id",
			Required: true,
		},
	},
	Action:          handleFormsRetrieve,
	HideHelpCommand: true,
}

var formsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a form",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "form-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "list-id",
			Usage:    "List to subscribe form submitters to",
			BodyPath: "listId",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			BodyPath: "status",
		},
		&requestflag.Flag[any]{
			Name:     "success-message",
			BodyPath: "successMessage",
		},
		&requestflag.Flag[string]{
			Name:     "title",
			BodyPath: "title",
		},
	},
	Action:          handleFormsUpdate,
	HideHelpCommand: true,
}

var formsList = cli.Command{
	Name:    "list",
	Usage:   "List forms",
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
			Usage:     `Filter by status. Omit or use "all" to return all forms.`,
			QueryPath: "status",
		},
	},
	Action:          handleFormsList,
	HideHelpCommand: true,
}

var formsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a form",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "form-id",
			Required: true,
		},
	},
	Action:          handleFormsDelete,
	HideHelpCommand: true,
}

var formsListSubmissions = cli.Command{
	Name:    "list-submissions",
	Usage:   "List form submissions",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "form-id",
			Required: true,
		},
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
	Action:          handleFormsListSubmissions,
	HideHelpCommand: true,
}

func handleFormsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("form-id") && len(unusedArgs) > 0 {
		cmd.Set("form-id", unusedArgs[0])
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
	_, err = client.Forms.Get(ctx, cmd.Value("form-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "forms retrieve", obj, format, transform)
}

func handleFormsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("form-id") && len(unusedArgs) > 0 {
		cmd.Set("form-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := vibedropper.FormUpdateParams{}

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
	_, err = client.Forms.Update(
		ctx,
		cmd.Value("form-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "forms update", obj, format, transform)
}

func handleFormsList(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := vibedropper.FormListParams{}

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
	_, err = client.Forms.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "forms list", obj, format, transform)
}

func handleFormsDelete(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("form-id") && len(unusedArgs) > 0 {
		cmd.Set("form-id", unusedArgs[0])
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
	_, err = client.Forms.Delete(ctx, cmd.Value("form-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "forms delete", obj, format, transform)
}

func handleFormsListSubmissions(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("form-id") && len(unusedArgs) > 0 {
		cmd.Set("form-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := vibedropper.FormListSubmissionsParams{}

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
	_, err = client.Forms.ListSubmissions(
		ctx,
		cmd.Value("form-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "forms list-submissions", obj, format, transform)
}

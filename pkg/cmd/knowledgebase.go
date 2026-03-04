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

var knowledgeBasesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a knowledge base",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "kb-id",
			Required: true,
		},
	},
	Action:          handleKnowledgeBasesRetrieve,
	HideHelpCommand: true,
}

var knowledgeBasesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a knowledge base",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "kb-id",
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
		&requestflag.Flag[int64]{
			Name:     "sort-order",
			BodyPath: "sortOrder",
		},
	},
	Action:          handleKnowledgeBasesUpdate,
	HideHelpCommand: true,
}

var knowledgeBasesList = cli.Command{
	Name:            "list",
	Usage:           "List knowledge bases",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleKnowledgeBasesList,
	HideHelpCommand: true,
}

var knowledgeBasesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a knowledge base",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "kb-id",
			Required: true,
		},
	},
	Action:          handleKnowledgeBasesDelete,
	HideHelpCommand: true,
}

func handleKnowledgeBasesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("kb-id") && len(unusedArgs) > 0 {
		cmd.Set("kb-id", unusedArgs[0])
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
	_, err = client.KnowledgeBases.Get(ctx, cmd.Value("kb-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "knowledge-bases retrieve", obj, format, transform)
}

func handleKnowledgeBasesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("kb-id") && len(unusedArgs) > 0 {
		cmd.Set("kb-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := vibedropper.KnowledgeBaseUpdateParams{}

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
	_, err = client.KnowledgeBases.Update(
		ctx,
		cmd.Value("kb-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "knowledge-bases update", obj, format, transform)
}

func handleKnowledgeBasesList(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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
	_, err = client.KnowledgeBases.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "knowledge-bases list", obj, format, transform)
}

func handleKnowledgeBasesDelete(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("kb-id") && len(unusedArgs) > 0 {
		cmd.Set("kb-id", unusedArgs[0])
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

	return client.KnowledgeBases.Delete(ctx, cmd.Value("kb-id").(string), options...)
}

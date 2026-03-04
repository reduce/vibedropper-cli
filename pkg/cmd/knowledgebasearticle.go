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

var knowledgeBasesArticlesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create an article",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "kb-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "title",
			Required: true,
			BodyPath: "title",
		},
		&requestflag.Flag[any]{
			Name:     "category-id",
			BodyPath: "categoryId",
		},
		&requestflag.Flag[string]{
			Name:     "content",
			Usage:    "HTML or markdown content",
			BodyPath: "content",
		},
		&requestflag.Flag[any]{
			Name:     "excerpt",
			BodyPath: "excerpt",
		},
		&requestflag.Flag[bool]{
			Name:     "published",
			Default:  true,
			BodyPath: "published",
		},
	},
	Action:          handleKnowledgeBasesArticlesCreate,
	HideHelpCommand: true,
}

var knowledgeBasesArticlesList = cli.Command{
	Name:    "list",
	Usage:   "List articles in a knowledge base",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "kb-id",
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
	Action:          handleKnowledgeBasesArticlesList,
	HideHelpCommand: true,
}

func handleKnowledgeBasesArticlesCreate(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("kb-id") && len(unusedArgs) > 0 {
		cmd.Set("kb-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := vibedropper.KnowledgeBaseArticleNewParams{}

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
	_, err = client.KnowledgeBases.Articles.New(
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
	return ShowJSON(os.Stdout, "knowledge-bases:articles create", obj, format, transform)
}

func handleKnowledgeBasesArticlesList(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("kb-id") && len(unusedArgs) > 0 {
		cmd.Set("kb-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := vibedropper.KnowledgeBaseArticleListParams{}

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
	_, err = client.KnowledgeBases.Articles.List(
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
	return ShowJSON(os.Stdout, "knowledge-bases:articles list", obj, format, transform)
}

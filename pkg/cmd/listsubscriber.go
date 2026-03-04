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

var listsSubscribersList = cli.Command{
	Name:    "list",
	Usage:   "Returns all subscribers for the list ordered by subscribe date descending.\nIncludes linked customer data.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "list-id",
			Required: true,
		},
	},
	Action:          handleListsSubscribersList,
	HideHelpCommand: true,
}

var listsSubscribersAdd = cli.Command{
	Name:    "add",
	Usage:   "Creates or updates the matching customer record and adds a subscriber entry.\nReturns 400 with code `duplicate` if already subscribed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "list-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "email",
			Required: true,
			BodyPath: "email",
		},
		&requestflag.Flag[any]{
			Name:     "custom-fields",
			Usage:    "Arbitrary key-value metadata",
			BodyPath: "customFields",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "pickup-location-id",
			Usage:    "Pickup location ID (must belong to the given regionId)",
			BodyPath: "pickupLocationId",
		},
		&requestflag.Flag[string]{
			Name:     "region-id",
			Usage:    "Region ID to assign to the customer",
			BodyPath: "regionId",
		},
	},
	Action:          handleListsSubscribersAdd,
	HideHelpCommand: true,
}

var listsSubscribersRemove = cli.Command{
	Name:    "remove",
	Usage:   "Remove a subscriber from a list",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "list-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "subscriber-id",
			Required: true,
		},
	},
	Action:          handleListsSubscribersRemove,
	HideHelpCommand: true,
}

func handleListsSubscribersList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Lists.Subscribers.List(ctx, cmd.Value("list-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "lists:subscribers list", obj, format, transform)
}

func handleListsSubscribersAdd(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("list-id") && len(unusedArgs) > 0 {
		cmd.Set("list-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := vibedropper.ListSubscriberAddParams{}

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
	_, err = client.Lists.Subscribers.Add(
		ctx,
		cmd.Value("list-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "lists:subscribers add", obj, format, transform)
}

func handleListsSubscribersRemove(ctx context.Context, cmd *cli.Command) error {
	client := vibedropper.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscriber-id") && len(unusedArgs) > 0 {
		cmd.Set("subscriber-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := vibedropper.ListSubscriberRemoveParams{
		ListID: cmd.Value("list-id").(string),
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
	_, err = client.Lists.Subscribers.Remove(
		ctx,
		cmd.Value("subscriber-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "lists:subscribers remove", obj, format, transform)
}

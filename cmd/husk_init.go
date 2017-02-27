// Copyright 2016 Pulumi, Inc. All rights reserved.

package cmd

import (
	"github.com/spf13/cobra"
)

func newHuskInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init <husk>",
		Short: "Create an empty husk with the given name, ready for deployments",
		Long: "Create an empty husk with the given name, ready for deployments\n" +
			"\n" +
			"This command creates an empty husk with the given name.  It has no resources, but\n" +
			"afterwards it can become the target of a deployment using the `deploy` command.",
		Run: func(cmd *cobra.Command, args []string) {
			create(cmd, args)
		},
	}
}
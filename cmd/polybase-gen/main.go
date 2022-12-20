/*
 * Copyright © 2022 V1def
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"os"

	"github.com/v1def/go-polybase/codegen"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	configPath string

	root = &cobra.Command{
		Use:   "polybase-gen",
		Short: "Code generation for working with Polybase collection.",
		Long:  "polybase-gen generates the code needed to work with the Polybase collection.",
		Run: func(cmd *cobra.Command, args []string) {
			if configPath == "" {
				log.Fatal().Msg("Config file not specified!")
			}

			generator, err := codegen.New(configPath)
			if err != nil {
				log.Fatal().Err(err).Msg("error initialize code-generator")
			}

			if err := generator.Generate(); err != nil {
				log.Fatal().Err(err).Msg("error generating collection")
			}
		},
	}
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	root.Flags().StringVarP(&configPath, "config", "c", "", "config path")
}

func main() {
	if err := root.Execute(); err != nil {
		log.Fatal().Err(err).Msg("error executing cli")
	}
}

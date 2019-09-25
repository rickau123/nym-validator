// run.go - nym-tendermint node startup definition.
// Copyright (C) 2019  Nym Authors.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"

	"github.com/nymtech/nym-validator/daemon"
	"github.com/nymtech/nym-validator/tendermint/nymnode"
)

const (
	serviceName                = "nym-tendermint-node"
	defaultConfigFile          = "/tendermint/config/config.toml"
	defaultDataRoot            = "/tendermint"
	defaultEmptyBlocksInterval = 0
)

func cmdRun(args []string, usage string) {
	opts := daemon.NewOpts(serviceName, "run [OPTIONS]", usage)

	daemon.Start(func(args []string) daemon.Service {
		cfgFile := opts.Flags("--cfgFile").Label("CFGFILE").String(
			"Path to the main tendermint configuration file",
			defaultConfigFile,
		)
		dataRoot := opts.Flags("--dataRoot").Label("DATAROOT").String(
			"Path to the data root directory",
			defaultDataRoot,
		)
		createEmptyBlocks := opts.Flags("--createEmptyBlocks").Label("EMPTYBLOCKSFLAG").Bool(
			"Flag to indicate whether tendermint should create empty blocks",
		)
		emptyBlocksInterval := opts.Flags("--emptyBlocksInteral").Label("EMPTYBLOCKSINTERVAL").Duration(
			"(if applicable) used to indicate interval between empty blocks",
			defaultEmptyBlocksInterval,
		)

		params := opts.Parse(args)
		if len(params) != 0 {
			opts.PrintUsage()
			os.Exit(-1)
		}

		node, err := nymnode.CreateNymNode(*cfgFile, *dataRoot, *createEmptyBlocks, *emptyBlocksInterval)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create NymNode: %v\n", err)
			os.Exit(-1)
		}
		return node
	}, args)
}

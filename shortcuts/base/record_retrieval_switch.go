// Copyright (c) 2026 Lark Technologies Pte. Ltd.
// SPDX-License-Identifier: MIT

package base

import (
	"context"

	"github.com/larksuite/cli/shortcuts/common"
)

var BaseRecordRetrievalSwitch = common.Shortcut{
	Service:     "base",
	Command:     "+record-retrieval-switch",
	Description: "Enable or disable record retrieval for a base",
	Risk:        "write",
	Scopes:      []string{"base:record:retrieve"},
	AuthTypes:   authTypes(),
	Flags: []common.Flag{
		baseTokenFlag(true),
		{Name: "enable", Desc: "enable or disable retrieval", Required: true, Enum: []string{"true", "false"}},
	},
	DryRun: dryRunRecordRetrievalSwitch,
	Execute: func(ctx context.Context, runtime *common.RuntimeContext) error {
		return executeRecordRetrievalSwitch(runtime)
	},
}

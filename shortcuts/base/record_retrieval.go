// Copyright (c) 2026 Lark Technologies Pte. Ltd.
// SPDX-License-Identifier: MIT

package base

import (
	"context"

	"github.com/larksuite/cli/shortcuts/common"
)

var BaseRecordRetrieval = common.Shortcut{
	Service:     "base",
	Command:     "+record-retrieval",
	Description: "Retrieve records by semantic query",
	Risk:        "read",
	Scopes:      []string{"base:record:retrieve"},
	AuthTypes:   authTypes(),
	Flags: []common.Flag{
		baseTokenFlag(true),
		{Name: "query", Desc: "search query string", Required: true},
	},
	DryRun: dryRunRecordRetrieval,
	Execute: func(ctx context.Context, runtime *common.RuntimeContext) error {
		return executeRecordRetrieval(runtime)
	},
}

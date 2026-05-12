// Copyright (c) 2026 Lark Technologies Pte. Ltd.
// SPDX-License-Identifier: MIT

package base

import (
	"context"

	"github.com/larksuite/cli/shortcuts/common"
)

var BaseRecordRetrievalPoll = common.Shortcut{
	Service:     "base",
	Command:     "+record-retrieval-poll",
	Description: "Poll record retrieval result by retrieval ID",
	Risk:        "read",
	Scopes:      []string{"base:record:retrieve"},
	AuthTypes:   authTypes(),
	Flags: []common.Flag{
		baseTokenFlag(true),
		{Name: "retrieval-id", Desc: "retrieval ID returned by +record-retrieval", Required: true},
	},
	DryRun: dryRunRecordRetrievalPoll,
	Execute: func(ctx context.Context, runtime *common.RuntimeContext) error {
		return executeRecordRetrievalPoll(runtime)
	},
}

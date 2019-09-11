/*
 * MIT License
 *
 * Copyright (c) 2019 Dariusz Depta Engos Software
 *
 * License details in LICENSE.
 */

package suites

import (
    o "github.com/EngosSoftware/oxyde"
    tc "software/engos/tarolas/tests/cases"
    c "software/engos/tarolas/tests/common"
)

func All(ctx *c.Context, dtx *o.DocContext) {
    c.Display(ctx)
    c.TsNames(ctx, dtx)
    tc.TsDirectoryCreate(ctx, dtx)
    tc.TsDirectoryDelete(ctx, dtx)
    tc.TsDirectoryRead(ctx, dtx)
    tc.TsFileAppend(ctx, dtx)
    tc.TsFileExists(ctx, dtx)
    tc.TsFileChecksum(ctx, dtx)
}

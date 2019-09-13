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

func Single(ctx *c.Context, dtx *o.DocContext) {
    c.Display(ctx)
    ctx.Verbose = false
    tc.TcFileAppend(ctx, dtx)

    // start documentation preview server
    o.StartPreview(dtx)
}

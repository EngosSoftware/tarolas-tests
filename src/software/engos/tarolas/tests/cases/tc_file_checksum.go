/*
 * MIT License
 *
 * Copyright (c) 2019 Dariusz Depta Engos Software
 *
 * License details in LICENSE.
 */

package cases

import (
    "github.com/EngosSoftware/oxyde"
    c "software/engos/tarolas/tests/common"
)

const (
    FileChecksumUrl = "/file/checksum"
)

type FileChecksumParams struct {
    Name *string `json:"name"  api:"Name of the file for which checksum will be calculated."`
}

type FileChecksumResult struct {
    Data   c.FileDto    `json:"data"`
    Errors []c.ErrorDto `json:"errors"`
}

func TsFileChecksum(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    TdFileChecksum(ctx, dtx)
}

func TdFileChecksum(ctx *c.Context, dtx *oxyde.DocContext) {
    const summary = `Calculates the file's checksum.`
    const description = `(to be updated)`
    c.Display(ctx)
    dtx.NewEndpoint(ctx.Version, c.FilesTag, summary, description)
    TcFileChecksumSmall(ctx, dtx)
    TcFileChecksumBig(ctx, dtx)
    TcFileChecksumNotFound(ctx, dtx)
    dtx.SaveEndpoint()
}

func TcFileChecksumSmall(ctx *c.Context, dtx *oxyde.DocContext) {
    const summary = `
Calculates checksum.
`
    const description = `
`
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    fileName := c.RootDirName + c.FileNames[c.FileA]
    FileAppend(ctx, dtx, fileName, []byte(c.FileContents[c.FileA]))
    params := FileChecksumParams{Name: &fileName}
    var result FileChecksumResult
    dtx.CollectAll(summary, description)
    oxyde.HttpGET(ctx, dtx, FileChecksumUrl, nil, &params, &result, 200)
    c.DisplayOK(ctx)
}

func TcFileChecksumBig(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    fileName := c.RootDirName + c.FileNames[c.FileA]
    FileAppendBig(ctx, dtx, fileName, 1024, 10)
    params := FileChecksumParams{Name: &fileName}
    var result FileChecksumResult
    oxyde.HttpGET(ctx, dtx, FileChecksumUrl, nil, &params, &result, 200)
    c.DisplayOK(ctx)
}

func TcFileChecksumNotFound(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    fileName := c.RootDirName + c.FileNames[c.FileA]
    params := FileChecksumParams{Name: &fileName}
    var result FileChecksumResult
    oxyde.HttpGET(ctx, dtx, FileChecksumUrl, nil, &params, &result, 400)
    c.AssertError(result.Errors, 400, "file not found", fileName, 10938)
    c.DisplayOK(ctx)
}

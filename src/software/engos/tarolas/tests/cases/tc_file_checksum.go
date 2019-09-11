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
    c.Display(ctx)
    TcFileChecksumSmall(ctx, dtx)
    TcFileChecksumBig(ctx, dtx)
    TcFileChecksumNotFound(ctx, dtx)
}

func TcFileChecksumSmall(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    FileAppend(ctx, dtx, c.FileNames[c.FileA], []byte(c.FileContents[c.FileA]))
    params := FileChecksumParams{Name: &c.FileNames[c.FileA]}
    var result FileChecksumResult
    oxyde.HttpGET(ctx, dtx, FileChecksumUrl, nil, &params, &result, 200)
    c.DisplayOK(ctx)
}

func TcFileChecksumBig(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    FileAppendBig(ctx, dtx, c.FileNames[c.FileA], 1024, 10)
    params := FileChecksumParams{Name: &c.FileNames[c.FileA]}
    var result FileChecksumResult
    oxyde.HttpGET(ctx, dtx, FileChecksumUrl, nil, &params, &result, 200)
    c.DisplayOK(ctx)
}

func TcFileChecksumNotFound(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    params := FileChecksumParams{Name: &c.FileNames[c.FileA]}
    var result FileChecksumResult
    oxyde.HttpGET(ctx, dtx, FileChecksumUrl, nil, &params, &result, 400)
    c.AssertError(result.Errors, 400, "file not found", c.FileNames[c.FileA], 10938)
    c.DisplayOK(ctx)
}

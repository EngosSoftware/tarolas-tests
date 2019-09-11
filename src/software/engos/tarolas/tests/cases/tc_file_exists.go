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
    FileExistsUrl = "/file/exists"
)

type FileExistsParams struct {
    Name *string `json:"name"  api:"Name of the file to be checked for existence."`
}

type FileExistsResult struct {
    Data   c.FileDto    `json:"data"`
    Errors []c.ErrorDto `json:"errors"`
}

func TsFileExists(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    TdFileExists(ctx, dtx)
}

func TdFileExists(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    TcFileExists(ctx, dtx)
    TcFileExistsBig(ctx, dtx)
    TcFileExistsNotFound(ctx, dtx)
}

func TcFileExists(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    FileAppend(ctx, dtx, c.FileNames[c.FileA], []byte(c.FileContents[c.FileA]))
    params := FileExistsParams{Name: &c.FileNames[c.FileA]}
    var result FileExistsResult
    oxyde.HttpGET(ctx, dtx, FileExistsUrl, nil, &params, &result, 200)
    c.DisplayOK(ctx)
}

func TcFileExistsBig(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    FileAppendBig(ctx, dtx, c.FileNames[c.FileA], 1024, 10)
    params := FileExistsParams{Name: &c.FileNames[c.FileA]}
    var result FileExistsResult
    oxyde.HttpGET(ctx, dtx, FileExistsUrl, nil, &params, &result, 200)
    c.DisplayOK(ctx)
}

func TcFileExistsNotFound(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    params := FileExistsParams{Name: &c.FileNames[c.FileA]}
    var result FileExistsResult
    oxyde.HttpGET(ctx, dtx, FileExistsUrl, nil, &params, &result, 400)
    c.AssertError(result.Errors, 400, "file not found", c.FileNames[c.FileA], 10938)
    c.DisplayOK(ctx)
}

/*
 * MIT License
 *
 * Copyright (c) 2019 Dariusz Depta Engos Software
 *
 * License details in LICENSE.
 */

package cases

import (
    "fmt"
    o "github.com/EngosSoftware/oxyde"
    c "software/engos/tarolas/tests/common"
)

const (
    FileExistsUrl = "/file/exists"
)

type FileExistsParams struct {
    Name *string `json:"name"  api:"Name of the file to be checked if exists."`
}

type FileExistsDto struct {
    Name   *string `json:"name"      api:"File name without parent path."`
    Size   *int64  `json:"size"      api:"?Optional file size in bytes, set only when file exists."`
    Exists *bool   `json:"exists"    api:"Flag indicating if file exists."`
}

type FileExistsResult struct {
    Data   FileExistsDto `json:"data"    api:"?File searching result when processing this request was successful."`
    Errors []c.ErrorDto  `json:"errors"  api:"?List of errors when something went wrong during processing this request."`
}

func TsFileExists(ctx *c.Context, dtx *o.DocContext) {
    c.Display(ctx)
    TdFileExists(ctx, dtx)
}

func TdFileExists(ctx *c.Context, dtx *o.DocContext) {
    const summary = `Checks if file exists.`
    const description = `(to be updated)`
    c.Display(ctx)
    dtx.NewEndpoint(ctx.Version, c.FilesTag, summary, description)
    TcFileExists(ctx, dtx)
    TcFileExistsInDirectory(ctx, dtx)
    TcFileExistsBig(ctx, dtx)
    TcFileExistsNotFound(ctx, dtx)
    TcFileExistsNotFoundInDirectory(ctx, dtx)
    // TODO add test of invalid file names
    // TODO add test to check for file when directory name is given
    dtx.SaveEndpoint()
}

func TcFileExists(ctx *c.Context, dtx *o.DocContext) {
    const summary = `Checks if file exists.`
    var description = fmt.Sprintf(`Assuming that file **%s** exists in a root directory.`, c.FileNames[c.FileA])

    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    fileName := c.RootDirName + c.FileNames[c.FileA]
    FileAppend(ctx, dtx, fileName, []byte(c.FileContents[c.FileA]))
    params := FileExistsParams{Name: &fileName}
    var result FileExistsResult
    dtx.CollectAll(summary, description)
    o.HttpGET(ctx, dtx, FileExistsUrl, nil, &params, &result, 200)
    o.AssertEqualStringNullable(&fileName, result.Data.Name)
    o.AssertEqualInt64Nullable(&c.FileSizes[c.FileA], result.Data.Size)
    o.AssertTrue(*result.Data.Exists)
    c.DisplayOK(ctx)
}

func TcFileExistsInDirectory(ctx *c.Context, dtx *o.DocContext) {
    const summary = `Checks if file exists in directory.`
    var description = fmt.Sprintf(`Assuming that file **%s** exists in directory **%s**.`, c.FileNames[c.FileA], c.DirectoryNames[c.DirectoryA])
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    directoryName := c.RootDirName + c.DirectoryNames[c.DirectoryA]
    DirectoryCreate(ctx, dtx, directoryName, false)
    fileName := directoryName + "/" + c.FileNames[c.FileA]
    FileAppend(ctx, dtx, fileName, []byte(c.FileContents[c.FileA]))
    params := FileExistsParams{Name: &fileName}
    var result FileExistsResult
    dtx.CollectUsage(summary, description)
    o.HttpGET(ctx, dtx, FileExistsUrl, nil, &params, &result, 200)
    o.AssertEqualStringNullable(&fileName, result.Data.Name)
    o.AssertEqualInt64Nullable(&c.FileSizes[c.FileA], result.Data.Size)
    o.AssertTrue(*result.Data.Exists)
    c.DisplayOK(ctx)
}

func TcFileExistsBig(ctx *c.Context, dtx *o.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    fileName := c.RootDirName + c.FileNames[c.FileA]
    FileAppendBig(ctx, dtx, fileName, 1024, 10)
    params := FileExistsParams{Name: &fileName}
    var result FileExistsResult
    o.HttpGET(ctx, dtx, FileExistsUrl, nil, &params, &result, 200)
    c.DisplayOK(ctx)
}

func TcFileExistsNotFound(ctx *c.Context, dtx *o.DocContext) {
    const summary = `Checks if the file does not exist.`
    var description = fmt.Sprintf("Assuming that there is no **%s** file in a root directory.", c.FileNames[c.FileB])

    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    fileName := c.RootDirName + c.FileNames[c.FileB]
    params := FileExistsParams{Name: &fileName}
    var result FileExistsResult
    dtx.CollectUsage(summary, description)
    o.HttpGET(ctx, dtx, FileExistsUrl, nil, &params, &result, 200)
    o.AssertEqualStringNullable(&fileName, result.Data.Name)
    o.AssertNil(result.Data.Size)
    o.AssertFalse(*result.Data.Exists)
    c.DisplayOK(ctx)
}

func TcFileExistsNotFoundInDirectory(ctx *c.Context, dtx *o.DocContext) {
    const summary = `Checks if file does not exist in directory.`
    var description = fmt.Sprintf(`Assuming that there is no **%s** file in directory **%s**.`, c.FileNames[c.FileA], c.DirectoryNames[c.DirectoryA])
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    directoryName := c.RootDirName + c.DirectoryNames[c.DirectoryA]
    DirectoryCreate(ctx, dtx, directoryName, false)
    fileName := directoryName + "/" + c.FileNames[c.FileB]
    params := FileExistsParams{Name: &fileName}
    var result FileExistsResult
    dtx.CollectUsage(summary, description)
    o.HttpGET(ctx, dtx, FileExistsUrl, nil, &params, &result, 200)
    o.AssertEqualStringNullable(&fileName, result.Data.Name)
    o.AssertNil(result.Data.Size)
    o.AssertFalse(*result.Data.Exists)
    c.DisplayOK(ctx)
}

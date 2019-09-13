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
    FileDeleteUrl = "/file/delete"
)

type FileDeleteParams struct {
    Name *string `json:"name"  api:"Name of the file to be deleted."`
}

type FileDeleteDto struct {
    Name *string `json:"name"      api:"Deleted file name without parent path."`
    Size *int64  `json:"size"      api:"Deleted file size before deletion."`
}

type FileDeleteResult struct {
    Data   FileDeleteDto `json:"data"    api:"?Deleted file details when processing this request was successful."`
    Errors []c.ErrorDto  `json:"errors"  api:"?List of errors when something went wrong during processing this request."`
}

func TsFileDelete(ctx *c.Context, dtx *o.DocContext) {
    c.Display(ctx)
    TdFileDelete(ctx, dtx)
}

func TdFileDelete(ctx *c.Context, dtx *o.DocContext) {
    const summary = `Deletes an existing file.`
    const description = `(to be updated)`
    c.Display(ctx)
    dtx.NewEndpoint(ctx.Version, c.FilesTag, summary, description)
    TcFileDelete(ctx, dtx)
    TcFileDeleteInDirectory(ctx, dtx)
    TcFileDeleteNotFound(ctx, dtx)
    TcFileDeleteNotFoundInDirectory(ctx, dtx)
    // TODO add test of invalid file names
    // TODO add test to check for file when directory name is given
    dtx.SaveEndpoint()
}

func TcFileDelete(ctx *c.Context, dtx *o.DocContext) {
    // SET-UP
    c.Display(ctx)
    // DOCUMENTATION
    const summary = `Deletes an existing file.`
    var description = fmt.Sprintf(`Assuming that file **%s** exists in a root directory.`, c.FileNames[c.FileA])
    // GIVEN
    RemoveRootContents(ctx, dtx)
    fileName := c.RootDirName + c.FileNames[c.FileA]
    FileAppend(ctx, dtx, fileName, []byte(c.FileContents[c.FileA]))
    params := FileDeleteParams{Name: &fileName}
    var result FileDeleteResult
    // WHEN
    dtx.CollectAll(summary, description)
    o.HttpDELETE(ctx, dtx, FileDeleteUrl, nil, &params, nil, &result, 200)
    // THEN
    o.AssertEqualStringNullable(&fileName, result.Data.Name)
    o.AssertEqualInt64Nullable(&c.FileSizes[c.FileA], result.Data.Size)
    // root directory should be empty
    dir := DirectoryRead(ctx, dtx, c.RootDirName)
    c.AssertEmptyDirectory(dir, c.RootDirName)
    // TEAR-DOWN
    c.DisplayOK(ctx)
}

func TcFileDeleteInDirectory(ctx *c.Context, dtx *o.DocContext) {
    // SET-UP
    c.Display(ctx)
    // DOCUMENTATION
    const summary = `Deletes an existing file in subdirectory.`
    var description = fmt.Sprintf(`Assuming that file **%s** exists in directory **%s**.`, c.FileNames[c.FileA], c.DirectoryNames[c.DirectoryA])
    // GIVEN
    RemoveRootContents(ctx, dtx)
    directoryName := c.RootDirName + c.DirectoryNames[c.DirectoryA]
    DirectoryCreate(ctx, dtx, directoryName, false)
    fileName := directoryName + "/" + c.FileNames[c.FileA]
    FileAppend(ctx, dtx, fileName, []byte(c.FileContents[c.FileA]))
    params := FileDeleteParams{Name: &fileName}
    var result FileDeleteResult
    // WHEN
    dtx.CollectUsage(summary, description)
    o.HttpDELETE(ctx, dtx, FileDeleteUrl, nil, &params, nil, &result, 200)
    // THEN
    o.AssertEqualStringNullable(&fileName, result.Data.Name)
    o.AssertEqualInt64Nullable(&c.FileSizes[c.FileA], result.Data.Size)
    // subdirectory should be empty
    dir := DirectoryRead(ctx, dtx, directoryName)
    c.AssertEmptyDirectory(dir, c.DirectoryNames[c.DirectoryA])
    // TEAR-DOWN
    c.DisplayOK(ctx)
}

func TcFileDeleteNotFound(ctx *c.Context, dtx *o.DocContext) {
    // SET-UP
    c.Display(ctx)
    // DOCUMENTATION
    const summary = `Attempts to delete a file that does not exist in root directory.`
    var description = fmt.Sprintf("Assuming that there is no **%s** file in a root directory.", c.FileNames[c.FileB])
    // GIVEN
    RemoveRootContents(ctx, dtx)
    fileName := c.RootDirName + c.FileNames[c.FileB]
    params := FileDeleteParams{Name: &fileName}
    var result FileDeleteResult
    // WHEN
    dtx.CollectUsage(summary, description)
    o.HttpDELETE(ctx, dtx, FileDeleteUrl, nil, &params, nil, &result, 400)
    // THEN
    c.AssertError(result.Errors, 400, "file not found", fileName, 10938)
    // TEAR-DOWN
    c.DisplayOK(ctx)
}

func TcFileDeleteNotFoundInDirectory(ctx *c.Context, dtx *o.DocContext) {
    // SET-UP
    c.Display(ctx)
    // DOCUMENTATION
    const summary = `Attempts to delete a file that does not exist in subdirectory.`
    var description = fmt.Sprintf(`Assuming that there is no **%s** file in directory **%s**.`, c.FileNames[c.FileA], c.DirectoryNames[c.DirectoryA])
    // GIVEN
    RemoveRootContents(ctx, dtx)
    directoryName := c.RootDirName + c.DirectoryNames[c.DirectoryA]
    DirectoryCreate(ctx, dtx, directoryName, false)
    fileName := directoryName + "/" + c.FileNames[c.FileB]
    params := FileDeleteParams{Name: &fileName}
    var result FileDeleteResult
    // WHEN
    dtx.CollectUsage(summary, description)
    o.HttpDELETE(ctx, dtx, FileDeleteUrl, nil, &params, nil, &result, 400)
    // THEN
    c.AssertError(result.Errors, 400, "file not found", fileName, 10938)
    // TEAR-DOWN
    c.DisplayOK(ctx)
}

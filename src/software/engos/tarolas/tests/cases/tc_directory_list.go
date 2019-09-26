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
    "sort"
)

const (
    DirectoryListUrl = "/directory/list"
)

type DirectoryListResult struct {
    Data   []string     `json:"data"`
    Errors []c.ErrorDto `json:"errors"`
}

type DirectoryListParams struct {
    Name *string `json:"name"  api:"Name of the directory to be listed."`
}

func TsDirectoryList(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    TdDirectoryList(ctx, dtx)
}

func TdDirectoryList(ctx *c.Context, dtx *oxyde.DocContext) {
    const summary = `Lists the directory tree.`
    const description = `(to be updated)`
    c.Display(ctx)
    dtx.NewEndpoint(ctx.Version, c.DirectoriesTag, summary, description)
    TcDirectoryListRootEmpty(ctx, dtx)
    TcDirectoryListRoot(ctx, dtx)
    TcDirectoryListInvalidName(ctx, dtx)
    TcDirectoryListNoPermissions(ctx, dtx)
    TcDirectoryListMultiple(ctx, dtx)
    TcDirectoryListMultipleDeep(ctx, dtx)
    dtx.SaveEndpoint()
}

func TcDirectoryListRootEmpty(ctx *c.Context, dtx *oxyde.DocContext) {
    const summary = `Listing the content of empty root directory.`
    const description = `If root directory is empty, then returned list of directories should be empty.`
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    params := DirectoryListParams{Name: &c.RootDirName}
    var result DirectoryListResult
    dtx.CollectUsage(summary, description)
    oxyde.HttpGET(ctx, dtx, DirectoryListUrl, nil, &params, &result, 200)
    oxyde.AssertNotNil(result.Data)
    oxyde.AssertEqualInt(0, len(result.Data))
    c.DisplayOK(ctx)
}

func TcDirectoryListRoot(ctx *c.Context, dtx *oxyde.DocContext) {
    const summary = `
Listing the content of root directory.
`
    const description = `
aaa
`
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    directoryName := c.RootDirName + c.DirectoryNames[c.DirectoryA]
    DirectoryCreate(ctx, dtx, directoryName, false)
    params := DirectoryListParams{Name: &c.RootDirName}
    var result DirectoryListResult
    dtx.CollectAll(summary, description)
    oxyde.HttpGET(ctx, dtx, DirectoryListUrl, nil, &params, &result, 200)
    oxyde.AssertNotNil(result.Data)
    oxyde.AssertEqualInt(1, len(result.Data))
    oxyde.AssertEqualString(directoryName, result.Data[0])
    c.DisplayOK(ctx)
}

func TcDirectoryListInvalidName(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    var result DirectoryListResult
    f := func(params DirectoryListParams, errorMessage string, errorDetails string, errorCode int) {
        oxyde.HttpGET(ctx, dtx, DirectoryListUrl, nil, &params, &result, 400)
        c.AssertError(result.Errors, 400, errorMessage, errorDetails, errorCode)
    }
    f(DirectoryListParams{Name: nil}, "required parameter is missing", "name", 10837)
    f(DirectoryListParams{Name: &c.EmptyValue}, "required parameter is empty", "name", 10767)
    f(DirectoryListParams{Name: &c.SpaceValue}, "required parameter is empty", "name", 10767)
    f(DirectoryListParams{Name: &c.WhiteSpaceValue}, "required parameter is empty", "name", 10767)
    f(DirectoryListParams{Name: &c.NoRootSlashName}, "file or directory name must begin with slash", "noRootSlashName", 10767)
    c.DisplayOK(ctx)
}

func TcDirectoryListMultiple(ctx *c.Context, dtx *oxyde.DocContext) {
    const summary = `Listing the content of root directory with multiple subdirectories.`
    const description = `If there is more than one subdirectory, all of them should be returned in the directory list.`
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    directoryNameA := c.RootDirName + c.DirectoryNames[c.DirectoryA]
    directoryNameB := c.RootDirName + c.DirectoryNames[c.DirectoryB]
    directoryNameC := c.RootDirName + c.DirectoryNames[c.DirectoryC]
    DirectoryCreate(ctx, dtx, directoryNameA, false)
    DirectoryCreate(ctx, dtx, directoryNameB, false)
    DirectoryCreate(ctx, dtx, directoryNameC, false)
    params := DirectoryListParams{Name: &c.RootDirName}
    var result DirectoryListResult
    dtx.CollectUsage(summary, description)
    oxyde.HttpGET(ctx, dtx, DirectoryListUrl, nil, &params, &result, 200)
    oxyde.AssertNotNil(result.Data)
    oxyde.AssertEqualInt(3, len(result.Data))
    names := []string{directoryNameA, directoryNameB, directoryNameC}
    sort.Strings(names)
    oxyde.AssertEqualString(names[0], result.Data[0])
    oxyde.AssertEqualString(names[1], result.Data[1])
    oxyde.AssertEqualString(names[2], result.Data[2])
    c.DisplayOK(ctx)
}

func TcDirectoryListMultipleDeep(ctx *c.Context, dtx *oxyde.DocContext) {
    const summary = `Listing the content of root directory with multiple subdirectories.`
    const description = `If there is more than one subdirectory, all of them should be returned in the directory list.`
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    directoryNameA := c.RootDirName + c.DirectoryNames[c.DirectoryA]
    directoryNameB := directoryNameA + "/" + c.DirectoryNames[c.DirectoryB]
    directoryNameC := directoryNameB + "/" + c.DirectoryNames[c.DirectoryC]
    DirectoryCreate(ctx, dtx, directoryNameC, true)
    params := DirectoryListParams{Name: &c.RootDirName}
    var result DirectoryListResult
    dtx.CollectUsage(summary, description)
    oxyde.HttpGET(ctx, dtx, DirectoryListUrl, nil, &params, &result, 200)
    oxyde.AssertNotNil(result.Data)
    oxyde.AssertEqualInt(3, len(result.Data))
    oxyde.AssertEqualString(directoryNameA, result.Data[0])
    oxyde.AssertEqualString(directoryNameB, result.Data[1])
    oxyde.AssertEqualString(directoryNameC, result.Data[2])
    c.DisplayOK(ctx)
}

func TcDirectoryListNoPermissions(ctx *c.Context, dtx *oxyde.DocContext) {
    const summary = `Listing the content of root directory without permissions.`
    const description = `If there is a directory without required permissions in the tree, then error should be returned.`
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    directoryNameA := c.RootDirName + c.DirectoryNames[c.DirectoryA]
    DirectoryCreate(ctx, dtx, directoryNameA, true)
    directoryNameB := directoryNameA + "/" + c.DirectoryNames[c.DirectoryB]
    DirectoryCreate(ctx, dtx, directoryNameB, false)
    c.ChangeMode(directoryNameB, 222)
    params := DirectoryListParams{Name: &c.RootDirName}
    var result DirectoryListResult
    dtx.CollectUsage(summary, description)
    oxyde.HttpGET(ctx, dtx, DirectoryListUrl, nil, &params, &result, 400)
    c.AssertError(result.Errors, 400, "walking directory tree failed", "check server log for details", 10177)
    c.ChangeMode(directoryNameB, 755)
    c.DisplayOK(ctx)
}

func DirectoryList(ctx *c.Context, dtx *oxyde.DocContext, name string) []string {
    params := DirectoryListParams{Name: &name}
    var result DirectoryListResult
    oxyde.HttpGET(ctx, dtx, DirectoryListUrl, nil, &params, &result, 200)
    return result.Data
}

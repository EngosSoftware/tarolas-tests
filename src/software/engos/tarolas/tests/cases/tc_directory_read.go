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
    DirectoryReadUrl = "/directory/read"
)

type DirectoryReadResult struct {
    Data   c.DirectoryDto `json:"data"`
    Errors []c.ErrorDto   `json:"errors"`
}

type DirectoryReadParams struct {
    Name *string `json:"name"  api:"Name of the directory to be read."`
}

func TsDirectoryRead(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    TdDirectoryRead(ctx, dtx)
}

func TdDirectoryRead(ctx *c.Context, dtx *oxyde.DocContext) {
    const summary = `Reads the directory content.`
    const description = `(to be updated)`
    c.Display(ctx)
    dtx.NewEndpoint(ctx.Version, c.DirectoriesTag, summary, description)
    TcDirectoryReadRoot(ctx, dtx)
    TcDirectoryReadEmptyRoot(ctx, dtx)
    dtx.SaveEndpoint()
}

func TcDirectoryReadRoot(ctx *c.Context, dtx *oxyde.DocContext) {
    const summary = `
Reading the content of directory.
`
    const description = `
aaa
`
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    directoryName := c.RootDirName + c.DirectoryNames[c.DirectoryA]
    DirectoryCreate(ctx, dtx, directoryName, false)
    params := DirectoryReadParams{Name: &c.RootDirName}
    var result DirectoryReadResult
    dtx.CollectAll(summary, description)
    oxyde.HttpGET(ctx, dtx, DirectoryReadUrl, nil, &params, &result, 200)
    oxyde.AssertNotNil(result.Data)
    c.AssertNoFiles(result.Data)
    oxyde.AssertEqualStringNullable(&c.RootDirName, result.Data.Name)
    oxyde.AssertNotNil(result.Data.Directories)
    subDir := result.Data.Directories[0]
    oxyde.AssertNotNil(subDir.Name)
    oxyde.AssertEqualStringNullable(&c.DirectoryNames[c.DirectoryA], subDir.Name)
    c.AssertNoFiles(subDir)
    c.AssertNoDirectories(subDir)
    c.DisplayOK(ctx)
}

func TcDirectoryReadEmptyRoot(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    var result DirectoryReadResult
    f := func(params DirectoryReadParams, errorMessage string, errorCode int) {
        oxyde.HttpGET(ctx, dtx, DirectoryReadUrl, nil, &params, &result, 400)
        c.AssertError(result.Errors, 400, errorMessage, "name", errorCode)
    }
    f(DirectoryReadParams{Name: nil},"required parameter is missing",10837)
    f(DirectoryReadParams{Name: &c.EmptyValue},"required parameter is empty",10767)
    f(DirectoryReadParams{Name: &c.SpaceValue},"required parameter is empty",10767)
    f(DirectoryReadParams{Name: &c.WhiteSpaceValue},"required parameter is empty",10767)
    c.DisplayOK(ctx)
}

func DirectoryRead(ctx *c.Context, dtx *oxyde.DocContext, name string) c.DirectoryDto {
    params := DirectoryReadParams{Name: &name}
    var result DirectoryReadResult
    oxyde.HttpGET(ctx, dtx, DirectoryReadUrl, nil, &params, &result, 200)
    return result.Data
}

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
    c.Display(ctx)
    TcDirectoryReadEmptyRoot(ctx, dtx)
    TcDirectoryReadRoot(ctx, dtx)
}

func TcDirectoryReadEmptyRoot(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    var result DirectoryReadResult
    f := func(params DirectoryReadParams) {
        oxyde.HttpGET(ctx, dtx, DirectoryReadUrl, nil, &params, &result, 200)
        oxyde.AssertNotNil(result.Data)
        oxyde.AssertEqualInt(0, len(result.Data.Directories))
        oxyde.AssertEqualInt(0, len(result.Data.Files))
        oxyde.AssertEqualStringNullable(&c.RootDirName, result.Data.Name)
    }
    f(DirectoryReadParams{Name: nil})
    f(DirectoryReadParams{Name: &c.RootDirName})
    f(DirectoryReadParams{Name: &c.EmptyValue})
    f(DirectoryReadParams{Name: &c.SpaceValue})
    f(DirectoryReadParams{Name: &c.WhiteSpaceValue})
    c.DisplayOK(ctx)
}

func TcDirectoryReadRoot(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    DirectoryCreate(ctx, dtx, c.DirectoryNames[c.DirectoryA], false)
    var result DirectoryReadResult
    params := DirectoryReadParams{Name: nil}
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

func DirectoryRead(ctx *c.Context, dtx *oxyde.DocContext, name string) c.DirectoryDto {
    params := DirectoryReadParams{Name: &name}
    var result DirectoryReadResult
    oxyde.HttpGET(ctx, dtx, DirectoryReadUrl, nil, &params, &result, 200)
    return result.Data
}

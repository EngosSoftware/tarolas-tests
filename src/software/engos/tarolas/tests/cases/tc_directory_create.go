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
    DirectoryCreateUrl = "/directory/create"
)

type DirectoryCreateResult struct {
    Data   c.DirectoryDto `json:"data"`
    Errors []c.ErrorDto   `json:"errors"`
}

type DirectoryCreateParams struct {
    Name *string `json:"name"  api:"Name of the directory to be created."`
    All  *bool   `json:"all"   api:"Flag indicating if non existing parent directories should be also created."`
}

func TsDirectoryCreate(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    TdDirectoryCreate(ctx, dtx)
}

func TdDirectoryCreate(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    TcDirectoryCreateRoot(ctx, dtx)
    TcDirectoryCreateSingle(ctx, dtx)
    TcDirectoryCreateSingleInvalidParameters(ctx, dtx)
    TcDirectoryCreateSingleMethodNotSupported(ctx, dtx)
    TcDirectoryCreateMultiple(ctx, dtx)
}

func TcDirectoryCreateRoot(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    var result DirectoryCreateResult
    params := DirectoryCreateParams{
        Name: &c.RootDirName,
        All:  &c.FlagFalse}
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 400)
    c.AssertError(result.Errors, c.Http400, "directory already exists", "/", 10237)
    // root directory should still be empty
    dir := DirectoryRead(ctx, dtx, c.RootDirName)
    c.AssertEmptyDirectory(dir, c.RootDirName)
    c.DisplayOK(ctx)
}

func TcDirectoryCreateSingle(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    // create single directory
    var result DirectoryCreateResult
    params := DirectoryCreateParams{Name: &c.DirectoryNames[c.DirectoryA], All: &c.FlagFalse}
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 200)
    oxyde.AssertNotNil(result.Data)
    c.AssertNoFiles(result.Data)
    c.AssertNoDirectories(result.Data)
    oxyde.AssertEqualStringNullable(&c.DirectoryNames[c.DirectoryA], result.Data.Name)
    // root directory should contain only one directory
    dir := DirectoryRead(ctx, dtx, c.RootDirName)
    oxyde.AssertEqualStringNullable(&c.RootDirName, dir.Name)
    oxyde.AssertEqualInt(1, len(dir.Directories))
    c.AssertNoFiles(dir)
    subDir := dir.Directories[0]
    c.AssertEmptyDirectory(subDir, c.DirectoryNames[c.DirectoryA])
    c.DisplayOK(ctx)
}

func TcDirectoryCreateSingleInvalidParameters(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    var result DirectoryCreateResult
    // no parameters
    params := DirectoryCreateParams{Name: nil, All: nil}
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 400)
    c.AssertError(result.Errors, 400, "required parameter is missing", "name", 10837)
    // empty directory name
    params = DirectoryCreateParams{Name: &c.EmptyValue, All: nil}
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 400)
    c.AssertError(result.Errors, 400, "required parameter is empty", "name", 10767)
    // directory name has only spaces
    params = DirectoryCreateParams{Name: &c.SpaceValue, All: nil}
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 400)
    c.AssertError(result.Errors, 400, "required parameter is empty", "name", 10767)
    // directory name has only white spaces
    params = DirectoryCreateParams{Name: &c.WhiteSpaceValue, All: nil}
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 400)
    c.AssertError(result.Errors, 400, "required parameter is empty", "name", 10767)
    // directory name has invalid characters
    a := string([]byte{'/', 'd', 'i', 'r', 0})
    params = DirectoryCreateParams{Name: &a, All: nil}
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 400)
    c.AssertError(result.Errors, 400, "creating directory failed", a, 10152)
    // root directory should still be empty
    dir := DirectoryRead(ctx, dtx, c.RootDirName)
    c.AssertEmptyDirectory(dir, c.RootDirName)
    c.DisplayOK(ctx)
}

func TcDirectoryCreateSingleMethodNotSupported(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    var result DirectoryCreateResult
    all := false
    params := DirectoryCreateParams{
        Name: &c.DirectoryNames[c.DirectoryA],
        All:  &all}
    oxyde.HttpGET(ctx, dtx, DirectoryCreateUrl, nil, &params, &result, 400)
    c.AssertMethodNotSupportedErrorGET(result.Errors)
    oxyde.HttpPUT(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 400)
    c.AssertMethodNotSupportedErrorPUT(result.Errors)
    oxyde.HttpDELETE(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 400)
    c.AssertMethodNotSupportedErrorDELETE(result.Errors)
    // root directory should still be empty
    dir := DirectoryRead(ctx, dtx, c.RootDirName)
    c.AssertEmptyDirectory(dir, c.RootDirName)
    c.DisplayOK(ctx)
}

func TcDirectoryCreateMultiple(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    var result DirectoryCreateResult
    // create two levels
    name := c.DirectoryNames[c.DirectoryA] + "/" + c.DirectoryNames[c.DirectoryB]
    params := DirectoryCreateParams{Name: &name, All: &c.FlagTrue}
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 200)
    c.AssertEmptyDirectory(result.Data, c.DirectoryNames[c.DirectoryB])
    dir := DirectoryRead(ctx, dtx, c.RootDirName)
    c.AssertOneSubdirectory(dir, c.RootDirName, c.DirectoryNames[c.DirectoryA])
    dir = DirectoryRead(ctx, dtx, c.DirectoryNames[c.DirectoryA])
    c.AssertOneSubdirectory(dir, c.DirectoryNames[c.DirectoryA], c.DirectoryNames[c.DirectoryB])
    // create three levels
    name = c.DirectoryNames[c.DirectoryA] + "/" + c.DirectoryNames[c.DirectoryB] + "/" + c.DirectoryNames[c.DirectoryC]
    params = DirectoryCreateParams{Name: &name, All: &c.FlagTrue}
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 200)
    c.AssertEmptyDirectory(result.Data, c.DirectoryNames[c.DirectoryC])
    dir = DirectoryRead(ctx, dtx, c.RootDirName)
    c.AssertOneSubdirectory(dir, c.RootDirName, c.DirectoryNames[c.DirectoryA])
    dir = DirectoryRead(ctx, dtx, c.DirectoryNames[c.DirectoryA])
    c.AssertOneSubdirectory(dir, c.DirectoryNames[c.DirectoryA], c.DirectoryNames[c.DirectoryB])
    dir = DirectoryRead(ctx, dtx, c.DirectoryNames[c.DirectoryA]+"/"+c.DirectoryNames[c.DirectoryB])
    c.AssertOneSubdirectory(dir, c.DirectoryNames[c.DirectoryB], c.DirectoryNames[c.DirectoryC])
    c.DisplayOK(ctx)
}

func DirectoryCreate(ctx *c.Context, dtx *oxyde.DocContext, name string, all bool) {
    params := DirectoryCreateParams{Name: &name, All: &all}
    var result DirectoryCreateResult
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 200)
}

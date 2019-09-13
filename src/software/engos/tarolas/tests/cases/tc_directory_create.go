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

type DirectoryCreateParams struct {
    Name *string `json:"name"  api:"Name of the directory to be created. Directory name may contain relative path when needed."`
    All  *bool   `json:"all"   api:"?Optional flag indicating if not existing parent directories should be created when needed."`
}

type DirectoryCreateResult struct {
    Data   c.DirectoryDto `json:"data"    api:"Created directory object when the request was processed successfully."`
    Errors []c.ErrorDto   `json:"errors"  api:"Error details when processing the request failed."`
}

func TsDirectoryCreate(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    TdDirectoryCreate(ctx, dtx)
}

func TdDirectoryCreate(ctx *c.Context, dtx *oxyde.DocContext) {
    const summary = `Creates a new directory.`
    const description = `(to be updated)`
    c.Display(ctx)
    dtx.NewEndpoint(ctx.Version, c.DirectoriesTag, summary, description)
    TcDirectoryCreateSingle(ctx, dtx)
    TcDirectoryCreateMultipleOneSubdirectory(ctx, dtx)
    TcDirectoryCreateMultipleTwoSubdirectories(ctx, dtx)
    TcDirectoryCreateRoot(ctx, dtx)
    TcDirectoryCreateSingleInvalidParameters(ctx, dtx)
    TcDirectoryCreateSingleMethodNotSupported(ctx, dtx)
    dtx.SaveEndpoint()
}

func TcDirectoryCreateSingle(ctx *c.Context, dtx *oxyde.DocContext) {
    // SET-UP
    c.Display(ctx)
    // DOCUMENTATION
    const summary = `
Creating a new directory.
`
    const description = `
Creates a new, single directory that will be placed directly in the server's root directory,
because the directory name does not contain any relative path.
`
    // GIVEN
    RemoveRootContents(ctx, dtx)
    directoryName := c.RootDirName + c.DirectoryNames[c.DirectoryA]
    params := DirectoryCreateParams{Name: &directoryName}
    var result DirectoryCreateResult
    // WHEN
    dtx.CollectAll(summary, description)
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 200)
    // THEN
    oxyde.AssertNotNil(result.Data)
    c.AssertNoFiles(result.Data)
    c.AssertNoDirectories(result.Data)
    oxyde.AssertEqualStringNullable(&directoryName, result.Data.Name)
    // root directory should contain only one directory
    dir := DirectoryRead(ctx, dtx, c.RootDirName)
    oxyde.AssertEqualStringNullable(&c.RootDirName, dir.Name)
    oxyde.AssertEqualInt(1, len(dir.Directories))
    c.AssertNoFiles(dir)
    subDir := dir.Directories[0]
    c.AssertEmptyDirectory(subDir, c.DirectoryNames[c.DirectoryA])
    // TEAR-DOWN
    c.DisplayOK(ctx)
}

func TcDirectoryCreateRoot(ctx *c.Context, dtx *oxyde.DocContext) {
    const summary = `
Creating root directory is disabled.
`
    const description = `
Creating root directory has no effect. Any attempt to create the root directory ends with error message.
`
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    params := DirectoryCreateParams{Name: &c.RootDirName}
    var result DirectoryCreateResult
    dtx.CollectUsage(summary, description)
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 400)
    c.AssertError(result.Errors, c.Http400, "directory already exists", "/", 10237)
    // root directory should still be empty
    dir := DirectoryRead(ctx, dtx, c.RootDirName)
    c.AssertEmptyDirectory(dir, c.RootDirName)
    c.DisplayOK(ctx)
}

func TcDirectoryCreateSingleInvalidParameters(ctx *c.Context, dtx *oxyde.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    var result DirectoryCreateResult

    // no parameters
    dtx.CollectUsage("Attempt to create directory without specifying its name.", "")
    params := DirectoryCreateParams{Name: nil, All: nil}
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 400)
    c.AssertError(result.Errors, 400, "required parameter is missing", "name", 10837)

    // empty directory name
    dtx.CollectUsage("Attempt to create directory with empty name.", "")
    params = DirectoryCreateParams{Name: &c.EmptyValue, All: nil}
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 400)
    c.AssertError(result.Errors, 400, "required parameter is empty", "name", 10767)

    // directory name has only spaces
    dtx.CollectUsage("Attempt to create directory with name containing only spaces.", "")
    params = DirectoryCreateParams{Name: &c.SpaceValue, All: nil}
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 400)
    c.AssertError(result.Errors, 400, "required parameter is empty", "name", 10767)

    // directory name has only white spaces
    dtx.CollectUsage("Attempt to create directory with name containing white characters.", "")
    params = DirectoryCreateParams{Name: &c.WhiteSpaceValue, All: nil}
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 400)
    c.AssertError(result.Errors, 400, "required parameter is empty", "name", 10767)

    // directory name has invalid characters
    dtx.CollectUsage("Attempt to create directory with invalid characters.", "")
    a := string([]byte{'/', 'j', 'u', 'n', 0, 'g', 'l', 'e'})
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
    params := DirectoryCreateParams{Name: &c.DirectoryNames[c.DirectoryA], All: &c.FlagTrue}

    dtx.CollectUsage("Attempt to create directory with GET request.", "")
    oxyde.HttpGET(ctx, dtx, DirectoryCreateUrl, nil, &params, &result, 400)
    c.AssertMethodNotSupportedErrorGET(result.Errors)

    dtx.CollectUsage("Attempt to create directory with PUT request.", "")
    oxyde.HttpPUT(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 400)
    c.AssertMethodNotSupportedErrorPUT(result.Errors)

    dtx.CollectUsage("Attempt to create directory with DELETE request.", "")
    oxyde.HttpDELETE(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 400)
    c.AssertMethodNotSupportedErrorDELETE(result.Errors)

    // root directory should still be empty
    dir := DirectoryRead(ctx, dtx, c.RootDirName)
    c.AssertEmptyDirectory(dir, c.RootDirName)
    c.DisplayOK(ctx)
}

func TcDirectoryCreateMultipleOneSubdirectory(ctx *c.Context, dtx *oxyde.DocContext) {
    const summary = `
Creating a new directory with parent directories (one subdirectory).
`
    const description = `
Creates a new, single directory that will be placed in subdirectory. 
Because the subdirectory does not exist, it will be created (**all** flag is set to **true**).
`
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    var result DirectoryCreateResult
    // create two levels
    directoryName := c.RootDirName + c.DirectoryNames[c.DirectoryA] + "/" + c.DirectoryNames[c.DirectoryB]
    params := DirectoryCreateParams{Name: &directoryName, All: &c.FlagTrue}
    dtx.CollectUsage(summary, description)
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 200)
    c.AssertEmptyDirectory(result.Data, c.DirectoryNames[c.DirectoryB])
    dir := DirectoryRead(ctx, dtx, c.RootDirName)
    c.AssertOneSubdirectory(dir, c.RootDirName, c.DirectoryNames[c.DirectoryA])
    dir = DirectoryRead(ctx, dtx, c.RootDirName+c.DirectoryNames[c.DirectoryA])
    c.AssertOneSubdirectory(dir, c.DirectoryNames[c.DirectoryA], c.DirectoryNames[c.DirectoryB])
    c.DisplayOK(ctx)
}

func TcDirectoryCreateMultipleTwoSubdirectories(ctx *c.Context, dtx *oxyde.DocContext) {
    const summary = `
Creating a new directory with parent directories (two subdirectories).
`
    const description = `
Creates a new, single directory that will be placed in subdirectories. 
Because the subdirectories do not exist, they will be created (**all** flag is set to **true**).
`
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    var result DirectoryCreateResult
    // create three levels
    name := c.RootDirName + c.DirectoryNames[c.DirectoryA] + "/" + c.DirectoryNames[c.DirectoryB] + "/" + c.DirectoryNames[c.DirectoryC]
    params := DirectoryCreateParams{Name: &name, All: &c.FlagTrue}
    dtx.CollectUsage(summary, description)
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 200)
    c.AssertEmptyDirectory(result.Data, c.DirectoryNames[c.DirectoryC])
    dir := DirectoryRead(ctx, dtx, c.RootDirName)
    c.AssertOneSubdirectory(dir, c.RootDirName, c.DirectoryNames[c.DirectoryA])
    dir = DirectoryRead(ctx, dtx, c.RootDirName+c.DirectoryNames[c.DirectoryA])
    c.AssertOneSubdirectory(dir, c.DirectoryNames[c.DirectoryA], c.DirectoryNames[c.DirectoryB])
    dir = DirectoryRead(ctx, dtx, c.RootDirName+c.DirectoryNames[c.DirectoryA]+"/"+c.DirectoryNames[c.DirectoryB])
    c.AssertOneSubdirectory(dir, c.DirectoryNames[c.DirectoryB], c.DirectoryNames[c.DirectoryC])
    c.DisplayOK(ctx)
}

func DirectoryCreate(ctx *c.Context, dtx *oxyde.DocContext, name string, all bool) {
    params := DirectoryCreateParams{Name: &name, All: &all}
    var result DirectoryCreateResult
    oxyde.HttpPOST(ctx, dtx, DirectoryCreateUrl, nil, &params, nil, &result, 200)
}

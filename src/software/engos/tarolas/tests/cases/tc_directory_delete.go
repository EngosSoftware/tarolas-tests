/*
 * MIT License
 *
 * Copyright (c) 2019 Dariusz Depta Engos Software
 *
 * License details in LICENSE.
 */

package cases

import (
    o "github.com/EngosSoftware/oxyde"
    c "software/engos/tarolas/tests/common"
)

const (
    DirectoryDeleteUrl = "/directory/delete"
)

type DirectoryDeleteParams struct {
    Name *string `json:"name"  api:"Name of the directory to be deleted. Directory name may contain parent directories."`
    All  *bool   `json:"all"   api:"?Optional flag indicating if whole directory content should be deleted, including subdirectories and files."`
}

type DirectoryDeleteResult struct {
    Data   c.DirectoryDto `json:"data"`
    Errors []c.ErrorDto   `json:"errors"`
}

func TsDirectoryDelete(ctx *c.Context, dtx *o.DocContext) {
    c.Display(ctx)
    TdDirectoryDelete(ctx, dtx)
}

func TdDirectoryDelete(ctx *c.Context, dtx *o.DocContext) {
    const summary = `Deletes an existing directory.`
    const description = `(to be updated)`
    c.Display(ctx)
    dtx.NewEndpoint(ctx.Version, c.DirectoriesTag, summary, description)
    TcDirectoryDelete(ctx, dtx)
    TcDirectoryDeleteRoot(ctx, dtx)
    TcDirectoryDeleteSubdirectory(ctx, dtx)
    TcDirectoryDeleteMethodNotSupported(ctx, dtx)
    dtx.SaveEndpoint()
}

func TcDirectoryDelete(ctx *c.Context, dtx *o.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    directoryName := c.RootDirName + c.DirectoryNames[c.DirectoryA]
    DirectoryCreate(ctx, dtx, directoryName, c.FlagFalse)
    params := DirectoryDeleteParams{Name: &directoryName, All: &c.FlagFalse}
    var result DirectoryDeleteResult
    dtx.CollectAll("Deletes single directory", "")
    o.HttpDELETE(ctx, dtx, DirectoryDeleteUrl, nil, &params, nil, &result, 200)
    dir := DirectoryRead(ctx, dtx, c.RootDirName)
    c.AssertEmptyDirectory(dir, c.RootDirName)
    c.DisplayOK(ctx)
}

func TcDirectoryDeleteRoot(ctx *c.Context, dtx *o.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    var result DirectoryDeleteResult
    // with flag 'all' == false deleting empty root directory should have no effect
    params := DirectoryDeleteParams{Name: &c.RootDirName, All: &c.FlagFalse}
    o.HttpDELETE(ctx, dtx, DirectoryDeleteUrl, nil, &params, nil, &result, 200)
    dir := DirectoryRead(ctx, dtx, c.RootDirName)
    c.AssertEmptyDirectory(dir, c.RootDirName)
    // with flag 'all' == true deleting empty root directory should have no effect
    params = DirectoryDeleteParams{Name: &c.RootDirName, All: &c.FlagTrue}
    o.HttpDELETE(ctx, dtx, DirectoryDeleteUrl, nil, &params, nil, &result, 200)
    dir = DirectoryRead(ctx, dtx, c.RootDirName)
    c.AssertEmptyDirectory(dir, c.RootDirName)
    // with flag 'all' == false deleting non empty root directory should have no effect
    DirectoryCreate(ctx, dtx, c.RootDirName+c.DirectoryNames[c.DirectoryA], c.FlagFalse)
    params = DirectoryDeleteParams{Name: &c.RootDirName, All: &c.FlagFalse}
    o.HttpDELETE(ctx, dtx, DirectoryDeleteUrl, nil, &params, nil, &result, 200)
    dir = DirectoryRead(ctx, dtx, c.RootDirName)
    c.AssertOneSubdirectory(dir, c.RootDirName, c.DirectoryNames[c.DirectoryA])
    // with flag 'all' == true deleting non empty root directory should delete the content
    params = DirectoryDeleteParams{Name: &c.RootDirName, All: &c.FlagTrue}
    o.HttpDELETE(ctx, dtx, DirectoryDeleteUrl, nil, &params, nil, &result, 200)
    dir = DirectoryRead(ctx, dtx, c.RootDirName)
    c.AssertEmptyDirectory(dir, c.RootDirName)
    c.DisplayOK(ctx)
}

func TcDirectoryDeleteSubdirectory(ctx *c.Context, dtx *o.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    name := c.RootDirName + c.DirectoryNames[c.DirectoryA] + "/" + c.DirectoryNames[c.DirectoryB]
    DirectoryCreate(ctx, dtx, name, c.FlagTrue)
    var result DirectoryDeleteResult
    params := DirectoryDeleteParams{Name: &name, All: &c.FlagFalse}
    o.HttpDELETE(ctx, dtx, DirectoryDeleteUrl, nil, &params, nil, &result, 200)
    dir := DirectoryRead(ctx, dtx, c.RootDirName+c.DirectoryNames[c.DirectoryA])
    c.AssertEmptyDirectory(dir, c.DirectoryNames[c.DirectoryA])
    c.DisplayOK(ctx)
}

func TcDirectoryDeleteMethodNotSupported(ctx *c.Context, dtx *o.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    DirectoryCreate(ctx, dtx, c.RootDirName+c.DirectoryNames[c.DirectoryA], false)
    var result DirectoryDeleteResult
    all := false
    params := DirectoryDeleteParams{
        Name: &c.DirectoryNames[c.DirectoryA],
        All:  &all}
    o.HttpGET(ctx, dtx, DirectoryDeleteUrl, nil, &params, &result, 400)
    c.AssertMethodNotSupportedErrorGET(result.Errors)
    o.HttpPOST(ctx, dtx, DirectoryDeleteUrl, nil, &params, nil, &result, 400)
    c.AssertMethodNotSupportedErrorPOST(result.Errors)
    o.HttpPUT(ctx, dtx, DirectoryDeleteUrl, nil, &params, nil, &result, 400)
    c.AssertMethodNotSupportedErrorPUT(result.Errors)
    c.DisplayOK(ctx)
}

func DirectoryDelete(ctx *c.Context, dtx *o.DocContext, name string, all bool) {
    var result DirectoryDeleteResult
    params := DirectoryDeleteParams{Name: &name, All: &all}
    o.HttpDELETE(ctx, dtx, DirectoryDeleteUrl, nil, &params, nil, &result, 200)
}

func RemoveRootContents(ctx *c.Context, dtx *o.DocContext) {
    DirectoryDelete(ctx, dtx, "/", true)
}

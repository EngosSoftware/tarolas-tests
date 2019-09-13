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
    FileAppendUrl = "/file/append"
)

type FileAppendBody struct {
    Base64 *string `json:"-"  api:"File content to be appended, Base64 encoded."`
}

type FileAppendParams struct {
    Name *string `json:"name"  api:"Name of the file to be appended."`
}

type FileAppendResult struct {
    Data   c.FileDto    `json:"data"`
    Errors []c.ErrorDto `json:"errors"`
}

func TsFileAppend(ctx *c.Context, dtx *o.DocContext) {
    c.Display(ctx)
    TdFileAppend(ctx, dtx)
}

func TdFileAppend(ctx *c.Context, dtx *o.DocContext) {
    const summary = `Appends data to a file.`
    const description = `(to be updated)`
    c.Display(ctx)
    dtx.NewEndpoint(ctx.Version, c.FilesTag, summary, description)
    TcFileAppend(ctx, dtx)
    TcFileAppendParts(ctx, dtx)
    dtx.SaveEndpoint()
}

func TcFileAppend(ctx *c.Context, dtx *o.DocContext) {
    const summary = `
Appends data to file.
`
    const description = `
`
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    fileName := c.RootDirName + c.FileNames[c.FileA]
    params := FileAppendParams{Name: &fileName}
    content := c.EncodeToString(c.FileContents[c.FileA])
    body := FileAppendBody{Base64: &content}
    var result FileAppendResult
    dtx.CollectAll(summary, description)
    o.HttpPUT(ctx, dtx, FileAppendUrl, nil, &params, &body, &result, 200)
    c.DisplayOK(ctx)
}

func TcFileAppendParts(ctx *c.Context, dtx *o.DocContext) {
    c.Display(ctx)
    RemoveRootContents(ctx, dtx)
    var result FileAppendResult
    fileName := c.RootDirName + c.FileNames[c.FileB]
    params := FileAppendParams{Name: &fileName}

    content := c.EncodeToString(c.FileContents[c.FileA])
    body := FileAppendBody{Base64: &content}
    o.HttpPUT(ctx, dtx, FileAppendUrl, nil, &params, &body, &result, 200)

    content = c.EncodeToString(c.FileContents[c.FileB])
    body = FileAppendBody{Base64: &content}
    o.HttpPUT(ctx, dtx, FileAppendUrl, nil, &params, &body, &result, 200)

    content = c.EncodeToString(c.FileContents[c.FileC])
    body = FileAppendBody{Base64: &content}
    o.HttpPUT(ctx, dtx, FileAppendUrl, nil, &params, &body, &result, 200)

    content = c.EncodeToString(c.FileContents[c.FileD])
    body = FileAppendBody{Base64: &content}
    o.HttpPUT(ctx, dtx, FileAppendUrl, nil, &params, &body, &result, 200)

    c.DisplayOK(ctx)
}

func FileAppend(ctx *c.Context, dtx *o.DocContext, name string, content []byte) {
    fileContent := c.EncodeToStringBytes(content)
    params := FileAppendParams{Name: &name}
    body := FileAppendBody{Base64: &fileContent}
    var result FileAppendResult
    o.HttpPUT(ctx, dtx, FileAppendUrl, nil, &params, &body, &result, 200)
}

func FileAppendBig(ctx *c.Context, dtx *o.DocContext, name string, len, count int) {
    params := FileAppendParams{Name: &name}
    var result FileAppendResult
    for i := 0; i < count; i++ {
        fileContent := c.EncodeToStringBytes(c.RandomContent(len))
        body := FileAppendBody{Base64: &fileContent}
        o.HttpPUT(ctx, dtx, FileAppendUrl, nil, &params, &body, &result, 200)
    }
}

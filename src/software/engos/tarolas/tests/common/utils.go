/*
 * MIT License
 *
 * Copyright (c) 2019 Dariusz Depta Engos Software
 *
 * License details in LICENSE.
 */

package common

import (
    "bytes"
    "fmt"
    o "github.com/EngosSoftware/oxyde"
    "math/rand"
    "strconv"
    "strings"
)

// Context struct contains all contextual info needed to execute test.
type Context struct {
    Url     string // Address of tested endpoint.
    Verbose bool   // Flag indicating if test execution details should be displayed.
    Version string // API version.
}

func NewContext(url string, verbose bool) *Context {
    return &Context{
        Url:     url,
        Verbose: verbose,
        Version: ""}
}

func (c *Context) GetUrl() string {
    return c.Url
}

func (c *Context) GetVerbose() bool {
    return c.Verbose
}

func (c *Context) GetVersion() string {
    return c.Version
}

func Display(ctx *Context) {
    DisplayLevel(ctx, 3)
}

func DisplayLevel(ctx *Context, level int) {
    const format = "%-110s %-5s"
    text := o.FunctionName(level)
    if strings.HasPrefix(text, "ts_") {
        fmt.Printf(format+"\n", " >> "+text, ctx.Version)
    } else if strings.HasPrefix(text, "td_") {
        fmt.Printf(format+"\n", "    > "+text, ctx.Version)
    } else if strings.HasPrefix(text, "tc_") {
        fmt.Printf(format, "      * "+text, ctx.Version)
    } else {
        fmt.Printf(format+"\n", ">>> "+text, ctx.Version)
    }
}

func DisplayOK(ctx *Context) {
    fmt.Println("OK")
}

func RandomContent(len int) []byte {
    var buffer bytes.Buffer
    width := 0
    for i := 0; i < len; i++ {
        if width == 80 {
            buffer.WriteByte('\n')
            width = 0
        } else {
            buffer.WriteByte(byte(65 + rand.Intn(25)))
        }
        width++
    }
    return buffer.Bytes()
}

func AssertOneSubdirectory(dir DirectoryDto, name string, subName string) {
    o.AssertEqualStringNullable(&name, dir.Name)
    o.AssertEqualInt(1, len(dir.Directories))
    AssertNoFiles(dir)
    subDir := dir.Directories[0]
    o.AssertEqualStringNullable(&subName, subDir.Name)
    AssertNoDirectories(subDir)
    AssertNoFiles(subDir)
}

func AssertEmptyDirectory(dir DirectoryDto, name string) {
    o.AssertEqualStringNullable(&name, dir.Name)
    AssertNoDirectories(dir)
    AssertNoFiles(dir)
}

func AssertNoDirectories(dir DirectoryDto) {
    o.AssertEqualInt(0, len(dir.Directories))
}

func AssertNoFiles(dir DirectoryDto) {
    o.AssertEqualInt(0, len(dir.Files))
}

func AssertError(errors []ErrorDto, status int, title string, detail string, code int) {
    o.AssertEqualInt(1, len(errors))
    s := strconv.Itoa(status)
    o.AssertEqualStringNullable(&s, errors[0].Status)
    o.AssertEqualStringNullable(&title, errors[0].Title)
    o.AssertEqualStringNullable(&detail, errors[0].Detail)
    c := strconv.Itoa(code)
    o.AssertEqualStringNullable(&c, errors[0].Code)
}

func AssertMethodNotSupportedErrorGET(errors []ErrorDto) {
    AssertError(errors, Http400, "request method not supported", "GET", 10001)
}

func AssertMethodNotSupportedErrorPOST(errors []ErrorDto) {
    AssertError(errors, Http400, "request method not supported", "POST", 10001)
}

func AssertMethodNotSupportedErrorPUT(errors []ErrorDto) {
    AssertError(errors, Http400, "request method not supported", "PUT", 10001)
}

func AssertMethodNotSupportedErrorDELETE(errors []ErrorDto) {
    AssertError(errors, Http400, "request method not supported", "DELETE", 10001)
}

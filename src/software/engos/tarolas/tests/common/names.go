/*
 * MIT License
 *
 * Copyright (c) 2019 Dariusz Depta Engos Software
 *
 * License details in LICENSE.
 */

package common

import (
    o "github.com/EngosSoftware/oxyde"
)

const (
    DirectoriesTag = "Directories"
    FilesTag       = "Files"
)

var (
    RootDirName     = "/"
    FlagFalse       = false
    FlagTrue        = true
    EmptyValue      = ""
    SpaceValue      = "     "
    WhiteSpaceValue = "   \n   \t    "
    Http400         = 400
)

const (
    DirectoryA = iota
    DirectoryB
    DirectoryC
    DirectoryD
)

var (
    DirectoryCount = 4
    DirectoryNames = []string{"books", "letters", "music", "photos"}
)

const (
    FileA = iota
    FileB
    FileC
    FileD
)

var (
    FileCount    = 4
    FileNames    = []string{"file_a.txt", "file_B.txt", "FILE_C.dat", "FILE_D.dat"}
    FileContents = []string{
        `(1) Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean sit amet rhoncus purus. Quisque venenatis, eros eget rutrum consectetur, risus felis sollicitudin leo, eu varius velit justo non lacus. Interdum et malesuada fames ac ante ipsum primis in faucibus. Nulla vel nibh eros. Nulla eu eleifend eros. Quisque consectetur, mauris vel tempor ultrices, ex lacus volutpat augue, sit amet sagittis dolor tellus in odio. Pellentesque convallis libero et consectetur mollis. Sed maximus at est non imperdiet. Ut in mi a dolor eleifend facilisis. Donec posuere, nibh at dictum sodales, erat libero suscipit est, in pharetra sem eros non velit. Nam sed nulla nec mauris rhoncus bibendum sed nec augue. Integer tincidunt ante ligula, non lacinia velit tempus ut. In sagittis, enim in dapibus pulvinar, tortor nibh ornare ante, ac dignissim turpis est sed nibh [...]`,
        `(2) Donec at facilisis felis, in fringilla odio. Etiam gravida lectus et mauris pretium, a convallis nunc sollicitudin. Curabitur efficitur viverra massa, eu dictum sapien consectetur ut. Duis ut tellus quis nunc posuere ullamcorper. Donec pellentesque nunc mauris, et interdum libero blandit et. Aenean ut sollicitudin nisl, maximus faucibus massa. Fusce eleifend pretium nibh. Etiam eget urna risus. Aliquam erat volutpat. Aliquam et tortor et augue faucibus dictum [...]`,
        `(3) Maecenas ultricies egestas feugiat. Sed justo urna, placerat ut pellentesque vel, efficitur consectetur neque. Nulla turpis nibh, hendrerit in risus eu, varius sodales nibh. Cras ligula orci, tristique tincidunt ex sed, cursus laoreet dui. Donec luctus tristique est vitae dictum. Etiam malesuada leo libero, sed consectetur nisi mattis nec. Duis sed nulla vestibulum, egestas felis eget, luctus leo. Sed augue nunc, convallis eu erat a, semper vulputate tortor [...]`,
        `(4) Vestibulum eu est sit amet enim congue venenatis id at elit. In porttitor dapibus odio, eget tempor sapien aliquam tempor. Praesent sit amet egestas orci. Nunc non varius nulla. Cras blandit, metus at blandit fermentum, tellus ipsum vehicula ipsum, sed eleifend tortor quam eu nisl. Aenean id diam porta, volutpat ligula at, finibus nibh. Proin urna sem, lobortis nec diam vel, consequat laoreet nunc. Vestibulum maximus tortor nec dui iaculis, luctus vestibulum leo porttitor. Integer eu nibh in nulla ornare rhoncus. Ut nec sodales sem. Morbi rutrum sapien ut enim mattis, non suscipit magna placerat. Suspendisse turpis dolor, gravida eget enim at, pharetra feugiat urna. Nunc enim ligula, consectetur porta consequat et, elementum a mauris [...]`}
    FileSizes = []int64{
        int64(len(FileContents[FileA])),
        int64(len(FileContents[FileB])),
        int64(len(FileContents[FileC])),
        int64(len(FileContents[FileD]))}
)

func TsNames(ctx *Context, dtx *o.DocContext) {
    o.AssertEqualString("/", RootDirName)
    o.AssertEqualInt(DirectoryCount, len(DirectoryNames))
    o.AssertEqualInt(FileCount, len(FileNames))
    o.AssertEqualInt(FileCount, len(FileContents))
    o.AssertEqualInt(FileCount, len(FileSizes))
}

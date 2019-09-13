/*
 * MIT License
 *
 * Copyright (c) 2019 Dariusz Depta Engos Software
 *
 * License details in LICENSE.
 */

package common

type DirectoryDto struct {
    Name        *string        `json:"name"         api:"Name of the directory without parent path."`
    Directories []DirectoryDto `json:"directories"  api:"Child directories in this directory."`
    Files       []FileDto      `json:"files"        api:"List of files in this directory."`
}

type FileDto struct {
    Name     *string `json:"name"      api:"?File name with extension (if present) without path to file."`
    Size     *int64  `json:"size"      api:"?File size in bytes."`
    Checksum *string `json:"checksum"  api:"?File checksum (SHA256)."`
    Exists   *bool   `json:"exists"    api:"?Flag indicating if a file exists."`
}

type ErrorDto struct {
    Status *string `json:"status"  api:"The HTTP status code applicable to this problem."`
    Code   *string `json:"code"    api:"An application-specific error code."`
    Title  *string `json:"title"   api:"A short, human-readable summary of the problem that SHOULD NOT change from occurrence to occurrence of the problem."`
    Detail *string `json:"detail"  api:"A human-readable explanation specific to this occurrence of the problem."`
}

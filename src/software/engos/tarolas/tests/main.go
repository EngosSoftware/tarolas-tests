/*
 * MIT License
 *
 * Copyright (c) 2019 Dariusz Depta Engos Software
 *
 * License details in LICENSE.
 */

package main

import (
    "fmt"
    o "github.com/EngosSoftware/oxyde"
    "os"
    c "software/engos/tarolas/tests/common"
    "software/engos/tarolas/tests/suites"
)

const Url = "http://localhost:15001"

func showUsage() {
    // TODO extend usage message
    fmt.Printf("TODO....\n")
}

// Function main is a starting point for running tests.
func main() {
    if len(os.Args) == 2 {
        switch os.Args[1] {
        case "ALL":
            ctx := c.NewContext(Url, false)
            dtx := o.CreateDocContext()
            suites.All(ctx, dtx)
        case "SINGLE":
            ctx := c.NewContext(Url, false)
            dtx := o.CreateDocContext()
            suites.Single(ctx, dtx)
        }
    } else {
        showUsage()
    }
}
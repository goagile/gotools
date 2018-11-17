package main

import (
    "fmt"
    "log"

    "github.com/sajari/docconv"
)

func main() {
    res, err := docconv.ConvertPath("Fpdf_AddPage.pdf")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(res)
}

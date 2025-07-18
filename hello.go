package main

import "github.com/fatih/color"

func main() {
    c := color.New(color.FgCyan).Add(color.Underline).Add(color.BgMagenta)
    c.Println("Dependency graph rulez!")
}

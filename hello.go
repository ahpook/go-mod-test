package main

import ( 
    "github.com/fatih/color"
    "github.com/briandowns/spinner"
    "time"
)

func main() {
    
    // Create a new spinner with a specific character
    s := spinner.New(spinner.CharSets[27], 100*time.Millisecond)
    s.Start()
    time.Sleep(2 * time.Second)
    s.Stop()

    c := color.New(color.FgCyan).Add(color.Underline).Add(color.BgMagenta)
    c.Println("Dependency graph rulez!")
}

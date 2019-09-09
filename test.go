package main

import "os"

func main() {
    f, _ := os.Open("1.txt")
    defer f.Close()
}

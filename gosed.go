package main
import (
    "fmt"
    "os"
    "regexp"
    "io/ioutil"
)

func main() {

    if len(os.Args) == 2 && os.Args[1] == "--help"{
        fmt.Println("inline text replacement utility")
        fmt.Println("usage: gosed <regex> <replacement> <fileName>")
        fmt.Println("replaces all matches of regex in fileName with replacement")
        fmt.Println("note, not suitable for large files as this loads the entire file into memory")
        fmt.Println("supports go regex and expansions https://golang.org/pkg/regexp/syntax/")
        os.Exit(0)
    }

    if len(os.Args) <= 3 {
        fmt.Println("Not enough args")
        fmt.Println("try --help")
        os.Exit(0xAC1D)
    }

    regex := os.Args[1]
    repl := os.Args[2]
    fileName := os.Args[3]
    fmt.Println(regex)
    fmt.Println(repl)
    fmt.Println(fileName)

    rex := regexp.MustCompile(regex)

    fileContent, err := ioutil.ReadFile(fileName)
    check(err)

    replaced := rex.ReplaceAllString(string(fileContent), repl)

    err = ioutil.WriteFile(fileName, []byte(replaced), 0666)

    check(err)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

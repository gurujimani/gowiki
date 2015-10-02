package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    files, err := ioutil.ReadDir("./*.txt")
    if err != nil {
      fmt.Println(err)
      return
    }
    for _, f := range files {
            name := f.Name()
//	    fmt.Printf("%s = %d\n", name, length)
            fmt.Println(name[:(len(name)- 3)])
    }
}

package main

import (
    "log"
    "os"
    "net/http"
    "fmt"
    "path/filepath"
)

func main() {

    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if len(os.Args) < 2 {
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Printf("* Usage : %s [site-url]\n", dir)
        }
        return
    }

    url := os.Args[1]
    requestWakeup(dir, url)
}

func requestWakeup(dir, url string) {

    // set log file
    fpLog, err := os.OpenFile(dir + "/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        panic(err)
    }
    
    defer fpLog.Close()

    // set output
    log.SetOutput(fpLog)

    // request wake up
    if _, err := http.Get(url); err != nil {
        log.Println("Request wake-up : %s", err)
    } else {
        log.Println("Request wake-up : success")
    }
}

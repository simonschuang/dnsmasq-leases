package main

import (
       "fmt"
       "bufio"
       "os"
       "log"
       "strings"
       "encoding/json"
       "time"
       "strconv"
       "flag"
)

type LeaseData struct {
    Timestamp string
    Mac string
    Ip  string
    Host string
    Id string
}

type LeaseArray struct {
    Lease []LeaseData
}

func main() {
    var l LeaseArray

    // url := flag.String("url", "http://192.168.1.1/api", "URL to send Json format lease file")
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
        fmt.Fprintf(os.Stderr, "%s <lease file> [options]\n", os.Args[0])
        flag.PrintDefaults()
    }
    flag.Parse()
    if flag.NArg() < 1 {
        return
    }

    f, err := os.Open(flag.Arg(0))
    if err != nil {
        return
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()

        words := strings.Fields(line)
        t, _ := strconv.ParseInt(words[0], 10, 64)
        tmp := LeaseData{time.Unix(t, 0).Format(time.RFC3339),
                         words[1],
                         words[2],
                         words[3],
                         words[4]}
        l.Lease = append(l.Lease, tmp)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    b, err := json.MarshalIndent(l, "", "    ")
    if err != nil {
        fmt.Println("json err:", err)
    }
    fmt.Println(string(b))
}

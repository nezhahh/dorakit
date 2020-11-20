package fio

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "regexp"
)

type FioCmdArgs struct {
    in      string
    out     string
    print   bool
}

type fioResult struct {
    ioengine    string
    iodepth     uint32
    iops        uint32
    bandwidth   uint64
    latency     map[uint32]uint64
}

func ParseFio(args FioCmdArgs) {
    // 先判断下in是文件还是文件夹

}

func ParseDir(d string) {
}


func ParseFile(f string) {
    file, err := os.Open(f)
    if err != nil {
        log.Fatalf("file %s error(%v)", f, err)
        return
    }

    defer file.Close()

    buf, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatalf("read %s error(%v)", f, err)
        return
    }

    ParseContext(string(buf))
}

func ParseContext(c string) {
    reg1 := regexp.MustCompile(".\\n")

    result1 := reg1.FindAllStringSubmatch(c, -1)
    result2 := reg1.FindStringSubmatch(c)
    reg1.FindStringSubmatch(c)
    fmt.Println(result1)
    for _, r := range result1 {
        fmt.Println(r)
    }
}
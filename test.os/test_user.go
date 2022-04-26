// user包
// 获取 linux 用户相关信息;
package main

import (
    "fmt"
    "os/user"
)

func main() {
    u, err := user.Current()
    if err != nil {
        fmt.Println("error")
    }
    fmt.Printf("u.Uid: %s, u.Gid: %s, u.Name: %s, u.HomeDir: %s, u.Username: %s\n",
        u.Uid, u.Gid, u.Name, u.HomeDir, u.Username)

    lu, err := user.Lookup("wohu")
    if err != nil {
        fmt.Println("error")
    }
    fmt.Printf("lu.Uid: %s, lu.Gid: %s, lu.Name: %s, lu.HomeDir: %s, lu.Username: %s\n",
        lu.Uid, lu.Gid, lu.Name, lu.HomeDir, lu.Username)

    li, err := user.LookupId("1000")
    if err != nil {
        fmt.Println("error")
    }
    fmt.Printf("li.Uid: %s, li.Gid: %s, li.Name: %s, li.HomeDir: %s, li.Username: %s\n",
        li.Uid, li.Gid, li.Name, li.HomeDir, li.Username)

    lg, err := user.LookupGroup("wohu")
    if err != nil {
        fmt.Println("error")
    }
    fmt.Printf("lg.Gid: %s, lg.Name: %s\n",
        lg.Gid, lg.Name)

    lgi, err := user.LookupGroupId("1000")
    if err != nil {
        fmt.Println("error")
    }
    fmt.Printf("lgi.Gid: %s, lgi.Name: %s\n",
        lgi.Gid, lgi.Name)

}

package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    // 判断字符串 s 是否以 prefix 开头
    var str string = "This is, an example of a string "
    hasprefix := strings.HasPrefix(str, "Th")
    fmt.Println(hasprefix)

    // 判断字符串 s 是否以 suffix 结尾
    suffix := strings.HasSuffix(str, "ig")
    fmt.Println(suffix)

    // 判断字符串 s 是否包含 substr
    constains := strings.Contains(str, "wxc")
    fmt.Println(constains)

    // 判断子字符串或字符在父字符串中出现的位置
    index := strings.Index(str, "a")
    fmt.Println(index)

    // 返回字符串 str 在字符串 s 中最后出现位置
    last_index := strings.LastIndex(str, "ex")
    fmt.Println(last_index)

    // 如果 ch 是非 ASCII 编码的字符，建议使用以下函数来对字符进行定位
    index_rune := strings.IndexRune(str, 'a')
    fmt.Println(index_rune)

    // 字符串替换
    replacer := strings.Replace(str, "is", "wxc", 1)
    fmt.Println(replacer)

    // 字符串统计
    count := strings.Count(str, "is")
    fmt.Println(count)

    // 字符串重复多少次
    repeat := strings.Repeat(str, 3)
    fmt.Println(repeat)

    // 小写
    tolower := strings.ToLower(str)
    fmt.Println(tolower)

    // 大写
    toupper := strings.ToUpper(str)
    fmt.Print(toupper)

    // 剔除开头结尾空白字符
    trim_space := strings.TrimSpace(str)
    fmt.Println(trim_space)

    // 剔除开头结尾指定字符
    trim := strings.Trim(str, "Th")
    fmt.Println(trim)

    // 剔除开头指定字符
    trim_left := strings.TrimLeft(str, "Th")
    fmt.Println(trim_left)

    // 分割字符串 空白符
    fields := strings.Fields(str)
    fmt.Println(fields[3])

    // 分割指定字符
    split := strings.Split(str, ",")
    fmt.Println(split[1])

    // 拼接字符
    join := strings.Join(split, ",")
    fmt.Println(join)

    // 从字符串读取内容指针
    new_reader := strings.NewReader(str)
    fmt.Println("new_reader字符串指针: ", new_reader)

    var orig string = "666"

    // 字符串转int
    atoi, _ := strconv.Atoi(orig)
    fmt.Printf("%T atoi: %d\n", atoi, atoi)

    // 数字变成字符串
    d := strconv.Itoa(123)
    d = d + "abc"
    fmt.Println(d)

    // 字符串转float64
    parse_float, _ := strconv.ParseFloat(orig, 64)
    fmt.Printf(`%t parse_float`, parse_float)
}

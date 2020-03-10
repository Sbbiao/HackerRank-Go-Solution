package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
    var res,counter int64 = 0,0
    // Edge case a only
    if s=="a"{
        return n
    }

    // Count 'a' in the string
    counter = int64(strings.Count(s,"a"))
    fmt.Println(counter)

    // Edge case no a
    if counter == 0{
        res = 0
        return res
    }else{
        if n>int64(len(s)){
            res = counter * (n/int64(len(s)))
            fmt.Println(res)
        }
    // Part 2
    if n%int64(len(s)) != 0{
        res += int64(strings.Count(s[:n % int64(len(s))], "a"))
    }
    }
    return res

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s := readLine(reader)

    n, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    result := repeatedString(s, n)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
    var res int32

    // var i,maximum int32
    // var counter [100] int32
    // Naive approach. Runtime error.
    // for i=0;i<n;i++{
    //     counter[ar[i]] += 1
    //     if ar[i]>maximum{
    //         maximum = ar[i]
    //     }
    // }
    // fmt.Println(counter)
    // for i=0;i<=maximum;i++{
    //     if counter[i] >=2{
    //         res += counter[i]/2
    //     }
    // }

    counter := make(map[string]int32)
    for _, sock := range ar{
        counter[strconv.FormatInt(int64(sock),10)] += 1
    }

    for _, v := range counter{
        res += v/2
    }
    return res

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arTemp := strings.Split(readLine(reader), " ")

    var ar []int32

    for i := 0; i < int(n); i++ {
        arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
        checkError(err)
        arItem := int32(arItemTemp)
        ar = append(ar, arItem)
    }

    result := sockMerchant(n, ar)

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

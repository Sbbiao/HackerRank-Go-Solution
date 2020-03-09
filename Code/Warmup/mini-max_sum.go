package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {

    var res int64 = 0
    var cmin int64 = 9999999999
    var cmax int64 = -9999999999
    for i:=0; i<len(arr);i++{
        res += int64(arr[i])
        if cmax < int64(arr[i]){
            cmax = int64(arr[i])
        }
        if cmin > int64(arr[i]){
            cmin = int64(arr[i])
        }
    }
    fmt.Printf("%d %d",res-cmax,res-cmin)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
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

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
    var length  = len(arr)
    var vp, vz, vm int = 0, 0, 0
    for i:=0; i<length; i++{
        if arr[i]<0{
            vm+=1
        }else if arr[i]==0{
            vz+=1
        }else{
            vp+=1
        }
    }

    var resm = float32(vm)/float32(length)
    var resz = float64(vz)/float64(length)
    var resp = float64(vp)/float64(length)
    fmt.Printf("%.6f\n",resp)
    fmt.Printf("%.6f\n",resm)
    fmt.Printf("%.6f\n",resz)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    plusMinus(arr)
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

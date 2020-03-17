package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the countSwaps function below.
func countSwaps(a []int32) {
    var res int32 = 0

    for i:=0;i<len(a);i++{

        for j:=0;j<len(a)-1;j++{

            if a[j] > a[j+1]{
                a[j] , a[j+1] = a[j+1], a[j]
                res++
            }
        }
    }
    fmt.Printf("Array is sorted in %d swaps.\n",res)
    fmt.Printf("First Element: %d\n",a[0])
    fmt.Printf("Last Element: %d\n",a[len(a)-1])

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    aTemp := strings.Split(readLine(reader), " ")

    var a []int32

    for i := 0; i < int(n); i++ {
        aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
        checkError(err)
        aItem := int32(aItemTemp)
        a = append(a, aItem)
    }

    countSwaps(a)
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

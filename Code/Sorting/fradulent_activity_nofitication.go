package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the activityNotifications function below.
func activityNotifications(expenditure []int32, d int32) int32 {
    var res int32
    var counter [201]int32

    // First count
    for i:=0; int32(i)<d; i++{
        counter[expenditure[i]]++
    }

    
    // Loop
    for i:=d; i<int32(len(expenditure)-1); i++{
        // Run the function to find 2* medium
        pmed := find_med(counter, d)

        // If larger than med,  res +=1 
        if pmed <= expenditure[i+1]{
            res +=1
        }

        // Remove previous counter
        // Add next int into counter
        counter[expenditure[i-d]]--
        counter[expenditure[i]]++
        fmt.Println(pmed/2)
    }
    
    return res
}

func find_med(c [201]int32, d int32 ) int32 {
    var res int32
    var count, a ,b int32

    // Even median
    if d%2 == 0{
        first:= d/2
        second:= first + 1
        for i:=0;i<201;i++{
            count += c[i]
            if first <= count{
                a = int32(i)
                break
            }
        }
        for i:=0;i<201;i++{
            count += c[i]
            if second <= count{
                b = int32(i)
                break
            }
        }
    // Odd median
    }else{
        first:=d/2 + 1
        for i:=0;i<201;i++{
            count += c[i]
            if first <= count{
                a = int32(i) * 2
                break
            }
        }
    }

    res = a+b

    // res is 2*Median
    return res
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nd := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nd[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    dTemp, err := strconv.ParseInt(nd[1], 10, 64)
    checkError(err)
    d := int32(dTemp)

    expenditureTemp := strings.Split(readLine(reader), " ")

    var expenditure []int32

    for i := 0; i < int(n); i++ {
        expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
        checkError(err)
        expenditureItem := int32(expenditureItemTemp)
        expenditure = append(expenditure, expenditureItem)
    }

    result := activityNotifications(expenditure, d)

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

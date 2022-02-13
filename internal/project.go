package main
import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func write()  {
	file, err := os.Create("./temp.txt")
    if err != nil {
        log.Fatal(err)
    }
    writer := bufio.NewWriter(file)
    linesToWrite := []string{"This is an example", "to show how", "to write to a file", "line by line."}
    for _, line := range linesToWrite {
        bytesWritten, err := writer.WriteString(line + "\n")
        if err != nil {
            log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
        }
        fmt.Printf("Bytes Written: %d\n", bytesWritten)
        fmt.Printf("Available: %d\n", writer.Available())
        fmt.Printf("Buffered : %d\n", writer.Buffered())
    }
    writer.Flush()	
}
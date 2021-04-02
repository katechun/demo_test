package main

import (
	"fmt"
	"io"
)

type SampleStore struct {
        data []byte
}

func (ss *SampleStore) Write(p []byte)(n int,err error){
        if (len(ss.data) == 10 ) {
                return 0,io.EOF
        }

        remainingCap := 10 - len(ss.data)

        writeLength := len(p)
        if remainingCap <= writeLength {
                writeLength = remainingCap;
                err = io.EOF
        }

        ss.data = append(ss.data,p[:writeLength]...)
        n = writeLength
        return
}



func main(){
        ss := &SampleStore{}
        bytesWritten1,err1 := io.WriteString(ss,"Hello!")
        fmt.Printf("Bytes written %d,error: %v\n", bytesWritten1,err1)
        fmt.Printf("Value of ss.data: %s\n",ss.data)

        bytesWritten2,err2 := io.WriteString(ss," Amazing")
         fmt.Printf("Bytes written %d,error: %v\n", bytesWritten2,err2)
        fmt.Printf("Value of ss.data: %s\n",ss.data)

        bytesWritten3,err3 := io.WriteString(ss," Amazing")
         fmt.Printf("Bytes written %d,error: %v\n", bytesWritten3,err3)
        fmt.Printf("Value of ss.data: %s\n",ss.data)

}


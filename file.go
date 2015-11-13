package file 

import (
	"os"
	"bufio"
)

func FileLines(fileToken string) a []string {
    file, err := os.Open(fileToken)
    if err != nil { panic(err) }
    // close file on exit and check for its returned error
    defer func() {
        if err := file.Close(); err != nil {
            panic(err)
        }
    }()
    r := bufio.NewReader(file)

    line, err := r.ReadString('\n')

    for err == nil {
    	a = append(a, line)
    	line, err = r.ReadString('\n')
    }
    return
}
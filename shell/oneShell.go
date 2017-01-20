package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "os/exec"
    "github.com/kr/pty"
)

func copyLinesUntil(r io.Reader, w io.Writer, match string) error {
    fmt.Println("inside copyLinesUntil")
    scanner := bufio.NewScanner(r)
    for scanner.Scan() {
        t := scanner.Text()
        if t == match {
                return nil
        }
        fmt.Fprintln(w, t)
    }
    return nil
}

func main() {
    c := exec.Command("sh", "-i")
    c.Stderr = os.Stderr
    f, err := pty.Start(c)
    if err != nil {
        panic(err)
    }

    f.Write([]byte("export PS1='----END OF INPUT---\\n> '\n")) // EOT

    go func() {
        // scanner := bufio.NewScanner(os.Stdin)
        // for scanner.Scan() {
        //     t := scanner.Text()
        //     fmt.Fprintf(f, "%s\n", t)
        // }
        //
        // fmt.Println("Exiting...")

        f.Write([]byte(". temp1.sh; echo uniqueid $? \n"))
        f.Write([]byte(". temp2.sh; echo xxxx $? \n"))
        f.Write([]byte{4}) // EOT
    }()
    copyLinesUntil(f, os.Stdout, "uniqueid 0")
    copyLinesUntil(f, os.Stdout, "xxxx 0")
    // io.Copy(os.Stdout, f)
}
package main

// import (
// 	"bytes"
// 	"fmt"
// 	"log"
// 	"os/exec"
// )

// func ShellOut(command string) (error, string, string) {
// 	var stdout bytes.Buffer
// 	var stderr bytes.Buffer
// 	cmd := exec.Command("bash", "-c", command)
// 	cmd.Stdout = &stdout
// 	cmd.Stderr = &stderr
// 	err := cmd.Run()
// 	return err, stdout.String(), stderr.String()
// }
// func main() {
// 	err, out, errout := ShellOut("ls -ltr")
// 	if err != nil {
// 		log.Printf("error:%v\n", err)
// 	}
// 	fmt.Println("--- stdout ---")
// 	fmt.Println(out)
// 	fmt.Println("--- stderr ---")
// 	fmt.Println(errout)

// }

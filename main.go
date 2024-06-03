package main

import (
	"bufio"
	"fmt"
	"lspserver_go/rpc"
	"os"
)

func main() {
  fmt.Println("test")

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(rpc.Split)

  for scanner.Scan() {
    msg := scanner.Text()
    HandleMessage(msg)  
  }
}

func HandleMessage (msg any) {

}

package main

import (
	"bufio"
	"log"
	"lspserver_go/rpc"
	"os"
)

func main() {
  logger := getLogger("/home/so7en/Desktop/lspserver-go/log.txt")
  logger.Println("logger started...")

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(rpc.Split)

  for scanner.Scan() {
    msg := scanner.Text()
    HandleMessage(msg)  
  }
}

func getLogger(filename string) *log.Logger {
  logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666) // 0666 -> creatiing file
  if err != nil {
    panic("wrong file")
  }
  return log.New(logfile, "[lspserver_go]", log.Ldate|log.Ltime|log.Lshortfile)
}

func HandleMessage (msg any) {

}

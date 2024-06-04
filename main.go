package main

import (
	"bufio"
	"encoding/json"
	"log"
	"lspserver_go/lsp"
	"lspserver_go/rpc"
	"os"
	"path"
)

func main() {
  root_folder, _ := os.Getwd()
  log_path := path.Join((root_folder), "log.txt")
  logger := getLogger(log_path)
  logger.Println("logger started...")

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(rpc.Split)

  for scanner.Scan() {
    msg := scanner.Bytes()
    method, content, err := rpc.DecodeMessage(msg)
    if err != nil {
      logger.Printf("Error: %s", err)
      continue
    }
    HandleMessage(logger, method, content)  
  }
}

func getLogger(filename string) *log.Logger {
  logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666) // 0666 -> creatiing file
  if err != nil {
    panic("wrong file")
  }
  return log.New(logfile, "[lspserver_go]", log.Ldate|log.Ltime|log.Lshortfile)
}

func HandleMessage (logger *log.Logger, method string, contents []byte) {
  logger.Printf("received msg with method: %s", method)
  switch method {
  case "initialize":
    var request lsp.InitializeRequest
    if err := json.Unmarshal(contents, &request); err != nil {
      logger.Printf("could not parse: %s", err)
    }
    logger.Printf("connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo)

    msg := lsp.NewInitializeResponse(request.ID)
    reply := rpc.EncodeMessage(msg)
    writer := os.Stdout
    writer.Write([]byte(reply))
    logger.Print("reply sent")

  case "textDocument/didOpen":
    var request lsp.DidOpenTextDocumentNotif
    if err := json.Unmarshal(contents, &request); err != nil {
      logger.Printf("could not parse: %s", err)
    }
    logger.Printf("opened to: %s %s", request.Params.TextDocument.URI, request.Params.TextDocument.Text)
  }
}

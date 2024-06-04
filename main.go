package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"lspserver_go/compiler"
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

  analyzer := compiler.NewAnalyze() 
  writer := os.Stdout

  for scanner.Scan() {
    msg := scanner.Bytes()
    method, content, err := rpc.DecodeMessage(msg)
    if err != nil {
      logger.Printf("Error: %s", err)
      continue
    }
    HandleMessage(logger, writer, analyzer,method, content)  
  }
}

func getLogger(filename string) *log.Logger {
  logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666) // 0666 -> creatiing file
  if err != nil {
    panic("wrong file")
  }
  return log.New(logfile, "[lspserver_go]", log.Ldate|log.Ltime|log.Lshortfile)
}

func WriteResponse(writer io.Writer, msg any) {
    reply := rpc.EncodeMessage(msg)
    writer.Write([]byte(reply))
}

func HandleMessage (logger *log.Logger, writer io.Writer, analyzer compiler.Analyzer, method string, contents []byte) {
  logger.Printf("received msg with method: %s", method)
  switch method {
  case"initialize":
    var request lsp.InitializeRequest
    if err := json.Unmarshal(contents, &request); err != nil {
      logger.Printf("could not parse: %s", err)
    }
    logger.Printf("connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo)

    msg := lsp.NewInitializeResponse(request.ID)
    WriteResponse(writer, msg)
    logger.Print("reply sent")

  case "textDocument/didOpen":
    var request lsp.DidOpenTextDocumentNotif
    if err := json.Unmarshal(contents, &request); err != nil {
      logger.Printf("textDoc/didOpen: %s", err)
    }
    logger.Printf("opened to: %s %s", request.Params.TextDocument.URI, request.Params.TextDocument.Text)
    analyzer.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
  case "textDocument/didChange":
    var request lsp.DidChangeTextDoucumentNotif
    if err := json.Unmarshal(contents, &request); err != nil {
      logger.Printf("textDoc/didChange: %s", err)
    }
    logger.Printf("changed : %s %s", request.Params.TextDocument.URI, request.Params.ContentChanges)
    for _, change := range request.Params.ContentChanges {
      analyzer.UpdateDocument(request.Params.TextDocument.URI, change.Text)
    }
  case "textDocument/hover":
    var request lsp.HoverRequest
    if err := json.Unmarshal(contents, &request); err != nil {
      logger.Printf("textDoc/hover: %s", err)
    }
    response := analyzer.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)
    WriteResponse(writer, response)
  case "textDocument/definition":
    var request lsp.DefinitionRequest
    if err := json.Unmarshal(contents, &request); err != nil {
      logger.Printf("textDoc/definition: %s", err)
    }
    response := analyzer.Definition(request.ID, request.Params.TextDocument.URI, request.Params.Position)
    WriteResponse(writer, response)
  case "textDocument/codeAction":
    var request lsp.CodeActionRequest
    if err := json.Unmarshal(contents, &request); err != nil {
      logger.Printf("textDoc/codeAction: %s", err)
    }
    response := analyzer.TextDocumentCodeAction(request.ID, request.Params.TextDocument.URI)
    WriteResponse(writer, response)

  }
}

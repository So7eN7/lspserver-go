package lsp

type TextDocumentItem struct {
  URI string `json:"uri"` // DocumentUri
  LanguageId string `json:"languageId"`
  Version int `json:"version"`
  Text string `json:"text"`
}

type DidOpenTextDocumentNotif struct {
  Notification 
  Params DidOpenTextDocumentParams `json:"params"`
}

type DidOpenTextDocumentParams struct {
  TextDocument TextDocumentItem `json:"textDocument"` 
}

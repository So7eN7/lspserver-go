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

type DidChangeTextDoucumentNotif struct {
  Notification
  Params DidChangeTextDoucumentParams `json:"params"`
}

type DidChangeTextDoucumentParams struct {
  TextDocument VersionTextDocumentIdentifer `json:"textdocument"` 
  ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type TextDocumentIdentifier struct {
  URI string `json:"uri"`
}

type VersionTextDocumentIdentifer struct {
  TextDocumentIdentifier
  Version int `json:"version"`
}

type TextDocumentContentChangeEvent struct {
  Text string `json:"text"`
}

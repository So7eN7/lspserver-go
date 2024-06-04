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

type HoverRequest struct {
  Request
  Params HoverParams `json:"params"`
}

type HoverParams struct {
  TextDocumentPositionParams
}

type HoverResponse struct {
  Response
  Result HoverResult `json:"result"`
}

type HoverResult struct {
  Contents string `json:"contents"`
}

type TextDocumentPositionParams struct {
  TextDocument TextDocumentIdentifier `json:"textDocument"`
  Position Position `json:"position"`
}

type Position struct {
  Line int `json:"line"`
  Char int `json:"char"`
}

type DefinitionRequest struct {
  Request
  Params HoverParams `json:"params"`
}

type DefinitionParams struct {
  TextDocumentPositionParams
}

type DefinitionResponse struct {
  Response
  Result Location `json:"result"`
}

type Location struct {
  URI string `json:"uri"`
  Range Range `json:"range"`
}

type Range struct {
  Start Position `json:"start"`
  End Position `json:"end"`
}

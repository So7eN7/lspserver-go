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

type CodeActionRequest struct {
  Request
  Params TextDocumentCodeActionParams `json:"params"`
}

type TextDocumentCodeActionParams struct {
  TextDocument TextDocumentIdentifier `json:"textDocument"`
  Range Range `json:"range"`
  Context CodeActionContext `json:"context"`
}

type TextDocumentCodeActionResponse struct {
  Response
  Result []CodeAction `json:"result"`
}

type CodeActionContext struct {

}

type CodeAction struct {
  Title string `json:"title"`
  Edit *WorkspaceEdit `json:"edit,omitempty"`
  Command *Command `json:"command,omitempty"`
}

type WorkspaceEdit struct {
  Changes map[string][]TextEdit `json:"changes"`
}

type TextEdit struct {
  Range Range `json:"range"`
  NewText string `json:"newText"`
}

type Command struct {
  Title string `json:"title"`
  Command string `json:"command"`
  Arguments []interface{} `json:"arguments,omitempty"`
}


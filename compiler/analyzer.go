package compiler

import (
	"fmt"
	"lspserver_go/lsp"
	"strings"
)

type Analyzer struct {
  Documents map[string]string
}

func NewAnalyze() Analyzer {
  return Analyzer{Documents: map[string]string{}}
}

func (a *Analyzer) OpenDocument(uri, text string) {
  a.Documents[uri] = text
}

func (a *Analyzer) UpdateDocument(uri, text string) {
  a.Documents[uri] = text
}

func (a *Analyzer) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
  document := a.Documents[uri]

  return lsp.HoverResponse{
      Response: lsp.Response{
        ID: &id,
        RPC: "2.0",        
      },
      Result: lsp.HoverResult{
        Contents: fmt.Sprintf("file: %s, chars: %d", uri, len(document)),
      },
    }
}

func (a *Analyzer) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
  
  return lsp.DefinitionResponse{
      Response: lsp.Response{
        ID: &id,
        RPC: "2.0",        
      },
      Result: lsp.Location{
        URI: uri,
        Range: lsp.Range{
          Start: lsp.Position{
            Line: position.Line - 1, // placeholder 
            Char: 0,
          },
          End: lsp.Position{
            Line: position.Line - 1,
            Char: 0, 
          },
        },
      },
    }
}

func (a *Analyzer) TextDocumentCodeAction(id int, uri string) lsp.TextDocumentCodeActionResponse {
  text := a.Documents[uri]
  actions := []lsp.CodeAction{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx >= 0 {
			replaceChange := map[string][]lsp.TextEdit{}
			replaceChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "Neovim",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Replace VS C*de with a superior editor",
				Edit:  &lsp.WorkspaceEdit{Changes: replaceChange},
			})

			censorChange := map[string][]lsp.TextEdit{}
			censorChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "VS C*de",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Censor to VS C*de",
				Edit:  &lsp.WorkspaceEdit{Changes: censorChange},
			})
		}
	}

	response := lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: actions,
	}

	return response
}


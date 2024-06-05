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

func GetDiagnosticsForFile(text string) []lsp.Diagnostics {
  diagnostics := []lsp.Diagnostics{}
  for row, line := range strings.Split(text, "\n"){
    if strings.Contains(line, "example") {
      idx := strings.Index(line, "example")
      diagnostics = append(diagnostics, lsp.Diagnostics{
        Range: LineRange(row, idx, idx+len("example")),
        Severity: 1,
        Source: "Halo",
        Message: "The Great Journey is nigh",
      })
    }
    if strings.Contains(line, "example") {
      idx := strings.Index(line, "example")
      diagnostics = append(diagnostics, lsp.Diagnostics{
        Range: LineRange(row, idx, idx+len("example")),
        Severity: 1,
        Source: "Halo",
        Message: "The Great Journey is nigh",
      })
    }
  }
  return diagnostics
}

func (a *Analyzer) OpenDocument(uri, text string) []lsp.Diagnostics {
  a.Documents[uri] = text
  return GetDiagnosticsForFile(text)
}

func (a *Analyzer) UpdateDocument(uri, text string) []lsp.Diagnostics {
  a.Documents[uri] = text
  return GetDiagnosticsForFile(text)
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
		idx := strings.Index(line, "example")
		if idx >= 0 {
			replaceChange := map[string][]lsp.TextEdit{}
			replaceChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("example")),
					NewText: "exampleNew",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "replace example with exampleNew",
				Edit:  &lsp.WorkspaceEdit{Changes: replaceChange},
			})

			censorChange := map[string][]lsp.TextEdit{}
			censorChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("example")),
					NewText: "e*ample",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "censor e*ample",
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

func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Char: start,
		},
		End: lsp.Position{
			Line:      line,
			Char: end,
		},
	}
}

func (a *Analyzer) TextDocumentCompletion(id int, uri string) lsp.CompletionResponse {
	items := []lsp.CompletionItem{
    {
      Label: "completionEx",
      Detail: "Detail for completionEx",
      Documentation: "Halo's light will shine through",
    },
  }
	
	response := lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}
	return response
}



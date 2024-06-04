package compiler

import (
	"fmt"
	"lspserver_go/lsp"
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

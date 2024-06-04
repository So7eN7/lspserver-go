package compiler

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

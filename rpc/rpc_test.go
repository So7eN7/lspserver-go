package rpc_test

import (
	"lspserver_go/rpc"
	"testing"
)

type EncodingExample struct {
  Testing bool
}

func TestEncode(t *testing.T) {
  expceted := "content_len: 16\r\n\r\n{\"Testing\":true}"
  actual := rpc.EncodeMessage(EncodingExample{Testing:true})
  if expceted != actual {
    t.Fatalf("expected: %s, actual: %s", expceted, actual)
  }
}

func TestDecode(t *testing.T) {
  incomingMessage := "content_len: 15\r\n\r\n{\"Method\":\"hi\"}" // "hi" is 4 in length  
  method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
  content_len := len(content)
  if err != nil {
    t.Fatal(err)
  }

  if content_len != 15 {
    t.Fatalf("expceted: 16, got: %d", content_len)
  }

  if method != "hi" {
    t.Fatalf("expceted hi, got: %s", method)
  }
}

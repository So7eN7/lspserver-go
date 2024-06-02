package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
  "errors"
)

type BaseMessage struct {
  Method string `json:"method"`
}

func EncodeMessage(msg any)  string {
  content, err := json.Marshal(msg)
  if err != nil {
    panic(err)
  }
  return fmt.Sprintf("content_len: %d\r\n\r\n%s", len(content), content)
}

func DecodeMessage(msg []byte) (string, []byte, error) {
  header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
  if !found {
    return "", nil, errors.New("did not find seperator")
  }

  contentLenBytes := header[len("content_len: "):] // after content_len
  contentlen, err := strconv.Atoi(string(contentLenBytes))
  if err != nil {
    return "", nil, err 
  }

  var baseMessage BaseMessage

  if err := json.Unmarshal(content[:contentlen], &baseMessage); err != nil {
    return "", nil, err
  }

  return baseMessage.Method, content[:contentlen] ,nil 
}

package message

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
  message := GetMessage()
  if message == ""{
    t.Error("Message returned nothing")
  }
}

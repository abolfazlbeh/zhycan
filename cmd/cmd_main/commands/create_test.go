package commands

import (
	"bytes"
	"io"
	"testing"
)

func Test_ExecuteCreateCmd(t *testing.T) {
	createCmd := NewCreateCmd()

	b := bytes.NewBufferString("")
	createCmd.SetOut(b)

	createCmd.Execute()
	_, err := io.ReadAll(b)
	if err != nil {
		t.Errorf("Expected to read without err, bot got: %v", err)
		return
	}
}

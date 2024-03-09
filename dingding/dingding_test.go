package dingding

import "testing"

func TestSendMarkdown(t *testing.T) {
	err := Default().SendMarkdownMessage("Test Title", "Test Text", nil, true)
	if err != nil {
		t.Fatal(err)
	}
}

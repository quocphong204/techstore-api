package utils

import (
	"testing"
)

func TestSendOrderEmail(t *testing.T) {
	err := SendOrderEmail("phong150718@gmail.com", "Test Email", "<b>This is a test</b>")
	if err != nil {
		t.Errorf("❌ SendOrderEmail failed: %v", err)
	} else {
		t.Log("✅ SendOrderEmail succeeded")
	}
}

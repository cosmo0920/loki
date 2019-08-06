package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLokiConfig(t *testing.T) {
	c, err := getLokiConfig("", "", "", "")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	assert.Equal(t, "http://localhost:3100/api/prom/push", c.url.String(), "Use default value of URL")
	assert.Equal(t, 1*time.Second, c.batchWait, "Use default value of batchWait")
	assert.Equal(t, 100*1024, c.batchSize, "Use default value of batchSize")

	// Invalid URL
	_, err = getLokiConfig("invalid---URL+*#Q(%#Q", "", "", "")
	if err == nil {
		t.Fatalf("failed test %#v", err)
	}

	// batchWait, batchSize

	c, err = getLokiConfig("", "15", "30720", "")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	assert.Equal(t, 15*time.Second, c.batchWait, "Use user-defined value of batchWait")
	assert.Equal(t, 30*1024, c.batchSize, "Use user-defined value of batchSize")

	// LabelSets
	labels := `{test="fluent-bit-go", lang="Golang"}`
	c, err = getLokiConfig("", "15", "30", labels)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	assert.Equal(t, "{lang=\"Golang\", test=\"fluent-bit-go\"}",
		c.labelSet.String(), "Use user-defined value of labels")

}

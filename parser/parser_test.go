package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessParser(t *testing.T) {
	TestData := []byte(
		`- name: process1
  url: url1
  selector: selector1
  extractor: extractor1
  actions:
    - name: print
`)
	p, err := NewProcessParser(TestData)
	assert.Nil(t, err)
	assert.Equal(t, "process1", p[0].Name)
	assert.Equal(t, "url1", p[0].Url)
	assert.Equal(t, "selector1", p[0].Selector)
	assert.Equal(t, "extractor1", p[0].Extractor)
	assert.Equal(t, "print", p[0].Actions[0].Name)
}

package parser

import (
	"testing"
)

var testData = []byte(`- name: job1
  actions:
    - name: get
      options:
        url: https://ip.xmcd.org/api
        useragent: myuseragent1
    - name: log
- name: job2
  actions:
    - name: get
      options:
        url: https://ip.xmcd.org/api
        useragent: myuseragent2
    - name: log`)

func TestParser(t *testing.T) {
	var actions []Actions
	t.Run("test parse", func(t *testing.T) {
		var err error
		if actions, err = Parse(testData); err != nil {
			t.Error(t.Name(), err)
		}
	})
	t.Run("test number of jobs", func(t *testing.T) {
		if len(actions) != 2 {
			t.Error("Error: number of jobs", len(actions))
		}
	})
	t.Run("test number of job actions", func(t *testing.T) {
		if len(actions[0].Actions) != 2 {
			t.Error("Error: number of job actions", len(actions))
		}
	})
}

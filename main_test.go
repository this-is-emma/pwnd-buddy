package main

import (
	"bytes"
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestMain_NoEmailFlag(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "No email flag provided",
			args:     []string{"cmd"},
			expected: "Error: Please provide an email address using the -email flag\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			os.Args = test.args

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			main()

			w.Close()
			os.Stdout = oldStdout

			var buf bytes.Buffer
			buf.ReadFrom(r)

			out := buf.String()
			if out != test.expected {
				t.Errorf("Expected: %s, Got: %s", test.expected, out)
			}
		})
	}
}

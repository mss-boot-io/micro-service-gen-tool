package pkg

import "testing"

func TestUpdate(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Install()
		})
	}
}

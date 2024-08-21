package console

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPortShower(t *testing.T) {
	tests := []struct {
		name string
		want PortShower
	}{
		{"new port shower", PortShower{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPortShower(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPortShower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPortShower_Show(t *testing.T) {
	type args struct {
		ports []string
	}
	tests := []struct {
		name   string
		p      PortShower
		args   args
		expect string
	}{
		{
			name: "show ports",
			args: args{
				ports: []string{"/dev/ttyUSB0", "/dev/ttyUSB1"},
			},
			expect: "Ports:\n\t-\t/dev/ttyUSB0\n\t-\t/dev/ttyUSB1\n",
		},
		{
			name: "show no ports",
			args: args{
				ports: []string{},
			},
			expect: "no ports found\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			shower := NewPortShower()

			// prepare to capture stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Act
			shower.Show(tt.args.ports)

			// restore stdout
			w.Close()
			os.Stdout = old

			// read the output of the command
			var buf bytes.Buffer
			_, _ = io.Copy(&buf, r)
			captured := buf.String()

			// Assert
			assert.Equal(t, tt.expect, captured, "Output does not match expected")
		})
	}
}

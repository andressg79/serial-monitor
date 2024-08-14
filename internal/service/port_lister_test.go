package service

import (
	"reflect"
	"testing"
)

func TestNewPortLister(t *testing.T) {
	type args struct {
		serial SerialLister
		shower PortShower
	}
	tests := []struct {
		name string
		args args
		want PortLister
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPortLister(tt.args.serial, tt.args.shower); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPortLister() = %v, want %v", got, tt.want)
			}
		})
	}
}

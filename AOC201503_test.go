package main

import (
	"reflect"
	"testing"
)

func Test_getNewHouse(t *testing.T) {
	type args struct {
		h House
		m string
	}
	tests := []struct {
		name string
		args args
		want House
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNewHouse(tt.args.h, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNewHouse() = %v, want %v", got, tt.want)
			}
		})
	}
}

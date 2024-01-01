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
		//
		{"Go up", args{House{0, 0}, "^"}, House{0, 1}},
		{"Go down", args{House{0, 0}, "v"}, House{0, -1}},
		{"Go left", args{House{0, 0}, "<"}, House{-1, 0}},
		{"Go right", args{House{0, 0}, ">"}, House{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNewHouse(tt.args.h, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNewHouse() = %v, want %v", got, tt.want)
			}
		})
	}
}

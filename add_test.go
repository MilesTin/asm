package main

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		x int64
		y int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				x:int64(1),
				y:int64(2),
			},
			want: int64(3),
		},
		{
			args: args{
				x:int64(10),
				y:int64(-10),
			},
			want: int64(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
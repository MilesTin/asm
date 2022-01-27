package main

import (
	"testing"

	"github.com/bytedance/gopkg/lang/fastrand"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSwapInt64(t *testing.T) {
	type args struct {
		addr *int64
		new  int64
	}

	old1 := int64(1)
	old2 := int64(2)

	tests := []struct {
		name    string
		args    args
		wantOld int64
	}{
		{
			args: args{
				addr: new(int64),
				new:  int64(1),
			},
			wantOld: int64(0),
		},
		{
			args: args{
				addr: &old1,
				new:  int64(2),
			},
			wantOld: 1,
		},
		{
			args: args{
				addr: &old2,
				new:  int64(100),
			},
			wantOld: int64(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOld := SwapInt64(tt.args.addr, tt.args.new); gotOld != tt.wantOld {
				t.Errorf("SwapInt64() = %v, want %v", gotOld, tt.wantOld)
			}
			if *tt.args.addr != tt.args.new {
				t.Errorf("SwapInt64() new = %v, want %v", *tt.args.addr, tt.args.new)
			}
		})
	}
}

func FuzzSwapInt64(f *testing.F) {
	f.Fuzz(func(t *testing.T, old int64, new int64) {
		orig := old
		oldV := SwapInt64(&old, new)
		Convey("SwapInt64", t, func() {
			So(oldV, ShouldEqual, orig)
			So(old, ShouldEqual, new)
		})
	})
}

func BenchmarkSwapInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		old := fastrand.Int63()
		new := fastrand.Int63()
		SwapInt64(&old, new)
	}
}

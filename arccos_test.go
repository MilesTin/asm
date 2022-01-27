package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/bytedance/gopkg/lang/fastrand"
	. "github.com/smartystreets/goconvey/convey"
)

func FuzzArccos(f *testing.F) {
	f.Fuzz(func(t *testing.T, x float64) {
		if x < -1.0 || x > 1.0 {
			return
		}
		Convey(fmt.Sprintf("arccos(%v)", x), t, func() {
			So(Arccos(x), ShouldAlmostEqual, math.Acos(x))
		})
	})
}

func TestArccos(t *testing.T) {
	Convey("arccos", t, func() {
		Convey("1.0", func() {
			So(Arccos(1.0), ShouldAlmostEqual, math.Acos(1.0))
		})
		Convey("-1.0", func() {
			So(Arccos(-1.0), ShouldAlmostEqual, math.Acos(-1.0))
		})
		Convey("0.5", func() {
			So(Arccos(0.5), ShouldAlmostEqual, math.Acos(0.5))
		})
	})
}

func BenchmarkArccos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Arccos(fastrand.Float64())
	}
}

func BenchmarkArcsin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Arcsin(fastrand.Float64())
	}
}

func BenchmarkMathArccos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Acos(fastrand.Float64())
	}
}

func BenchmarkMathAsin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Asin(fastrand.Float64())
	}
}

package main

import (
	"fmt"
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestArcsin(t *testing.T) {
	Convey("arcsin", t, func() {
		Convey("1.0", func() {
			So(Arcsin(1.0), ShouldAlmostEqual, math.Asin(1.0))
		})
		Convey("0.5", func() {
			So(Arcsin(0.5), ShouldAlmostEqual, math.Asin(0.5))
		})
		Convey("-1.0", func() {
			So(Arcsin(-1.0), ShouldAlmostEqual, math.Asin(-1.0))
		})
		Convey("-0.5142857142857142", func() {
			So(Arcsin(-0.5142857142857142), ShouldAlmostEqual, math.Asin(-0.5142857142857142))
		})
	})
}

func FuzzArcsin(f *testing.F) {
	f.Fuzz(func(t *testing.T, x float64) {
		if x < -1.0 || x > 1.0 {
			return
		}
		Convey(fmt.Sprintf("%v", x), t, func() {
			So(Arcsin(x), ShouldAlmostEqual, math.Asin(x))
		})
	})
}

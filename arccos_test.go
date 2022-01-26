package main

import (
	"fmt"
	"testing"
	"math"
	. "github.com/smartystreets/goconvey/convey"
)

func FuzzArccos(f *testing.F) {

	f.Fuzz(func(t *testing.T, x float64) {
		if x < -1.0 || x > 1.0 {
			return
		}
		Convey(fmt.Sprintf("arccos(%v)", x), t, func() {
			So(Arccos(x),ShouldEqual, math.Acos(x))
		})
	})
}
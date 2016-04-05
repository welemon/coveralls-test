package tests

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	Convey("Subject: Add", t, func() {
		Convey("Add 1 + 1 should be 2", func() {
			So(Add(1, 1), ShouldEqual, 2)
		})
	})
}

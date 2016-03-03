package cli

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCli(t *testing.T) {
	Convey("validateLongOption", t, func() {
		So(validateLongOption("ok"), ShouldBeNil)
		So(validateLongOption(""), ShouldNotBeNil)
		So(validateLongOption("1test"), ShouldNotBeNil)
		So(validateLongOption("test1"), ShouldBeNil)
		So(validateLongOption("test-it"), ShouldBeNil)
		So(validateLongOption("Test-It"), ShouldBeNil)
	})
}

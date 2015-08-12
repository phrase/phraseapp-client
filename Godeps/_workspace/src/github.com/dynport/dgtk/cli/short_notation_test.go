package cli

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type ActionWithShortNotation struct {
	Flag      bool     `cli:"opt -f --flag"`
	Option1   string   `cli:"opt -o --option required"`
	Option2   string   `cli:"opt -p --param default='something'"`
	Argument1 string   `cli:"arg required"`
	Argument2 []string `cli:"arg"`
}

func (a *ActionWithShortNotation) Run() error {
	return nil
}

func TestActionWithShortNotation(t *testing.T) {
	var act *ActionWithShortNotation
	var a *action
	var e error
	Convey("Given an action with short notation", t, func() {
		act = &ActionWithShortNotation{}
		Convey("When the reflect method is called on it", func() {
			a = testCreateAction("test", act)
			e = a.reflect()
			Convey("Then there is no error", func() {
				So(e, ShouldBeNil)
			})
			Convey("When empty parameters are parsed", func() {
				e = a.parseArgs([]string{})
				Convey("Then an error is returned", func() {
					So(e, ShouldNotBeNil)
					So(e.Error(), ShouldEqual, `option "Option1" is required but not set`)
				})
			})
			Convey("When the required parameters are given", func() {
				e = a.parseArgs([]string{"-o", "foo", "bar"})
				Convey("Then there is no error", func() {
					So(e, ShouldBeNil)
				})
				Convey("The option and argument are set accordingly", func() {
					So(act.Flag, ShouldBeFalse)
					So(act.Option1, ShouldEqual, "foo")
					So(act.Option2, ShouldEqual, "something")
					So(act.Argument1, ShouldEqual, "bar")
					So(len(act.Argument2), ShouldEqual, 0)
				})
			})
			Convey("When the flag is given", func() {
				e = a.parseArgs([]string{"-f", "-o", "foo", "bar"})
				Convey("Then there is no error", func() {
					So(e, ShouldBeNil)
				})
				Convey("The option and argument are set accordingly", func() {
					So(act.Flag, ShouldBeTrue)
					So(act.Option1, ShouldEqual, "foo")
					So(act.Option2, ShouldEqual, "something")
					So(act.Argument1, ShouldEqual, "bar")
					So(len(act.Argument2), ShouldEqual, 0)
				})
			})
			Convey("When the second option is given", func() {
				e = a.parseArgs([]string{"--param", "fuu", "-o", "foo", "bar"})
				Convey("Then there is no error", func() {
					So(e, ShouldBeNil)
				})
				Convey("The option and argument are set accordingly", func() {
					So(act.Flag, ShouldBeFalse)
					So(act.Option1, ShouldEqual, "foo")
					So(act.Option2, ShouldEqual, "fuu")
					So(act.Argument1, ShouldEqual, "bar")
					So(len(act.Argument2), ShouldEqual, 0)
				})
			})
		})
	})
}

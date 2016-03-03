package cli

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type ActionWithPointerTypes struct {
	Flag    *bool   `cli:"opt -f"`
	Option1 *string `cli:"opt -o"`
	Option2 *int    `cli:"opt -p"`
}

func (a *ActionWithPointerTypes) Run() error {
	return nil
}

func TestActionWithPointerTypes(t *testing.T) {
	var act *ActionWithPointerTypes
	var a *action
	var e error
	Convey("Given an action with short notation", t, func() {
		act = &ActionWithPointerTypes{}
		Convey("When the reflect method is called on it", func() {
			a = testCreateAction("test", act)
			e = a.reflect()
			Convey("Then there is no error", func() {
				So(e, ShouldBeNil)
			})
			Convey("When no parameters are given", func() {
				e = a.parseArgs([]string{})
				Convey("Then there is no error", func() {
					So(e, ShouldBeNil)
				})
				Convey("The option and argument are set accordingly", func() {
					So(act.Flag, ShouldEqual, nil)
					So(act.Option1, ShouldEqual, nil)
					So(act.Option2, ShouldEqual, nil)
				})
			})
			Convey("When the flag is given to false", func() {
				e = a.parseArgs([]string{"-f", "false"})
				Convey("Then there is no error", func() {
					So(e, ShouldBeNil)
				})
				Convey("The option and argument are set accordingly", func() {
					So(*act.Flag, ShouldBeFalse)
					So(act.Option1, ShouldEqual, nil)
					So(act.Option2, ShouldEqual, nil)
				})
			})
			Convey("When the flag is given to true", func() {
				e = a.parseArgs([]string{"-f", "true"})
				Convey("Then there is no error", func() {
					So(e, ShouldBeNil)
				})
				Convey("The option and argument are set accordingly", func() {
					So(*act.Flag, ShouldBeTrue)
					So(act.Option1, ShouldEqual, nil)
					So(act.Option2, ShouldEqual, nil)
				})
			})
			Convey("When the string parameters is given", func() {
				e = a.parseArgs([]string{"-o", "foo"})
				Convey("Then there is no error", func() {
					So(e, ShouldBeNil)
				})
				Convey("The option and argument are set accordingly", func() {
					So(act.Flag, ShouldEqual, nil)
					So(*act.Option1, ShouldEqual, "foo")
					So(act.Option2, ShouldEqual, nil)
				})
			})
			Convey("When the second option is given", func() {
				e = a.parseArgs([]string{"-p", "1"})
				Convey("Then there is no error", func() {
					So(e, ShouldBeNil)
				})
				Convey("The option and argument are set accordingly", func() {
					So(act.Flag, ShouldEqual, nil)
					So(act.Option1, ShouldEqual, nil)
					So(*act.Option2, ShouldEqual, 1)
				})
			})
			Convey("When all options are given", func() {
				e = a.parseArgs([]string{"-f", "true", "-o", "foo", "-p", "1"})
				Convey("Then there is no error", func() {
					So(e, ShouldBeNil)
				})
				Convey("The option and argument are set accordingly", func() {
					So(*act.Flag, ShouldBeTrue)
					So(*act.Option1, ShouldEqual, "foo")
					So(*act.Option2, ShouldEqual, 1)
				})
			})
		})
	})
}

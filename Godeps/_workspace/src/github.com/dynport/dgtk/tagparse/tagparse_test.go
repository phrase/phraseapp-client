package tagparse

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTagSplit(t *testing.T) {
	Convey("Given an empty tag string", t, func() {
		tagString := ""
		Convey("When the tag splitter is called", func() {
			fields, e := tagSplit(tagString)
			Convey("Then there is no error", func() {
				So(e, ShouldBeNil)
			})
			Convey("Then the field list is empty", func() {
				So(len(fields), ShouldEqual, 0)
			})
		})
	})

	Convey("Given an non empty tag string with a single value", t, func() {
		tagString := "foobar"
		Convey("When the tag splitter is called", func() {
			fields, e := tagSplit(tagString)
			Convey("Then there is no error", func() {
				So(e, ShouldBeNil)
			})
			Convey("Then the field list contains one entry", func() {
				So(len(fields), ShouldEqual, 1)
			})
			Convey("Then the contained value is set accordingly", func() {
				So(fields[0], ShouldEqual, tagString)
			})
		})
	})

	Convey("Given an non empty tag string with two values separated by one space", t, func() {
		tagString := "foo bar"
		Convey("When the tag splitter is called", func() {
			fields, e := tagSplit(tagString)
			Convey("Then there is no error", func() {
				So(e, ShouldBeNil)
			})
			Convey("Then the field list contains two entries", func() {
				So(len(fields), ShouldEqual, 2)
			})
			Convey("Then the contained values are set accordingly", func() {
				So(fields[0], ShouldEqual, "foo")
				So(fields[1], ShouldEqual, "bar")
			})
		})
	})

	Convey("Given an non empty tag string with two values separated by multiple whitespaces", t, func() {
		tagString := "foo \t bar"
		Convey("When the tag splitter is called", func() {
			fields, e := tagSplit(tagString)
			Convey("Then there is no error", func() {
				So(e, ShouldBeNil)
			})
			Convey("Then the field list contains two entries", func() {
				So(len(fields), ShouldEqual, 2)
			})
			Convey("Then the contained values are set accordingly", func() {
				So(fields[0], ShouldEqual, "foo")
				So(fields[1], ShouldEqual, "bar")
			})
		})
	})

	Convey("Given an non empty tag string with a valid quoted substring", t, func() {
		tagString := "foo'bar buz'"
		Convey("When the tag splitter is called", func() {
			fields, e := tagSplit(tagString)
			Convey("Then there is no error", func() {
				So(e, ShouldBeNil)
			})
			Convey("Then the field list contains two entries", func() {
				So(len(fields), ShouldEqual, 1)
			})
			Convey("Then the contained values are set accordingly", func() {
				So(fields[0], ShouldEqual, tagString)
			})
		})
	})

	Convey("Given an non empty tag string with a invalid quoted substring", t, func() {
		tagString := "prefix foo'bar 'buz'"
		Convey("When the tag splitter is called", func() {
			fields, e := tagSplit(tagString)
			Convey("Then there is an error", func() {
				So(e.Error(), ShouldEqual, "failed to parse tag due to erroneous quotes")
			})
			Convey("Then the field list returned is empty", func() {
				So(len(fields), ShouldEqual, 0)
			})
		})
	})
}

func customHandler(string) (string, string, error) {
	return "", "", fmt.Errorf("failed")
}

func TestParseTag(t *testing.T) {
	Convey("Given an empty tag string", t, func() {
		tagString := ""
		Convey("When the tag parser is called", func() {
			tagMap, e := parseTag(tagString, customHandler)
			Convey("Then there is no error", func() {
				So(e, ShouldBeNil)
			})
			Convey("Then the tag map should be empty", func() {
				So(len(tagMap), ShouldEqual, 0)
			})
		})
	})

	Convey("Given a tag string with a single key value pair", t, func() {
		tagString := "key=value"
		Convey("When the tag parser is called", func() {
			tagMap, e := parseTag(tagString, customHandler)
			Convey("Then there is no error", func() {
				So(e, ShouldBeNil)
			})
			Convey("Then the tag map should contain one key value pair", func() {
				So(len(tagMap), ShouldEqual, 1)
			})
			Convey("Then the tag map should contain the value for the key", func() {
				So(tagMap["key"], ShouldEqual, "value")
			})
		})
	})

	Convey("Given a tag string with two key value pairs", t, func() {
		tagString := "key1=value1 key2=value2"
		Convey("When the tag parser is called", func() {
			tagMap, e := parseTag(tagString, customHandler)
			Convey("Then there is no error", func() {
				So(e, ShouldBeNil)
			})
			Convey("Then the tag map should contain both key value pairs", func() {
				So(len(tagMap), ShouldEqual, 2)
			})
			Convey("Then the tag map should contain the values for the keys", func() {
				So(tagMap["key1"], ShouldEqual, "value1")
				So(tagMap["key2"], ShouldEqual, "value2")
			})
		})
	})

	Convey("Given a tag string with a quoted value", t, func() {
		tagString := "key='quoted value with spaces'"
		Convey("When the tag parser is called", func() {
			tagMap, e := parseTag(tagString, customHandler)
			Convey("Then there is no error", func() {
				So(e, ShouldBeNil)
			})
			Convey("Then the tag map should contain the key value pair", func() {
				So(len(tagMap), ShouldEqual, 1)
			})
			Convey("Then the tag map should contain the proper value for the key, with quotes being removed", func() {
				So(tagMap["key"], ShouldEqual, "quoted value with spaces")
			})
		})
	})

	Convey("Given a tag string with a quoted value with enclosing whitespace", t, func() {
		tagString := "key=' quoted value with spaces \t '"
		Convey("When the tag parser is called", func() {
			tagMap, e := parseTag(tagString, customHandler)
			Convey("Then there is no error", func() {
				So(e, ShouldBeNil)
			})
			Convey("Then the tag map should contain the key value pair", func() {
				So(len(tagMap), ShouldEqual, 1)
			})
			Convey("Then the tag map should contain the proper value for the key, with whitespace being preserved", func() {
				So(tagMap["key"], ShouldEqual, " quoted value with spaces \t ")
			})
		})
	})

	Convey("Given a tag string with a key without value", t, func() {
		tagString := "foo=bar keywithoutvalue"
		Convey("When the tag parser is called", func() {
			tagMap, e := parseTag(tagString, customHandler)
			Convey("Then there is an error", func() {
				So(e, ShouldNotBeNil)
				So(e.Error(), ShouldEqual, `failed to parse annotation (value missing): "keywithoutvalue"`)
			})
			Convey("Then the tag map be nil", func() {
				So(tagMap, ShouldEqual, nil)
			})
		})
	})

}

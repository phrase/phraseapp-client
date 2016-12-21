package paclient

import "testing"

func TestSourceGlob(t *testing.T) {
	s := &Source{File: "a/b/c/d.txt"}
	s2 := &Source{File: "a/**/*/d.txt"}
	s3 := &Source{File: "a/**/**/**/*c.txt"}

	pre, post := s.pattern(lg)
	pre2, post2 := s2.pattern(lg)
	pre3, post3 := s3.pattern(lg)
	tests := map[int]struct{ Has, Want interface{} }{
		1: {pre, "a/b/c/d.txt"},
		2: {post, ""},
		3: {pre2, "a"},
		4: {post2, "*/d.txt"},
		5: {pre3, "a"},
		6: {post3, "**/**/*c.txt"},
	}
	for i, tc := range tests {
		if tc.Has != tc.Want {
			t.Errorf("%d: want=%#v has=%#v", i, tc.Want, tc.Has)
		}
	}
}

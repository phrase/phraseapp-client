package version

import (
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func NewFromString(in string) (*Version, error) {
	v := &Version{}
	return v, v.Parse(in)
}

func New(major, minor, patch int) *Version {
	return &Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

func (version *Version) Parse(raw string) error {
	parts := strings.Split(raw, ".")
	if len(parts) == 3 {
		version.Major, _ = strconv.Atoi(parts[0])
		version.Minor, _ = strconv.Atoi(parts[1])
		version.Patch, _ = strconv.Atoi(parts[2])
		return nil
	}
	return fmt.Errorf("could not parse %s into version", raw)
}

func (version *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch)
}

func Parse(raw string) (v *Version, e error) {
	v = &Version{}
	e = v.Parse(raw)
	return v, e
}

func (version *Version) Less(other *Version) bool {
	if version.Major != other.Major {
		return version.Major < other.Major
	}
	if version.Minor != other.Minor {
		return version.Minor < other.Minor
	}
	return version.Patch < other.Patch
}

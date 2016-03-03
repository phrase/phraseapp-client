package cli

import "testing"

type OptionTestCommand struct {
	Opt1 []string `cli:"opt -o --one"`
	Opt2 []int    `cli:"opt -t --two"`
}

func (cmd *OptionTestCommand) Run() error {
	return nil
}

func TestSliceOption(t *testing.T) {
	cmd := new(OptionTestCommand)
	a := testCreateAction("test", cmd)

	err := a.reflect()
	if err != nil {
		t.Errorf("expected err to be empty, got %s", err)
	}

	err = a.parseArgs([]string{})
	if err != nil {
		t.Errorf("expected err to be empty, got %s", err)
	}
	if cmd.Opt1 != nil {
		t.Errorf("expected first option to be nil, got %s", cmd.Opt1)
	}
	if cmd.Opt2 != nil {
		t.Errorf("expected second option to be nil, got %d", cmd.Opt2)
	}

	err = a.parseArgs([]string{"-o", "a,b"})
	if err != nil {
		t.Errorf("expected err to be empty, got %s", err)
	}
	if cmd.Opt1 == nil {
		t.Errorf("expected first option to not be nil")
	}
	if len(cmd.Opt1) != 2 {
		t.Errorf("expected first option to have 2 elements, got %d", len(cmd.Opt1))
	}
	if cmd.Opt1[0] != "a" {
		t.Errorf("expected first option's first element to be %q, got %q", "a", cmd.Opt1[0])
	}
	if cmd.Opt1[1] != "b" {
		t.Errorf("expected first option's second element to be %q, got %q", "b", cmd.Opt1[1])
	}
	if cmd.Opt2 != nil {
		t.Errorf("expected second option to be nil, got %d", cmd.Opt2)
	}

	err = a.parseArgs([]string{"-t", "4,5,6"})
	if err != nil {
		t.Errorf("expected err to be empty, got %s", err)
	}
	if cmd.Opt2 == nil {
		t.Errorf("expected first option to not be nil")
	}
	if len(cmd.Opt2) != 3 {
		t.Errorf("expected first option to have 2 elements, got %d", len(cmd.Opt1))
	}
	if cmd.Opt2[0] != 4 {
		t.Errorf("expected first option's first element to be %d, got %d", 4, cmd.Opt2[0])
	}
	if cmd.Opt2[1] != 5 {
		t.Errorf("expected first option's second element to be %d, got %d", 5, cmd.Opt2[1])
	}
	if cmd.Opt2[2] != 6 {
		t.Errorf("expected first option's second element to be %d, got %d", 6, cmd.Opt2[2])
	}
}

type OptionWPtrTestCommand struct {
	Opt *string `cli:"opt -o --one"`
}

func (cmd *OptionWPtrTestCommand) Run() error {
	return nil
}

func TestPtrOption(t *testing.T) {
	cmd := new(OptionWPtrTestCommand)
	a := testCreateAction("test", cmd)

	err := a.reflect()
	if err != nil {
		t.Errorf("expected err to be empty, got %s", err)
	}

	err = a.parseArgs([]string{})
	if err != nil {
		t.Fatalf("expected err to be empty, got %q", err)
	}
	if cmd.Opt != nil {
		t.Fatalf("expected option to be nil, got %q", *cmd.Opt)
	}

	err = a.parseArgs([]string{"-o", ""})
	if err != nil {
		t.Fatalf("expected err to be empty, got %q", err)
	}
	if cmd.Opt == nil {
		t.Fatalf("expected option not to be nil, but was")
	} else if *cmd.Opt != "" {
		t.Fatalf("expected option to be %q, got %q", "", *cmd.Opt)
	}

	err = a.parseArgs([]string{"-o", "a"})
	if err != nil {
		t.Fatalf("expected err to be empty, got %q", err)
	}
	if cmd.Opt == nil {
		t.Fatalf("expected option not to be nil, but was")
	} else if *cmd.Opt != "a" {
		t.Fatalf("expected option to be %q, got %q", "a", *cmd.Opt)
	}
}

func TestPtrOptionWithPreset(t *testing.T) {
	s := "test1234"
	cmd := new(OptionWPtrTestCommand)
	cmd.Opt = &s
	a := testCreateAction("test", cmd)

	err := a.reflect()
	if err != nil {
		t.Errorf("expected err to be empty, got %s", err)
	}

	if a.opts[0].value != "test1234" {
		t.Errorf("expected option default value to be %q, got %q", s, a.opts[0].value)
	}
}

type OptionMapTestCommand struct {
	First  map[string]string `cli:"opt -f --first"`
	Second map[string]int64  `cli:"opt -s --second"`
}

func (cmd *OptionMapTestCommand) Run() error {
	return nil
}

func TestMapOption(t *testing.T) {
	cmd := new(OptionMapTestCommand)
	a := testCreateAction("test", cmd)

	err := a.reflect()
	if err != nil {
		t.Fatalf("expected err to be empty, got %s", err)
	}

	err = a.parseArgs([]string{})
	if err != nil {
		t.Errorf("expected err to be empty, got %s", err)
	}
	if cmd.First != nil {
		t.Errorf("expected first option to be nil, got %#v", cmd.First)
	}
	if cmd.Second != nil {
		t.Errorf("expected second option to be nil, got %#v", cmd.Second)
	}

	err = a.parseArgs([]string{"-f.a", "a", "-f.b", "b"})
	if err != nil {
		t.Errorf("expected err to be empty, got %s", err)
	}
	if v, found := cmd.First["a"]; !found || v != "a" {
		t.Errorf("expected value for %q to be %q, got %q", "a", "a", v)
	}
	if v, found := cmd.First["b"]; !found || v != "b" {
		t.Errorf("expected value for %q to be %q, got %q", "b", "b", v)
	}

	err = a.parseArgs([]string{"-s.c", "1", "-s.d", "2"})
	if err != nil {
		t.Errorf("expected err to de empty, got %s", err)
	}
	if v, found := cmd.Second["c"]; !found || v != 1 {
		t.Errorf("expected value for %d to be %d, got %d", 1, 1, v)
	}
	if v, found := cmd.Second["d"]; !found || v != 2 {
		t.Errorf("expected value for %d to be %d, got %d", 2, 2, v)
	}
}

type OptionFailBoolMapCommand struct {
	First map[string]bool `cli:"opt -m"`
}

func (cmd *OptionFailBoolMapCommand) Run() error {
	return nil
}

type OptionFailIntMapCommand struct {
	First map[string]int `cli:"opt -m"`
}

func (cmd *OptionFailIntMapCommand) Run() error {
	return nil
}

func TestFailMapOption(t *testing.T) {
	{
		cmd := new(OptionFailBoolMapCommand)
		a := testCreateAction("test", cmd)

		expErr := "OptionFailBoolMapCommand: Options with Map type must have string or int64 values, got bool"
		err := a.reflect()
		switch {
		case err == nil:
			t.Errorf("expected error, got none")
		case err.Error() != expErr:
			t.Errorf("expected error to be %q, got %q", expErr, err)
		}
	}

	{
		cmd := new(OptionFailIntMapCommand)
		a := testCreateAction("test", cmd)

		expErr := "OptionFailIntMapCommand: Options with Map type must have string or int64 values, got int"
		err := a.reflect()
		switch {
		case err == nil:
			t.Errorf("expected error, got none")
		case err.Error() != expErr:
			t.Errorf("expected error to be %q, got %q", expErr, err)
		}
	}
}

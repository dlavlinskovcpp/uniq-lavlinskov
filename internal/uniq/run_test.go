package uniq

import (
	"strings"
	"testing"
)

func TestRun_Basic(t *testing.T) {
	in := "a\na\nb\n"
	var out strings.Builder
	err := Run(strings.NewReader(in), &out, &Options{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "a\nb\n"
	if got := out.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRun_Count(t *testing.T) {
	in := "a\na\nb\n"
	var out strings.Builder
	err := Run(strings.NewReader(in), &out, &Options{Count: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "2 a\n1 b\n"
	if got := out.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRun_Dups(t *testing.T) {
	in := "x\nx\ny\nz\nz\n"
	var out strings.Builder
	err := Run(strings.NewReader(in), &out, &Options{Dups: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "x\nz\n"
	if got := out.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRun_Uniq(t *testing.T) {
	in := "x\nx\ny\nz\nz\n"
	var out strings.Builder
	err := Run(strings.NewReader(in), &out, &Options{Uniq: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "y\n"
	if got := out.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRun_IgnoreCase(t *testing.T) {
	in := "Hi\nhi\nHello\n"
	var out strings.Builder
	err := Run(strings.NewReader(in), &out, &Options{IgnoreCase: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "Hi\nHello\n"
	if got := out.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRun_SkipFields(t *testing.T) {
	in := "1 apple\n2 apple\n3 banana\n"
	var out strings.Builder
	err := Run(strings.NewReader(in), &out, &Options{SkipFields: 1})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "1 apple\n3 banana\n"
	if got := out.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRun_UTF8(t *testing.T) {
	in := "тест\nтест\nпример\n"
	var out strings.Builder
	err := Run(strings.NewReader(in), &out, &Options{SkipChars: 2})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "тест\nпример\n"
	if got := out.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// +build !bonus

package robotname

import (
	"testing"
)

func TestGenerateNames(t *testing.T) {
	names := GenerateNames()
	if len(names) != NumNames {
		t.Errorf(`Only generated %v names, want %v.`, len(names), NumNames)
	}
	if names[0] != "AA000" {
		t.Errorf(`Generated %v want %v.`, names[0], "AA000")
	}
	if names[NumNames-1] != "ZZ999" {
		t.Errorf(`Generated %v, want %v.`, names[NumNames-1], "ZZ999")
	}
}

func TestNameValid(t *testing.T) {
	n := New().getName(t, false)
	if !namePat.MatchString(n) {
		t.Errorf(`Invalid robot name %q, want form "AA###".`, n)
	}
}

func TestNameSticks(t *testing.T) {
	r := New()
	n1 := r.getName(t, false)
	n2 := r.getName(t, true)
	if n2 != n1 {
		t.Errorf(`Robot name changed.  Now %s, was %s.`, n2, n1)
	}
}

func TestSuccessiveRobotsHaveDifferentNames(t *testing.T) {
	n1 := New().getName(t, false)
	n2 := New().getName(t, false)
	if n1 == n2 {
		t.Errorf(`Robots with same name.  Two %s's.`, n1)
	}
}

func TestResetName(t *testing.T) {
	r := New()
	n1 := r.getName(t, false)
	r.Reset()
	if r.getName(t, false) == n1 {
		t.Errorf(`Robot name not cleared on reset.  Still %s.`, n1)
	}
}

// Note if you go for bonus points, this benchmark likely won't be
// meaningful.  Bonus thought exercise, why won't it be meaningful?
func BenchmarkName(b *testing.B) {
	// Benchmark combined time to create robot and name.
	for i := 0; i < b.N; i++ {
		New().getName(b, false)
	}
}

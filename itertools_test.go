package itertools

import (
	"testing"
)

func TestCountStep(t *testing.T) {
	expected := [...]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	size := len(expected)
	i := 0
	for val := range CountStep(2, 2) {
		if i == size {
			break
		}
		if expected[i] != val {
			t.Fail()
		}
		i += 1
	}
}

func TestCount(t *testing.T) {
	expected := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	size := len(expected)
	i := 0
	for val := range Count(1) {
		if i == size {
			break
		}
		if expected[i] != val {
			t.Fail()
		}
		i += 1
	}
}

func TestCycleByte(t *testing.T) {
	s := "abc"
	expected := s + s + s + s
	size := len(expected)
	i := 0
	for val := range CycleByte(s) {
		if i == size {
			break
		}
		if expected[i] != val {
			t.Fail()
		}
		i += 1
	}
}

//func TestCycleRune(t *testing.T) {
//	s := "日本語⌘"
//	expected := s
//	size := len(expected)
//	i := 0
//	for val := range CycleRune(s) {
//		if i == size {
//			break
//		}
//		if r, _ := utf8.DecodeRuneInString(expected[:i]); r != val {
//			t.Logf("Expected: %U. Actual: %U", r, val)
//			t.FailNow()
//		}
//		i += 1
//	}
//}

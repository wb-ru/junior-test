package secondres

import "testing"

func TestSum(tes *testing.T) {

	tes.Run("Input: s = egg, t=add", func(tes *testing.T) {
		s := "egg"
		t := "add"
		got := isIsomorphic(s, t)
		target := true

		if got != target {
			tes.Errorf("got %v target %v. Given: %s, %s", got, target, s, t)
		}
	})

	tes.Run("Input: s = foo, t=bar", func(tes *testing.T) {
		s := "foo"
		t := "bar"
		got := isIsomorphic(s, t)
		target := false

		if got != target {
			tes.Errorf("got %v target %v. Given: %s, %s", got, target, s, t)
		}
	})

	tes.Run("Input: s = paper, t=title", func(tes *testing.T) {
		s := "paper"
		t := "title"
		got := isIsomorphic(s, t)
		target := true

		if got != target {
			tes.Errorf("got %v target %v. Given: %s, %s", got, target, s, t)
		}
	})
}

package fizzbuzz_test

import (
	"testing"

	"github.com/tarathep/go101/t_unittest/fizzbuzz"
)

func TestInput1ShouldBeDisplay1(t *testing.T) {
	v := fizzbuzz.Fizzbuzz(1)
	if "1" != v {
		t.Error("fizzbuzz of 1 should be '1' but have", v)
	}
}

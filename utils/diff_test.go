package utils

import (
	"testing"
)

func TestSwitchDiff(t *testing.T) {
	f := SwitchDiff("115792089191727444327593391681087395856323862999734652148720336070939540036993")
	t.Log("diff in small number: ", f)
}

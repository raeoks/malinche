package malinche

import "testing"

func TestConstants(runner *testing.T) {
	if Version != "0.0.1" {
		runner.Errorf("malinche version does not match")
	}
	if Build == "" {
		runner.Errorf("malinche build date is empty")
	}
}

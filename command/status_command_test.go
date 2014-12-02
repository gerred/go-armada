package command

import "testing"

func TestNewStatusCommand(t *testing.T) {
	command := NewStatusCommand()

	if command.Name != "status" {
		t.Error("Expected status, got ", command.Name)
	}
}

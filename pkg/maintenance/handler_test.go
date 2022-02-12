package maintenance_test

import (
	"testing"

	"github.com/andygeiss/meridian59-build/pkg/maintenance"
)

func Test_Maintenance_Who(t *testing.T) {
	handler := maintenance.NewHandler()
	handler.Connect("127.0.0.1:59595")
	handler.Send("who")
	out := handler.Receive()
	err := handler.Error()
	if err != nil {
		t.Fatalf("error should be nil, but got %v", err)
	}
	if len(out) == 0 {
		t.Fatal("out should not be empty")
	}
}

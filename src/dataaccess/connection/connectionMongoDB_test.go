package connection

import (
	"github.com/Jacobbrewer1/chess-boards/src/custom"
	"testing"
)

func TestMongoDB_GetCollection(t *testing.T) {
	conn := MongoDB{
		Host:     "TestHost",
		Port:     "12345",
		Database: "TestDB",
	}

	conn.Collections = make(custom.Map[string, string])

	conn.Collections.Set("one", "1")
	conn.Collections.Set("two", "2")
	conn.Collections.Set("three", "3")
	conn.Collections.Set("four", "4")
	conn.Collections.Set("five", "5")

	if _, err := conn.GetCollection("one"); err != nil {
		t.Errorf("GetCollection() error: %s", err)
	}

	if _, err := conn.GetCollection("two"); err != nil {
		t.Errorf("GetCollection() error: %s", err)
	}

	if _, err := conn.GetCollection("three"); err != nil {
		t.Errorf("GetCollection() error: %s", err)
	}

	if _, err := conn.GetCollection("four"); err != nil {
		t.Errorf("GetCollection() error: %s", err)
	}

	if _, err := conn.GetCollection("five"); err != nil {
		t.Errorf("GetCollection() error: %s", err)
	}

	if _, err := conn.GetCollection("six"); err == nil {
		t.Errorf("GetCollection() expected error but returned nil")
	}
}

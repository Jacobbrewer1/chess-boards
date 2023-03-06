package connection

import (
	"testing"
)

func TestMongoDB_generateConnectionString_expectPanic(t *testing.T) {
	conn := MySql{
		User:     "",
		Password: "",
		Method:   "",
		Host:     "",
		Port:     "",
		Schema:   "",
	}

	var stringGenerated = false
	var completed = false
	defer func() {
		if stringGenerated || completed {
			t.Errorf("panic exected before end of program")
		} else {
			recovery := recover()
			if recovery != "invalid connection" {
				t.Errorf("panic recovered: %s", recovery)
			}
		}
	}()

	expected := "TestUser:TestPsw@TestMethod(TestHost:12345)/TestSchema"
	conn.generateConnectionString()

	stringGenerated = true

	if *conn.connectionString != expected {
		t.Errorf("generateConnectionString() got %s, expected %s", *conn.connectionString, expected)
	}

	completed = true
}

func TestMongoDB_generateConnectionString_noQuery(t *testing.T) {
	conn := MySql{
		User:     "TestUser",
		Password: "TestPsw",
		Method:   "TestMethod",
		Host:     "TestHost",
		Port:     "12345",
		Schema:   "TestSchema",
	}

	var stringGenerated = false
	defer func() {
		if !stringGenerated {
			t.Errorf("panic during connection string generation")
		}
	}()

	expected := "TestUser:TestPsw@TestMethod(TestHost:12345)/TestSchema"
	conn.generateConnectionString()

	stringGenerated = true

	if *conn.connectionString != expected {
		t.Errorf("generateConnectionString() got %s, expected %s", *conn.connectionString, expected)
	}
}

func TestMongoDB_generateConnectionString_withQuery(t *testing.T) {
	query := "timeout=2s&parseTime=true"
	conn := MySql{
		User:     "TestUser",
		Password: "TestPsw",
		Method:   "TestMethod",
		Host:     "TestHost",
		Port:     "12345",
		Schema:   "TestSchema",
		Query:    &query,
	}

	var stringGenerated = false
	defer func() {
		if !stringGenerated {
			t.Errorf("panic during connection string generation")
		}
	}()

	expected := "TestUser:TestPsw@TestMethod(TestHost:12345)/TestSchema?timeout=2s&parseTime=true"
	conn.generateConnectionString()

	stringGenerated = true

	if *conn.connectionString != expected {
		t.Errorf("generateConnectionString() got %s, expected %s", *conn.connectionString, expected)
	}
}

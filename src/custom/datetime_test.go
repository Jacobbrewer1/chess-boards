package custom

import (
	"log"
	"testing"
	"time"
)

func TestDateTime_Scan(t *testing.T) {
	var d Datetime

	err := d.Scan("2000-09-24 15:21:35 +0000 UTC")
	if err != nil {
		t.Error(err)
		return
	}

	parse, err := time.ParseInLocation(DateTimeLayout, "2000-09-24 15:21:35", time.Now().UTC().Location())
	if err != nil {
		log.Println(err)
		return
	}

	if time.Time(d) != parse.In(time.Now().Location()) {
		t.Errorf("times not matching got %v, expected %v", time.Time(d), parse)
	}
}

func TestDateTime_Display(t *testing.T) {
	parseTime := func(string2 string) Datetime {
		parsed, _ := time.Parse(time.RFC3339, string2)
		return Datetime(parsed)
	}

	tests := []struct {
		name     string
		input    Datetime
		expected string
	}{
		{"Jan", parseTime("2000-01-01T15:04:05Z"), "Saturday, 01 Jan 2000 15:04:05"},
		{"Feb", parseTime("2001-02-02T15:04:05Z"), "Friday, 02 Feb 2001 15:04:05"},
		{"Mar", parseTime("2002-03-03T15:04:05Z"), "Sunday, 03 Mar 2002 15:04:05"},
		{"Apr", parseTime("2003-04-04T15:04:05Z"), "Friday, 04 Apr 2003 15:04:05"},
		{"May", parseTime("2004-05-05T15:04:05Z"), "Wednesday, 05 May 2004 15:04:05"},
		{"Jun", parseTime("2005-06-06T15:04:05Z"), "Monday, 06 Jun 2005 15:04:05"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Display()
			if got != tt.expected {
				t.Errorf("Date.Display() = %s, expected %s", got, tt.expected)
			}
		})
	}
}

func TestDateTime_MarshalText(t *testing.T) {
	parseTime := func(string2 string) Datetime {
		parsed, _ := time.Parse(time.RFC3339, string2)
		return Datetime(parsed)
	}

	tests := []struct {
		name     string
		input    Datetime
		expected string
	}{
		{"Jan", parseTime("2000-01-01T15:04:05Z"), "2000-01-01T15:04:05Z"},
		{"Feb", parseTime("2001-02-02T15:04:05Z"), "2001-02-02T15:04:05Z"},
		{"Mar", parseTime("2002-03-03T15:04:05Z"), "2002-03-03T15:04:05Z"},
		{"Apr", parseTime("2003-04-04T15:04:05Z"), "2003-04-04T15:04:05Z"},
		{"May", parseTime("2004-05-05T15:04:05Z"), "2004-05-05T15:04:05Z"},
		{"Jun", parseTime("2005-06-06T15:04:05Z"), "2005-06-06T15:04:05Z"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := tt.input.MarshalText()
			if string(got) != tt.expected {
				t.Errorf("Date.Display() = %s, expected %s", got, tt.expected)
			}
		})
	}
}

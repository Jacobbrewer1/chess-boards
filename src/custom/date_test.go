package custom

import (
	"log"
	"testing"
	"time"
)

func TestDate_Scan(t *testing.T) {
	var d Date

	if err := d.Scan("2000-09-24 00:00:00 +0000 UTC"); err != nil {
		t.Error(err)
		return
	}

	parse, err := time.Parse(DateLayout, "2000-09-24")
	if err != nil {
		log.Println(err)
		return
	}

	if time.Time(d) != parse {
		t.Errorf("times not matching got %v, expected %v", time.Time(d), parse)
	}
}

func TestDate_Display(t *testing.T) {
	parseTime := func(string2 string) Date {
		parsed, _ := time.Parse(DateLayout, string2)
		return Date(parsed)
	}

	tests := []struct {
		name     string
		input    Date
		expected string
	}{
		{"Jan", parseTime("2000-01-01"), "Saturday, 01 Jan 2000"},
		{"Feb", parseTime("2001-02-02"), "Friday, 02 Feb 2001"},
		{"Mar", parseTime("2002-03-03"), "Sunday, 03 Mar 2002"},
		{"Apr", parseTime("2003-04-04"), "Friday, 04 Apr 2003"},
		{"May", parseTime("2004-05-05"), "Wednesday, 05 May 2004"},
		{"Jun", parseTime("2005-06-06"), "Monday, 06 Jun 2005"},
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

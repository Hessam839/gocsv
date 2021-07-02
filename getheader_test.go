package gocsv

import "testing"

type Q struct {
	Name   string `csv:"name"`
	Family string `csv:"family"`
	Age    int    `csv:"-"`
}

func Test_GetHeader(t *testing.T) {
	s := struct {
		Name   string `csv:"name"`
		Family string `csv:"family"`
	}{}

	header := GetHeader(s)
	t.Log(header)
}

func Test_WriteCSV(t *testing.T) {
	s := []Q{
		{
			Name:   "ali",
			Family: "ahmadi",
			Age:    12,
		},
		{
			Name:   "reza",
			Family: "naderi",
			Age:    32,
		},
		{
			Name:   "نادر",
			Family: "محمودی",
			Age:    44,
		},
	}

	_ = WriteToCSV("test.csv", s)
}

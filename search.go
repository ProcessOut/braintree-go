package braintree

import (
	"encoding/xml"
	"time"
)

type SearchQuery struct {
	XMLName string `xml:"search"`
	Fields  []interface{}
}

type SearchResults struct {
	XMLName  string `xml:"search-results"`
	PageSize string `xml:"page-size"`
	Ids      struct {
		Item []string `xml:"item"`
	} `xml:"ids"`
}

type TextField struct {
	XMLName    xml.Name
	Is         *string `xml:"is,omitempty"`
	IsNot      *string `xml:"is-not,omitempty"`
	StartsWith *string `xml:"starts-with,omitempty"`
	EndsWith   *string `xml:"ends-with,omitempty"`
	Contains   *string `xml:"contains,omitempty"`
}

type RangeField struct {
	XMLName xml.Name
	Is      float64 `xml:"is,omitempty"`
	Min     float64 `xml:"min,omitempty"`
	Max     float64 `xml:"max,omitempty"`
}

type TimeRangeField struct {
	XMLName xml.Name
	Is      *DateTime `xml:"is,omitempty"`
	Min     *DateTime `xml:"min,omitempty"`
	Max     *DateTime `xml:"max,omitempty"`
}

const iso8601 = "2006-01-02T15:04:05-0700"

func (t *TimeRangeField) SetIs(date *time.Time) {
	t.Is = &DateTime{
		Value: date.Format(iso8601),
	}
}
func (t *TimeRangeField) SetMax(date *time.Time) {
	t.Max = &DateTime{
		Value: date.Format(iso8601),
	}
}
func (t *TimeRangeField) SetMin(date *time.Time) {
	t.Min = &DateTime{
		Value: date.Format(iso8601),
	}
}

type DateTime struct {
	Value string `xml:",chardata"`
}

type MultiField struct {
	XMLName xml.Name
	Type    string   `xml:"type,attr"` // type=array
	Items   []string `xml:"item"`
}

func (s *SearchQuery) AddTextField(field string) *TextField {
	f := &TextField{XMLName: xml.Name{Local: field}}
	s.Fields = append(s.Fields, f)
	return f
}

func (s *SearchQuery) AddRangeField(field string) *RangeField {
	f := &RangeField{XMLName: xml.Name{Local: field}}
	s.Fields = append(s.Fields, f)
	return f
}

func (s *SearchQuery) AddTimeRangeField(field string) *TimeRangeField {
	f := &TimeRangeField{XMLName: xml.Name{Local: field}}
	s.Fields = append(s.Fields, f)
	return f
}

func (s *SearchQuery) AddMultiField(field string) *MultiField {
	f := &MultiField{
		XMLName: xml.Name{Local: field},
		Type:    "array",
	}
	s.Fields = append(s.Fields, f)
	return f
}

// generated by jsonenums -type=parseKind; DO NOT EDIT

package main

import (
	"encoding/json"
	"fmt"
)

var (
	_parseKindNameToValue = map[string]parseKind{
		"parseRaw":      parseRaw,
		"parseHTML":     parseHTML,
		"parseMarkdown": parseMarkdown,
	}

	_parseKindValueToName = map[parseKind]string{
		parseRaw:      "parseRaw",
		parseHTML:     "parseHTML",
		parseMarkdown: "parseMarkdown",
	}
)

func init() {
	var v parseKind
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_parseKindNameToValue = map[string]parseKind{
			interface{}(parseRaw).(fmt.Stringer).String():      parseRaw,
			interface{}(parseHTML).(fmt.Stringer).String():     parseHTML,
			interface{}(parseMarkdown).(fmt.Stringer).String(): parseMarkdown,
		}
	}
}

// MarshalJSON is generated so parseKind satisfies json.Marshaler.
func (r parseKind) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _parseKindValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid parseKind: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so parseKind satisfies json.Unmarshaler.
func (r *parseKind) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("parseKind should be a string, got %s", data)
	}
	v, ok := _parseKindNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid parseKind %q", s)
	}
	*r = v
	return nil
}

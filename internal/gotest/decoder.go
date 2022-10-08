// =====================================================================================================================
// = LICENSE:       Copyright (c) 2022 Kevin De Coninck
// =
// =                Permission is hereby granted, free of charge, to any person
// =                obtaining a copy of this software and associated documentation
// =                files (the "Software"), to deal in the Software without
// =                restriction, including without limitation the rights to use,
// =                copy, modify, merge, publish, distribute, sublicense, and/or sell
// =                copies of the Software, and to permit persons to whom the
// =                Software is furnished to do so, subject to the following
// =                conditions:
// =
// =                The above copyright notice and this permission notice shall be
// =                included in all copies or substantial portions of the Software.
// =
// =                THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// =                EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// =                OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// =                NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// =                HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// =                WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// =                FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// =                OTHER DEALINGS IN THE SOFTWARE.
// =====================================================================================================================

// Package gotest implements functions for working with Go's standard testing.
package gotest

import (
	"encoding/json"
	"fmt"
)

// TEvent is a single line of Go's standard testing output (formatted as JSON).
type TEvent struct {
	// Action is the executed action on this event, such as: "run", "output", ...
	Action string `json:"Action"`

	// Package is the Go package this event belongs to.
	Package string `json:"Package"`

	// Test is the name of the test this event belongs to.
	Test string `json:"Test"`

	// Elapsed is the total amount of elapsed seconds for this event.
	Elapsed float64 `json:"Elapsed"`
}

// ParseJSONOutput reads data into a TEvent.
func ParseJSONOutput(data string) (TEvent, error) {
	event := TEvent{}
	if err := json.Unmarshal([]byte(data), &event); err != nil {
		return event, fmt.Errorf("failed to decode JSON. %w", err)
	}
	return event, nil
}

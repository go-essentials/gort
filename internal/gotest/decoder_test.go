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

// Verify and measure the performance of the public API of the gotest package.
package gotest_test

import (
	"testing"

	"github.com/go-essentials/gort/internal/gotest"
)

// Verify that the TEvent.String function is implemented correctly.
func TestTEventString(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// DEFINITIONS.
	scenarios := []struct {
		name     string
		evt      gotest.TEvent
		expected string
	}{
		{
			name: "When the event is marked as `passed`, it's formatted correctly (Sample 1).",
			evt: gotest.TEvent{
				Action: "pass",
				Test:   "A_sample_test",
			},
			expected: "\033[32m✓ A sample test\033[0m",
		},
		{
			name: "When the event is marked as `passed`, it's formatted correctly (Sample 2).",
			evt: gotest.TEvent{
				Action: "pass",
				Test:   "Test/Test2",
			},
			expected: "\033[32m✓ Test2\033[0m",
		},
		{
			name: "When the event is marked as `failed`, it's formatted correctly (Sample 1).",
			evt: gotest.TEvent{
				Action: "fail",
				Test:   "A_sample_test",
			},
			expected: "\033[31m✗ A sample test\033[0m",
		},
		{
			name: "When the event is marked as `failed`, it's formatted correctly (Sample 2).",
			evt: gotest.TEvent{
				Action: "fail",
				Test:   "Test/Test2",
			},
			expected: "\033[31m✗ Test2\033[0m",
		},
	}

	// EXECUTION.
	for _, scenario := range scenarios {
		scenario := scenario // NOTE: Ensure that the t.Run function has the correct value when it's being executed.

		// EXECUTION.
		t.Run(scenario.name, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// ACT.
			result := scenario.evt.String()

			// ASSERT.
			if scenario.expected != result {
				t.Fatalf("\n\n"+
					"UT Name:  %s\n"+
					"\033[32mExpected: %q\033[0m\n"+
					"\033[31mActual:   %q\033[0m\n\n",
					scenario.name, scenario.expected, result)
			}
		})
	}
}

// Measure the performance of the TEvent.String function.
func BenchmarkTEventString(b *testing.B) {
	// DEFINITIONS.
	scenarios := []struct {
		name     string
		evt      gotest.TEvent
		expected string
	}{
		{
			name: "When the event is marked as `passed` (Sample 1).",
			evt: gotest.TEvent{
				Action: "pass",
				Test:   "A_sample_test",
			},
			expected: "\033[32m✓ A sample test\033[0m",
		},
		{
			name: "When the event is marked as `passed` (Sample 2).",
			evt: gotest.TEvent{
				Action: "pass",
				Test:   "Test/Test2",
			},
			expected: "\033[32m✓ Test2\033[0m",
		},
		{
			name: "When the event is marked as `failed` (Sample 1).",
			evt: gotest.TEvent{
				Action: "fail",
				Test:   "A_sample_test",
			},
			expected: "\033[31m✗ A sample test\033[0m",
		},
		{
			name: "When the event is marked as `failed` (Sample 2).",
			evt: gotest.TEvent{
				Action: "fail",
				Test:   "Test/Test2",
			},
			expected: "\033[31m✗ Test2\033[0m",
		},
	}

	// EXECUTION.
	for _, scenario := range scenarios {
		scenario := scenario // NOTE: Ensure that the b.Run function has the correct value when it's being executed.

		b.Run(scenario.name, func(b *testing.B) {
			// RUN.
			for n := 0; n < b.N; n++ {
				// ACT.
				_ = scenario.evt.String()
			}
		})
	}
}

// Verify that the TEvent.IsRelevant function is implemented correctly.
func TestTEventIsRelevantTrue(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// DEFINITIONS.
	scenarios := []struct {
		name string
		evt  gotest.TEvent
	}{
		{
			name: "When the event is marked as `passed` and it does represent a test, the event should be relevant.",
			evt: gotest.TEvent{
				Action: "pass",
				Test:   "My Test",
			},
		},

		{
			name: "When the event is marked as `failed` and it does represent a test, the event should be relevant.",
			evt: gotest.TEvent{
				Action: "fail",
				Test:   "My Test 2",
			},
		},
	}

	// EXECUTION.
	for _, scenario := range scenarios {
		scenario := scenario // NOTE: Ensure that the t.Run function has the correct value when it's being executed.

		// EXECUTION.
		t.Run(scenario.name, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// ACT / ASSERT.
			if !scenario.evt.IsRelevant() {
				t.Fatalf("\n\n"+
					"UT Name:  %s\n"+
					"\033[32mExpected: true\033[0m\n"+
					"\033[31mActual:   false\033[0m\n\n",
					scenario.name)
			}
		})
	}
}

// Verify that the TEvent.IsRelevant function is implemented correctly.
func TestTEventIsRelevantFalse(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// DEFINITIONS.
	scenarios := []struct {
		name string
		evt  gotest.TEvent
	}{
		{
			name: "When the event is marked as `passed`, but it does NOT represent a test, the event should NOT be " +
				"relevant.",
			evt: gotest.TEvent{
				Action: "pass",
			},
		},
		{
			name: "When the event is marked as `failed`, but it does NOT represent a test, the event should NOT be " +
				"relevant.",
			evt: gotest.TEvent{
				Action: "fail",
			},
		},
		{
			name: "When the event is NOT marked as `passed` or `failed`, but it does represent a test, the event " +
				"should NOT be relevant.",
			evt: gotest.TEvent{
				Test: "My test",
			},
		},
	}

	// EXECUTION.
	for _, scenario := range scenarios {
		scenario := scenario // NOTE: Ensure that the t.Run function has the correct value when it's being executed.

		// EXECUTION.
		t.Run(scenario.name, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// ACT / ASSERT.
			if scenario.evt.IsRelevant() {
				t.Fatalf("\n\n"+
					"UT Name:  %s\n"+
					"\033[32mExpected: false\033[0m\n"+
					"\033[31mActual:   true\033[0m\n\n",
					scenario.name)
			}
		})
	}
}

// Measure the performance of the TEvent.IsRelevant function.
func BenchmarkTEventIsRelevant(b *testing.B) {
	// DEFINITIONS.
	scenarios := []struct {
		name string
		evt  gotest.TEvent
	}{
		{
			name: "When the event is marked as `passed` and it does represent a test.",
			evt: gotest.TEvent{
				Action: "pass",
				Test:   "My Test",
			},
		},
		{
			name: "When the event is marked as `failed` and it does represent a test.",
			evt: gotest.TEvent{
				Action: "fail",
				Test:   "My Test 2",
			},
		},
		{
			name: "When the event is marked as `passed`, but it does NOT represent a test.",
			evt: gotest.TEvent{
				Action: "pass",
			},
		},

		{
			name: "When the event is marked as `failed`, but it does NOT represent a test.",
			evt: gotest.TEvent{
				Action: "fail",
			},
		},

		{
			name: "When the event is NOT marked as `passed` or `failed`, but it does represent a test.",
			evt: gotest.TEvent{
				Test: "My test",
			},
		},
	}

	// EXECUTION.
	for _, scenario := range scenarios {
		scenario := scenario // NOTE: Ensure that the b.Run function has the correct value when it's being executed.

		b.Run(scenario.name, func(b *testing.B) {
			// RUN.
			for n := 0; n < b.N; n++ {
				// ACT.
				_ = scenario.evt.IsRelevant()
			}
		})
	}
}

// Verify that the ParseJSONOutput function is implemented correctly.
func TestParseJSONOutputErr(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// DEFINITIONS.
	scenarios := []struct {
		name        string
		inputString string
	}{
		{
			name:        "When the input is an empty string, an error is returned.",
			inputString: "",
		},
		{
			name:        "When the input is NOT a JSON document, an error is returned.",
			inputString: "N.A.",
		},
	}

	// EXECUTION.
	for _, scenario := range scenarios {
		scenario := scenario // NOTE: Ensure that the t.Run function has the correct value when it's being executed.

		// EXECUTION.
		t.Run(scenario.name, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// ACT.
			_, err := gotest.ParseJSONOutput(scenario.inputString)

			// ASSERT.
			if nil == err {
				t.Fatalf("\n\n"+
					"UT Name:  %s\n"+
					"\033[32mExpected: Any error.\033[0m\n"+
					"\033[31mActual:   %v\033[0m\n\n",
					scenario.name, err)
			}
		})
	}
}

// Verify that the ParseJSONOutput function is implemented correctly.
func TestParseJSONOutputNoErr(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// DEFINITIONS.
	scenarios := []struct {
		name                                          string
		inputString                                   string
		expectedAction, expectedPackage, expectedTest string
		expectedElapsed                               float64
	}{
		{
			name:            "When the input is an empty JSON document, it's parsed correctly.",
			inputString:     "{}",
			expectedAction:  "",
			expectedPackage: "",
			expectedTest:    "",
			expectedElapsed: 0,
		},
		{
			name:            "When the input is a valid JSON document, it's parsed correctly (Sample 1).",
			inputString:     "{\"Action\":\"run\", \"Package\":\"pkg1\", \"Test\":\"tst1\",\"Elapsed\":5}",
			expectedAction:  "run",
			expectedPackage: "pkg1",
			expectedTest:    "tst1",
			expectedElapsed: 5,
		},
		{
			name:            "When the input is a valid JSON document, it's parsed correctly (Sample 2).",
			inputString:     "{\"Action\":\"output\", \"Package\":\"pkg2\", \"Test\":\"tst2\",\"Elapsed\":15}",
			expectedAction:  "output",
			expectedPackage: "pkg2",
			expectedTest:    "tst2",
			expectedElapsed: 15,
		},
	}

	// EXECUTION.
	for _, scenario := range scenarios {
		scenario := scenario // NOTE: Ensure that the t.Run function has the correct value when it's being executed.

		// EXECUTION.
		t.Run(scenario.name, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// ACT.
			result, err := gotest.ParseJSONOutput(scenario.inputString)

			// ASSERT.
			if nil != err {
				t.Fatalf("\n\n"+
					"UT Name:  %s\n"+
					"\033[32mExpected: NO error.\033[0m\n"+
					"\033[31mActual:   %s\033[0m\n\n",
					scenario.name, err)
			}

			if scenario.expectedAction != result.Action {
				t.Fatalf("\n\n"+
					"UT Name:  %s\n"+
					"\033[32mExpected (Action): %s\033[0m\n"+
					"\033[31mActual (Action):   %s\033[0m\n\n",
					scenario.name, scenario.expectedAction, result.Action)
			}

			if scenario.expectedPackage != result.Package {
				t.Fatalf("\n\n"+
					"UT Name:  %s\n"+
					"\033[32mExpected (Package): %s\033[0m\n"+
					"\033[31mActual (Package):   %s\033[0m\n\n",
					scenario.name, scenario.expectedPackage, result.Package)
			}

			if scenario.expectedTest != result.Test {
				t.Fatalf("\n\n"+
					"UT Name:  %s\n"+
					"\033[32mExpected (Test): %s\033[0m\n"+
					"\033[31mActual (Test):   %s\033[0m\n\n",
					scenario.name, scenario.expectedTest, result.Test)
			}

			if scenario.expectedElapsed != result.Elapsed {
				t.Fatalf("\n\n"+
					"UT Name:  %s\n"+
					"\033[32mExpected (Elapsed): %f\033[0m\n"+
					"\033[31mActual (Elapsed):   %f\033[0m\n\n",
					scenario.name, scenario.expectedElapsed, result.Elapsed)
			}
		})
	}
}

// Measure the performance of the ParseJSON function.
func BenchmarkParseJSON(b *testing.B) {
	// DEFINITIONS.
	scenarios := []struct {
		name        string
		inputString string
	}{
		{
			name:        "When the input is an empty string.",
			inputString: "",
		},
		{
			name:        "When the input is NOT a JSON document.",
			inputString: "N.A.",
		},
		{
			name:        "When the input is an empty JSON document.",
			inputString: "{}",
		},
		{
			name:        "When the input is a valid JSON document (Sample 1).",
			inputString: "{\"Action\":\"run\", \"Package\":\"pkg1\", \"Test\":\"tst1\",\"Elapsed\":5}",
		},
		{
			name:        "When the input is a valid JSON document (Sample 2).",
			inputString: "{\"Action\":\"output\", \"Package\":\"pkg2\", \"Test\":\"tst2\",\"Elapsed\":15}",
		},
	}

	// EXECUTION.
	for _, scenario := range scenarios {
		scenario := scenario // NOTE: Ensure that the b.Run function has the correct value when it's being executed.

		b.Run(scenario.name, func(b *testing.B) {
			// RUN.
			for n := 0; n < b.N; n++ {
				// ACT.
				_, _ = gotest.ParseJSONOutput(scenario.inputString)
			}
		})
	}
}

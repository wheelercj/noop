/*
Copyright © 2024 Chris Wheeler <mail@chriswheeler.dev>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"", 0},
		{"0", 0},
		{"00", 0},
		{"1", 1},
		{"9", 9},
		{"11", 2},
		{"33", 6},
		{"45", 9},
		{"54", 9},
		{"123", 6},
		{"398789237438947", 91},
		{"i", 1},
		{"v", 5},
		{"x", 10},
		{"l", 50},
		{"c", 100},
		{"d", 500},
		{"m", 1000},
		{"I", 1},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},
		{"vi", 6},
		{"iv", 6},
		{"iiv", 7},
		{"iiiv", 8},
		{"LVIII", 58},
		{"MCMLXXXIX", 2191},
		{"1i", 2},
		{"5x", 15},
		{"L337", 63},
		{"cC", 200},
		{"1.", -1},
		{".", 0},
		{".2", 2},
		{"1.2", 1},
		{"4.4", 0},
		{"5.3", -2},
		{"5.35", 3},
		{"i.v", 4},
		{"ii.v", 3},
		{"D.3", -497},
		{"29L8.1c829", 51},
		{"X0X0", 20},
		{"0.", 0},
		{".0", 0},
		{"m0c.x1000000000", -1089},
	}

	for _, test := range tests {
		name := fmt.Sprintf("add(%s)", test.input)
		t.Run(name, func(t *testing.T) {
			ans, err := add(test.input)
			if err != nil {
				t.Errorf("%s returned err `%s`, want %d", name, err, test.want)
			} else if ans != test.want {
				t.Errorf("%s = %d, want %d", name, ans, test.want)
			}
		})
	}
}

func TestAddInvalid(t *testing.T) {
	tests := []struct {
		input string
		want  error
	}{
		{"4+5", fmt.Errorf("invalid character: `+`")},
		{"4-5", fmt.Errorf("invalid character: `-`")},
		{"4*5", fmt.Errorf("invalid character: `*`")},
		{"4/5", fmt.Errorf("invalid character: `/`")},
		{"3 + 2", fmt.Errorf("invalid character: ` `")},
		{"3 2", fmt.Errorf("invalid character: ` `")},
		{"387a2", fmt.Errorf("invalid character: `a`")},
		{"1.2.3", fmt.Errorf("multiple periods")},
		{"..", fmt.Errorf("multiple periods")},
		{"..7", fmt.Errorf("multiple periods")},
		{"8..", fmt.Errorf("multiple periods")},
		{".9.", fmt.Errorf("multiple periods")},
		{"c.lm.3", fmt.Errorf("multiple periods")},
		{"c57.m39.X387.", fmt.Errorf("multiple periods")},
	}

	for _, test := range tests {
		name := fmt.Sprintf("add(%s)", test.input)
		t.Run(name, func(t *testing.T) {
			_, err := add(test.input)
			if err.Error() != test.want.Error() {
				t.Errorf("%s returned err:\n\t%s\nwant err\n\t%s", name, err, test.want)
			}
		})
	}
}

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
	"os"

	"github.com/spf13/cobra"
)

var version = "v0.0.4"

var values = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"i": 1,
	"v": 5,
	"x": 10,
	"l": 50,
	"c": 100,
	"d": 500,
	"m": 1000,
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "noop [number]",
	Version: version,
	Short:   "A calculator with no operators",
	Long: `A calculator with no operators

All digits are added together. Only single-character numbers can be
used, including Roman numeral characters. If the input has a period,
the digits to its left are negative.
`,
	Args: cobra.MaximumNArgs(1),
	RunE: runFunc,
}

func runFunc(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		fmt.Println("noop", version)
		return repl()
	} else {
		sum, err := add(args[0])
		if err != nil {
			return err
		}
		fmt.Println(sum)
		return nil
	}
}

// repl is a Read Eval Print Loop. It can be exited with a keyboard interrupt.
func repl() error {
	for {
		fmt.Print(">>> ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			return err
		}
		sum, err := add(input)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(sum)
		}
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}

func add(numStr string) (int, error) {
	var sum int
	var periodFound bool
	for _, ch := range numStr {
		s := string(ch)
		if s == "." {
			if periodFound {
				return 0, fmt.Errorf("multiple periods")
			}
			sum = -sum
			periodFound = true
		} else if v, ok := values[s]; ok {
			sum += v
		} else {
			return 0, fmt.Errorf("invalid character: `%s`", s)
		}
	}

	return sum, nil
}

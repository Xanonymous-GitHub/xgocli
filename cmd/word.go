package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
	"xgocli/internal/word"
)

type Mode string

const (
	ModeUpper                      Mode = "b"  // convert all word to uppercase.
	ModeLower                      Mode = "s"  // convert all word to lowercase.
	ModeUnderscoreToUpperCamelCase Mode = "ub" // convert to upper camel case.
	ModeUnderscoreTpLowerCamelCase Mode = "us" // convert to lower camel case.
	ModeCamelCaseToUnderscore      Mode = "u"  // convert to underscore.
	Undefined                      Mode = "x"
)

var doc = strings.Join([]string{
	"Word conversion supports various word format conversion.\n",
	"\tb:  Convert all characters to uppercase.",
	"\ts:  Convert all characters to lowercase.",
	"\tub: Convert underscore words to uppercase camel case words.",
	"\tus: Convert underscore words to lowercase camel case words.",
	"\tu:  Convert camel case words to underscore words.",
}, "\n")

var (
	data    string
	mode    Mode
	wordCmd = &cobra.Command{
		Use:   "word",
		Short: "word type conversion",
		Long:  doc,
		Run: func(cmd *cobra.Command, args []string) {
			var result string

			switch mode {
			case ModeUpper:
				result = word.ToUpper(data)
			case ModeLower:
				result = word.ToLower(data)
			case ModeUnderscoreToUpperCamelCase:
				result = word.UnderscoreToUpperCamelCase(data)
			case ModeUnderscoreTpLowerCamelCase:
				result = word.UnderscoreToLowerCamelCase(data)
			case ModeCamelCaseToUnderscore:
				result = word.CamelCaseToUnderScore(data)
			case Undefined:
				log.Fatalf("please choose a conversion mode.")
			default:
				log.Fatalf("unsupported conversion mode.")
			}

			fmt.Println(result)
		},
	}
)

func init() {
	wordCmd.Flags().StringVarP(&data, "str", "s", "", "some strings")
	wordCmd.Flags().StringVarP((*string)(&mode), "mode", "m", "x", "conversion mode")
}

# xgoCLI
The collection of some useful CLI tools.

### Build to use by yourself!
`go build main.go`

## Word converter
In daily life and work, we often see some strings named after a single word, and then convert them into names in various formats. For example, we have defined a name originally, but it may need to be converted into one or more constants. At this time, if you manually modify one by one, it will be too cumbersome. Not only is it possible to correct mistakes, but also the work efficiency is not good.

This program contains the following features:
- Convert all words to lowercase
- Convert all words to uppercase
- Convert bottom-line words to uppercase camel case words
- Underline words are converted to lowercase camel case words
- Convert hump words to underline words

```
  Word conversion supports various word format conversion.

	b:  Convert all characters to uppercase.
	s:  Convert all characters to lowercase.
	ub: Convert underscore words to uppercase camel case words.
	us: Convert underscore words to lowercase camel case words.
	u:  Convert camel case words to underscore words.

  Usage:
     word [flags]

  Flags:
    -h, --help          help for word
    -m, --mode string   conversion mode (default "x")
    -s, --str string    some strings
```

## Convenient time tool
When viewing the original data, it is often necessary to look at the personalized time after formatting, or the time stamp, etc. If the time format in different systems is different, and the comparison rules are different, then a round of conversion is required every time it is used. In many cases, the start time and end time of the import parameters of the business interface are values of a timestamp. At this time, it is necessary to rely on some external fast websites or internal web sites to obtain and adjust the time. First, you must connect to the Internet, and then enter the URL, etc., which is obviously not in line with efficiency thinking.

```
Time format conversion

Usage:
   time [flags]
   time [command]

Available Commands:
  calc        calculate time in need
  now         Get current time

Flags:
  -h, --help   help for time
```

## sql scripts to golang structure conversion
When initializing a project or adding a new data table, we often need to increase the model structure. At this time, we will encounter a new problem, that is, we need to write the model structure. If it is a handwritten model structure, it is too inefficient, so this program implements the conversion of the database table to the Golang structure.

Only MySQL is currently supported. 0.0

```
sql conversion and processes

Usage:
   sql [flags]
   sql [command]

Available Commands:
  struct      sql transfer

Flags:
  -h, --help   help for sql
```

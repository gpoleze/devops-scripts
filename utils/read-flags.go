package utils

import (
	"errors"
	"flag"
	"fmt"
	"slices"
)

type MyFlag struct {
	Name         string
	ShortName    string
	Description  string
	Value        string
	Required     bool
	DefaultValue string
}

type MyFlags []MyFlag

func (mfs MyFlags) UpdateDefaultValue(name, value string) {
	index := slices.IndexFunc(mfs, func(f MyFlag) bool { return f.Name == name })
	mf := mfs[index]
	mf.DefaultValue = value
	mfs[index] = mf
}

type ParsedFlags map[string]*string

var AwsFlags = []MyFlag{
	{
		Name:        "region",
		ShortName:   "r",
		Description: "AWS region",
	},
	{
		Name:        "profile",
		ShortName:   "p",
		Description: "AWS profile",
	},
	{
		Name:        "output-type",
		ShortName:   "o",
		Description: "Output type (table or json)",
	},
}

func ReadAwsFlags() (*string, *string, *string) {
	ReadFlags(AwsFlags)
	return &AwsFlags[0].Value, &AwsFlags[1].Value, &AwsFlags[2].Value
}

func ReadFlags(flags []MyFlag) (ParsedFlags, error) {
	parsedFlags := make(map[string]*string)
	for i := range flags {
		flag.StringVar(&(flags[i]).Value, flags[i].Name, flags[i].DefaultValue, flags[i].Description)
		if flags[i].ShortName != "" {
			flag.StringVar(&(flags[i]).Value, flags[i].ShortName, flags[i].DefaultValue, flags[i].Description+"(shorthand)")
		}
		parsedFlags[flags[i].Name] = &flags[i].Value
	}
	flag.Parse()

	var errorMessage string
	for _, myFlag := range flags {
		if myFlag.Required && myFlag.Value == "" {
			errorMessage += fmt.Sprintf("Error: %s is required\n", myFlag.Name)
		}
	}
	if errorMessage != "" {
		return ParsedFlags{}, errors.New(errorMessage)
	}

	return parsedFlags, nil
}

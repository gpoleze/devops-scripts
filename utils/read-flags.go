package utils

import "flag"

func ReadAwsFlags() (*string, *string, *string) {
	var region string
	var profile string
	var outputType string
	flag.StringVar(&region, "region", "", "AWS region")
	flag.StringVar(&region, "r", "", "AWS region (shorthand)")

	flag.StringVar(&profile, "profile", "", "AWS profile")
	flag.StringVar(&profile, "p", "", "AWS profile (shorthand)")

	flag.StringVar(&outputType, "o", "table", "output type type(shorthand)")
	flag.StringVar(&outputType, "output", "table", "output type (accepted types table and json")

	flag.Parse()
	return &region, &profile, &outputType
}

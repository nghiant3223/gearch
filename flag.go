package main

var fResultLimit = IntFlag{
	DefaultValue: 3,
	Flag: Flag{
		Name:        "c",
		Description: "Maximum number of results showing",
	},
}

package main

import "github.com/yoheimuta/protolint/plugin"

func main() {
	plugin.RegisterCustomRules(
		&RequiredOptionRule{
			requiredOptions: []string{"go_package"},
		},
	)
}

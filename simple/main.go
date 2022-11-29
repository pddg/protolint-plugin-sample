package main

import "github.com/yoheimuta/protolint/plugin"

func main() {
	plugin.RegisterCustomRules(
		&SimpleDenyRule{},
	)
}

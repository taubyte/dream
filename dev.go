//go:build dreamdev
// +build dreamdev

package main

import service "github.com/taubyte/dreamland/service"

func init() {
	service.Dev = true
}

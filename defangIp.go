package main

import "strings"

func defangIp(ip string) string {
  return strings.ReplaceAll(ip, ".", "[.]")
}

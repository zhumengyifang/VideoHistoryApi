package main

import "gindemo/api"

type Host struct {
	IP   string
	Name string
}

func main() {
	api.ApiInit()
}

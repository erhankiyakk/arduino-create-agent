// Code generated by goa v3.13.2, DO NOT EDIT.
//
// HTTP request path constructors for the tools service.
//
// Command:
// $ goa gen github.com/arduino/arduino-create-agent/design

package server

import (
	"fmt"
)

// AvailableToolsPath returns the URL path to the tools service available HTTP endpoint.
func AvailableToolsPath() string {
	return "/v2/pkgs/tools/available"
}

// InstalledToolsPath returns the URL path to the tools service installed HTTP endpoint.
func InstalledToolsPath() string {
	return "/v2/pkgs/tools/installed"
}

// InstallToolsPath returns the URL path to the tools service install HTTP endpoint.
func InstallToolsPath() string {
	return "/v2/pkgs/tools/installed"
}

// RemoveToolsPath returns the URL path to the tools service remove HTTP endpoint.
func RemoveToolsPath(packager string, name string, version string) string {
	return fmt.Sprintf("/v2/pkgs/tools/installed/%v/%v/%v", packager, name, version)
}

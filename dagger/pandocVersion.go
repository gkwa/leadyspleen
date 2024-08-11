package main

import "fmt"

// PandocVersion is a custom type to handle version logic
type PandocVersion string

// NewPandocVersion creates a new PandocVersion with a default if empty
func NewPandocVersion(version string) PandocVersion {
	if version == "" {
		return PandocVersion("latest")
	}
	return PandocVersion(version)
}

// String returns the string representation of the version
func (pv PandocVersion) String() string {
	return string(pv)
}

// ImageName returns the full image name for this version
func (pv PandocVersion) ImageName() string {
	return fmt.Sprintf("pandoc/core:%s", pv)
}

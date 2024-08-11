package main

import "fmt"

type PandocVersion string

func NewPandocVersion(version string) PandocVersion {
	if version == "" {
		return PandocVersion("latest")
	}
	return PandocVersion(version)
}

func (pv PandocVersion) String() string {
	return string(pv)
}

func (pv PandocVersion) ImageName() string {
	return fmt.Sprintf("pandoc/core:%s", pv)
}

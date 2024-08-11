package main

import (
	"context"
	"fmt"

	"dagger/leadyspleen/internal/dagger"
)

type Leadyspleen struct{}

func getImageName(pandocVersion string) string {
	return fmt.Sprintf("pandoc/core:%s", pandocVersion)
}

func (m *Leadyspleen) PandocRun(
	ctx context.Context,
	// +optional
	pandocVersion string,
	directoryArg *dagger.Directory,
) (string, error) {
	if pandocVersion == "" {
		pandocVersion = "latest"
	}
	image := getImageName(pandocVersion)
	return dag.Container().
		From(image).
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"pandoc", "--to", "native", "test.md"}).
		Stdout(ctx)
}

func (m *Leadyspleen) PandocVersion(
	ctx context.Context,
	// +optional
	pandocVersion string,
) (string, error) {
	if pandocVersion == "" {
		pandocVersion = "latest"
	}
	image := getImageName(pandocVersion)
	return dag.Container().
		From(image).
		WithExec([]string{"pandoc", "-v"}).
		Stdout(ctx)
}

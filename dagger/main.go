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

func (m *Leadyspleen) PandocRun(ctx context.Context, pandocVersion string, directoryArg *dagger.Directory) (string, error) {
	image := getImageName(pandocVersion)
	return dag.Container().
		From(image).
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"pandoc", "--to", "native", "test.md"}).
		Stdout(ctx)
}

func (m *Leadyspleen) PandocVersion(ctx context.Context, pandocVersion string) (string, error) {
	image := getImageName(pandocVersion)
	return dag.Container().
		From(image).
		WithExec([]string{"pandoc", "-v"}).
		Stdout(ctx)
}

package main

import (
	"context"

	"dagger/leadyspleen/internal/dagger"
)

type Leadyspleen struct{}

func (m *Leadyspleen) PandocRun(
	ctx context.Context,
	directoryArg *dagger.Directory,
	// +optional
	pandocVersion string,
) (string, error) {
	pv := NewPandocVersion(pandocVersion)

	return dag.Container().
		From(pv.ImageName()).
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
	pv := NewPandocVersion(pandocVersion)
	return dag.Container().
		From(pv.ImageName()).
		WithExec([]string{"pandoc", "-v"}).
		Stdout(ctx)
}

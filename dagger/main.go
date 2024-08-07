package main

import (
	"context"

	"dagger/leadyspleen/internal/dagger"
)

type Leadyspleen struct{}

func (m *Leadyspleen) PandocRun(ctx context.Context, directoryArg *dagger.Directory) (string, error) {
	return dag.Container().
		From("pandoc/core:3.2.1").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"pandoc", "-t", "native", "test.md"}).
		Stdout(ctx)
}

func (m *Leadyspleen) PandocVersion(ctx context.Context) (string, error) {
	return dag.Container().
		From("pandoc/core:3.2.1").
		WithExec([]string{"pandoc", "-v"}).
		Stdout(ctx)
}

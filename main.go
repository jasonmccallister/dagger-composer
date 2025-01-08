// A generated module for Composer functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/dagger-composer/internal/dagger"
	"errors"
	"fmt"
)

type Composer struct {
	Version string
	Dir     *dagger.Directory
}

func New(
	// +optional
	// +default="latest"
	version string,
	// +optional
	// +defaultPath="."
	dir *dagger.Directory,
) *Composer {
	return &Composer{
		Version: version,
		Dir:     dir,
	}
}

// Takes a directory and installs composer dependencies.
func (m *Composer) Install(ctx context.Context) (*dagger.Directory, error) {
	// make sure there is a composer.json file
	if !exists(ctx, m.Dir, "composer.json") {
		return nil, errors.New("missing composer.json file in path")
	}

	// if there is no composer.lock, log it but continue
	if !exists(ctx, m.Dir, "composer.lock") {
		// TODO(jasonmccallister) log to stdout in the dagger way
		fmt.Println("composer.lock file not found in path")
	}

	return dag.Container().
		From("composer:"+m.Version).
		WithMountedDirectory("/app", m.Dir).
		WithWorkdir("/app").
		WithExec([]string{"composer", "install", "--prefer-dist", "--optimize-autoloader", "--ignore-platform-reqs"}).
		Directory("/app/vendor"), nil
}

// can be simplified when this is completed: https://github.com/dagger/dagger/issues/6713
func exists(ctx context.Context, src *dagger.Directory, path string) bool {
	if _, err := src.File(path).Sync(ctx); err != nil {
		return false
	}

	return true
}

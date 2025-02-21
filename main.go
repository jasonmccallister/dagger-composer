// Dagger Composer Module
//
// This module provides a way to install composer dependencies in a directory for a
// PHP project. It uses the composer image to install the dependencies and returns the
// vendor directory with the dependencies installed.

package main

import (
	"context"
	"dagger/dagger-composer/internal/dagger"
	"errors"
	"fmt"
)

type Composer struct {
	Version string            // +private
	Source  *dagger.Directory // +private
}

func New(
	// +optional
	// +default="latest"
	// The version of the composer image to use.
	version string,
	// +optional
	// +defaultPath="."
	// The directory that contains the composer.json/composer.lock files.
	source *dagger.Directory,
) *Composer {
	return &Composer{
		Version: version,
		Source:  source,
	}
}

// Takes a directory and installs composer dependencies.
func (m *Composer) Install(
	ctx context.Context,
	// +optional
	// +default=["--prefer-dist", "--optimize-autoloader", "--ignore-platform-reqs"]
	// Additional arguments to pass to the composer install command.
	args []string) (*dagger.Directory, error) {
	// make sure there is a composer.json file
	if !exists(ctx, m.Source, "composer.json") {
		return nil, errors.New("missing composer.json file in path")
	}

	// if there is no composer.lock, log it but continue
	if !exists(ctx, m.Source, "composer.lock") {
		// TODO(jasonmccallister) log to stdout in the dagger way
		fmt.Println("composer.lock file not found in path")
	}

	exec := append([]string{"composer", "install"}, args...)

	return dag.Container().
		From("composer:"+m.Version).
		WithMountedDirectory("/app", m.Source).
		WithWorkdir("/app").
		WithExec([]string{"mkdir", "/composer"}).
		WithEnvVariable("COMPOSER_HOME", "/composer").
		WithExec(exec).
		Directory("/app/vendor"), nil
}

// can be simplified when this is completed: https://github.com/dagger/dagger/issues/6713
func exists(ctx context.Context, src *dagger.Directory, path string) bool {
	if _, err := src.File(path).Sync(ctx); err != nil {
		return false
	}

	return true
}

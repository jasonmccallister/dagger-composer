# Composer Dagger Module

This is a [Dagger](https://dagger.io) module to handling installing dependencies in a PHP project with Composer.

## Usage

To manage dependencies in the current working directory:

```
dagger call -m github.com/jasonmccallister/dagger-composer install
```

To provide a specific directory (default is the current directory) that contains the `composer.lock` and `composer.json` files:

```
dagger call -m github.com/jasonmccallister/dagger-composer --dir ./path/to/directory install
```

To use a specific version of composer (default is `latest`) use the `--version` flag.

```
dagger call -m github.com/jasonmccallister/dagger-composer --version 2.2 install
```

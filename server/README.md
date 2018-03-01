# go-starter

This repository is a barebones golang web service scaffold that sets up routes, SSL, static file serving, and a global pooled mysql database connection. [Glide](https://github.com/Masterminds/glide) is used for dependency management. This is still in progress!

## Initial Installation
Since Golang uses the directory structure of your project for import/build paths, to use this code all import instances with "github.com/suyashkumar/go-starter" will have to be replaced with your import path. A script is provided to automatically do this for you. 
1. `git clone https://github.com/suyashkumar/go-starter <path-of-your-project>`
2. `cd <path-of-your-project>`
3. `./install <go-import-path-of-your-project>` You must supply the go import path for your project. The current path is simply "github.com/suyashkumar/go-starter" but yours might be "github.com/bob/hello-world" if your project lives in `$GOPATH/src/github.com/bob/hello-world`. 
4. `glide install` to install dependencies.
5. (Optional) you may want to `rm -rf .git` and `git init` your new git repo from scratch. 

## General Usage
* `make` to run tests and build your project
* `make release` to build win/linux/darwin binaries of your project
* `make test` to just run all tests in the project
* Docker configuration coming soon!

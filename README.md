# Clean Architecture

<img src="/ref/logo.png" align="right" width="210">

> Implementation of clean architecture based on Uncle Bob's book

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/waffle-frame/clean-architecture-template)](https://github.com/gomods/athens) ![Swagger Validator](https://img.shields.io/swagger/valid/3.0?specUrl=https://raw.githubusercontent.com/waffle-frame/clean-architecture-template/master/pkg/docs/swagger.yaml) [![wakatime](https://wakatime.com/badge/user/ec493241-c2a0-40a9-8ff1-637bdb54b2f1/project/d9c08ded-4928-49a3-8921-49aa3087f700.svg)](https://wakatime.com/badge/user/ec493241-c2a0-40a9-8ff1-637bdb54b2f1/project/d9c08ded-4928-49a3-8921-49aa3087f700)

Clean Architecture was proposed by Robert Martin (also known as Uncle Bob) in his book "Clean Architecture: A Craftsman's Guide to Software Structure and Design" published in 2017. The goal of Clean Architecture is to create a flexible, extensible and easily maintainable software architecture.

The main idea of Clean Architecture is to divide the program into layers, where each layer has its own responsibility and depends only on the layers below. This makes it easy to change or replace components within each layer without affecting the rest of the system.

---

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=10 orderedList=false} -->

<!-- code_chunk_output -->

- [Clean Architecture](#clean-architecture)
  - [Project Navigation](#project-navigation)
  - [Installation](#installation)
    - [Step 1. Check dependencies](#step-1-check-dependencies)
    - [Step 2. Setting up the environment](#step-2-setting-up-the-environment)
      - [2.1 Database creation](#21-database-creation)
      - [2.2 Сonfiguration file setup](#22-сonfiguration-file-setup)
      - [2.3 Swagger setup (Option)](#23-swagger-setup-option)
    - [Step 3. Run](#step-3-run)
  - [TODO](#todo)

<!-- /code_chunk_output -->

## Project Navigation

| Name            | Link                      |
| --------------- | ------------------------- |
| Database schema | <http://link.com>        |
| Design          | <http://figma.com>       |
| Technical task  | <http://docs.google.com> |

## Installation

### Step 1. Check dependencies

Clone project from GitHub:

```bash
git clone https://gitlab.humo.tj/orzu/applications/bridge.git
cd clean-architecture-template
```

Before running or building a project, make sure that all the necessary dependencies are installed on your OS.
This will also install the dependencies required by the golong application and install them if necessary.

To find out information about it, you can use the command:

```bash
make check-dependencies
```

**Note**: To add one or more applications to the list of dependencies, you will need to edit the `DEPENDENCIES` variable in the `Makefile`

Example:

```makefile
# Project dependencies
DEPENDENCIES = Golang="go version" \
                Keycloak="docker exec 1ca7b3a28abb cat /opt/keycloak/version.txt" \
                ...
```

In this example I accessed the docker image and got the keycloak version

### Step 2. Setting up the environment

#### 2.1 Database creation

```bash
sudo -u postgres psql -c 'CREATE DATABASE clean_arch;'
```

#### 2.2 Сonfiguration file setup

An instance of the config file is located along the [path](ref/example-config.json).
You can supplement it with your own data. Also, in order for you to be able to use the configurations in the application, you need to add an identical structure in [config.go](pkg/config/module.go) file.

If successful, the `ImportConfigs` function will return your entry in the content.

**Note**: If you don't want to populate the config file, you can use the instance using the following command:

```bash
cp ref/example-config.json config.json
```

#### 2.3 Swagger setup (Option)

For the swagger to work properly, you need to replace the `{{}}` constructs in the [swagger.go](pkg/docs/swagger.go) file.
The names of the variables speak for themselves, so there is no need to explain the meaning of the variables.

To apply the changes, you need to run the command:

```bash
make docs
```

### Step 3. Run

To run in development mode, use the command:

```bash
make run
```

To build the application, use the command:

```bash
make
# or
make build
```

> The assembly contains a dependency check, the generation of a swagger file and the assembly itself
> The binary file will be located on the path ./build/app

---

## TODO

- [ ] Implement CI/CD
- [ ] Add information to other readme files
- [ ] Сreate a project that will have an example of using this blog-based template

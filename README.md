# SpyCode CLI

**SpyCLI** is a CLI library for work with iac
projects, blueprints, modules, etc

### Usage:
`spycli [command]`

### Available Commands:
- `blueprint`   Manipulate iac blueprints
- `completion`  Generate the autocompletion script for the specified shell
- `help`        Help about any command
- `module`      Manipulate modules
- `project`     Manipulate iac projects

#### Flags:
- `-h`, `--help`      help for spycli
- `-V`, `--verbose`   verbose output

Use "spycli [command] --help" for more information about a command.

## Project
Use project commands

### Usage:
`spycli project [command]`

### Available Commands:
- `init`        Initialize a project
- `new`         Create new project

#### Flags:
- `h`, --help   help for project

## Project new
`new`: creates a new project

Ex:

`spycli project new -n "Prj Simple Web App" -b /src/bp-aws-nearform -s simple-web-app -l "/src/tf-modules-aws" -r us-east-1 -e dev -e prd`

### Usage:
`spycli project new [flags]`

#### Flags:
- `-b`, `--blueprint` string           Blueprint
- `-v`, `--blueprint-version` string   Blueprint version
- `-d`, `--directory` string           Base directory where the files will be writen
- `-e`, `--environment` strings        Pass a list of environments (default [dev])
- `-h`, `--help`                       help for new
- `-l`, `--library` string             Library (ex: git@github.com:spycode-io/tf-components.git
- `-n`, `--name` string                Element name (ex: my-project or my-blueprint)
- `-p`, `--platform` string            Plataform or service (aws|azure) (default "aws")
- `-r`, `--region` strings             Pass a list of environments (default [us-east-1])
- `-s`, `--stack` string               Stack name
- `-k`, `--version` string             Library version (or tag if it's a git repository)

Global Flags:
- `-V`, `--verbose`   verbose output

## Project init
Use project init on a project folder to sync blueprint files

Ex:

`spycli project init`

### Usage:

`spycli project init [flags]`

### Flags:
- `-d`, `--directory` string   Base directory where the files will be writen (default ".")
- `-h`, `--help`               help for init
- `-l`, `--link`               Base directory where the files will be writen

Global Flags:
- `-V`, `--verbose`   verbose output

## Blueprint
Use project new

### Usage:
`spycli blueprint [command]`

Available Commands:
- `new`         Create new project

### Flags:
- `-h`, `--help`   help for blueprint

### Global Flags:
- `-V`, `--verbose`   verbose output

Use `spycli blueprint [command] --help` for more information about a command.

## Blueprint new
Use blueprint commands

`new`: creates a new blueprint

Ex:

`spycli blueprint new -n "BP AWS Nearform" -s simple-web-app -b "git@github.com:spycode-io/bp-test.git" -r us-east-1`

### Usage:
  spycli blueprint new [flags]

### Flags:
 - `-b`, `--blueprint` string   Blueprint
 - `-d`, `--directory` string   Base directory where the files will be writen
 - `-h`, `--help`               help for new
 - `-n`, `--name` string        Element name (ex: my-project or my-blueprint)
 - `-r`, `--region` strings     Pass a list of regions
 - `-s`, `--stack` string       Stack name
 - `-v`, `--version` string     Blueprint version (default "v0.0.0")

### Global Flags:
- `-V`, `--verbose`   verbose output

## Module
Use module commands

### Usage:
`spycli module [command]`

### Available Commands:
- `include`     Include module

### Flags:
- `-h`, `--help`   help for module

Global Flags:
- `-V`, `--verbose`   verbose output

Use `spycli module [command] --help` for more information about a command.

## Module include
`include`: includes a new module in a blueprint region

To create a vpc module called web-app-vpc:

`spycli module new -m vpc -n "VPC Web App"`

### Usage:
- `spycli module include [flags]`

### Flags:
- `-d`, `--directory` string   Base directory where the files will be writen
- `-h`, `--help`               help for include
- `-m`, `--module` string      Module (ex: aws/vpc)
- `-n`, `--name` string        Element name (ex: my-project or my-blueprint)

### Global Flags:
- `-V`, `--verbose`   verbose output
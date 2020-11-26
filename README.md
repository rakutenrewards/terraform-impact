# terraform-impact :boom:
Terraform Impact is a simple tool and is meant to stay that way. It does its static analysis in 4 basic steps:

1. Lists all terraform states
2. For each terraform state, creates a file and module dependency tree.
3. Lists states impacted by any of the input files
4. Outputs impacted states

## Usage
```
Terraform Impact.

This tool takes a list of files as input and outputs a list of all the Terraform states
impacted by any of those files. An impact is described as a file creation, modification
or deletion.

Usage:
  impact [--pattern <string>] [--rootdir <dir>] <files>...
  impact -h | --help
  impact -v | --version

Options:
  -r --rootdir <dir>     The directory from where the state discovery begins [default: .]
  -p --pattern <string>  A string to filter states. Only states whose path contains the
                         string will be taken into account. [default: ]
  -h --help              Show this screen.
  -v --version           Show version.
```

## Examples
Some call examples for the `test_resources` repository.
```
impact \
  --rootdir ./test_resources/terraform \
  --pattern /gcp/ \
  test_resources/terraform/gcp/modules/unused_module/outputs.tf \
  test_resources/terraform/gcp/modules/google/runtime_config/variables.tf
```


## Steps
Details for each steps

### 1. State listing
#### Discovery
The tool recursively looks for all directories in the file system from the provided `rootdir`.

To decide whether a directory `d` is a state or not, the tool looks if there's the following blocks in the `d/main.tf` file.

```
terraform {
    backend {
        ...
    }
}
```
The search can be further filtered by passing a `string pattern`. This makes the tool ignore all directories not containing the provided pattern in their path.

### 2. Dependency tree
For each state, the tool recursively looks for `module` blocks and builds a dependency tree where each node contains the `path` to the module and a list of `nodes` as dependencies.

```
// in trees/node.go
type Node struct {
	Path         string
	Dependencies []*Node
}
```
Important note, the dependency tree builder follows file symlinks to add them in the dependencies.

### 3. List impacted states
Filters states list by looking if any of the input files are in the state dependency tree.

### 4. Outputs impacted states
Takes the result from (`3. List impacted states`) and outputs it in the terminal `stdout`.

## Contributing
### Tests
In `test-resources`, you'll find a wannabe Terraform project which is used for tests.

In order to make the tests run from the root of the repository, the following needs to imported
```
_ "github.com/RakutenReady/terraform-impact/testutils/setup"
```

### Makefile
Every useful commands are in the `Makefile`. Here's an explicit list:
```
make clean
make format
make unit-tests
make integration-tests
```

### Command-line library
`Docopt` is the chosen one. See in [Github](https://github.com/docopt/docopt.go)
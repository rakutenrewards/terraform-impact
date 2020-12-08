# terraform-impact :boom:
Terraform Impact is a simple tool and is meant to stay that way. It does its static analysis in 4 basic steps:

1. Lists all terraform states
2. For each terraform state, creates a file and module dependency tree.
3. Lists states impacted by any of the input files
4. Outputs impacted states

## Usage
```
./terraform-impact -h
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
The search can be further filtered by passing a `regexp pattern`. This makes the tool ignore all directories not matching the provided pattern in their path.

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
Takes the result from (`3. List impacted states`) and outputs it in the terminal or in a file.

## Contributing
### Tests
In `test-resources`, you'll find a wannabe Terraform project which is used for tests.

In order to make the tests run from the root of the repository, the following needs to imported
```
_ "github.com/RakutenReady/terraform-impact/testutils/setup"
```

### Github tests
See [README.md](https://github.com/RakutenReady/terraform-impact/tree/github-test-main)

To run these tests, you'll need to setup the following env vars.
```
GITHUB_USERNAME=<your Github username>
GITHUB_PASSWORD=<a generated Github token>
```

Those tests are located in `e2etests` meaning it only affects `make integration-tests`

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
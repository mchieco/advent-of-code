# Advent of Code 2023

A lightweight execution environment for the [advent of code](https://adventofcode.com/) challenges written in Go.

### Test the solutions

Before submitting a solution, you can test the algorithms with the template input. Navigate to the day you want to check and run the
following command:

```sh
go test
```

Alternatively, you can also run the command from the project root.

```sh
go test ./year/day_xx/
```

### Input Files

The input files are located in the `inputs` folder, which are gitignored. The file names are in the following format: `day_xx.txt`. The input files are not included in the repository, you have to download them from the [advent of code](https://adventofcode.com/) website, and are generated specific to the user.

### Compile and run

You can compile code and run with the actual input if all the tests pass. For this, first, run this command:

```sh
go build .
```

To run the solution, you need to provide a few extra arguments.
* the `--year` flag must be set to specify which year's solution should run
* the `--day` flag must be set to specify which day's solution should run
* the `--input` flag specifies the path of the file containing the actual input
* optionally, you can add the `--mode` flag to run one part of the daily challenge, accepted values are 1 and 2

And now the complete command:

```sh
go run . --year y --day x --input path_to_input --mode 1
```
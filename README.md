# 🐹 go-tests

[![🐳 On Master - Build and Test Docker Image](https://github.com/01-edu/go-tests/actions/workflows/ga-image-build-master.yml/badge.svg?branch=master)](https://github.com/01-edu/go-tests/actions/workflows/ga-image-build-master.yml)

Private repository that holds the files needed to build the Go tests Docker image.

To run the tests make sure the two repositories are right next to each other:

- https://github.com/01-edu/piscine-go
- https://github.com/01-edu/go-tests

To test an exercise, run this command in the `go-tests` folder:

```
./test_one.sh isnegative
```

To run all the exercises, run this command in the `go-tests` folder:

```
./test_all.sh
```

### Run an exercise in the Docker container

First you need to build the image, for example to build it locally you may run:

```
docker build -t go_tests .
```

To run an exercise as close as it can get to production environment you can call `test_with_docker.sh` as follows:

```
./test_with_docker.sh [EXERCISE_NAME] [EXERCISE_PATH] [ALLOWED_FUNCTION]*
```

As an example:

```console
./test_with_docker getalpha getalpha/main.go --allow-builtin github.com/01-edu/z01.PrintRune strconv.Atoi os.Args
```

For convenience to do bulk tests you may want to store this info in a `meta.txt` file that will separate them with `:` and then use `awk` to run tests on all of them.

As an example:

```console
# The meta.txt file
getalpha:getalpha/main.go:--allow-builtin github.com/01-edu/z01.PrintRune strconv.Atoi os.Args

# The command to run
cat meta.txt | awk -F ":" '{print system("./test_with_docker.sh "$1" "$2" "$3)}'
```

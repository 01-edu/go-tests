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

# go-tests

To run the tests make sure the two repositories are right next to each other:

- github.com/01-edu/piscine-go
- github.com/01-edu/go-tests

To test an exercise, run this command in the `go-tests`:

```
go run github.com/01-edu/go-tests/tests/isnegative_test
```

Relative paths work:

```
go run ./tests/printalphabet_test
cd tests
go run ./isnegative_test
```

No output means success.

`kallax` does not support module yet. Must be run from GOPATH.

Generate kallax support code:

```
kallax gen
```

Generate migrations:

```
kallax migrate --input . --out migrations
```

Run all migrations:

```
kallax migrate up --dsn 'localhost/kallax_test?sslmode=disable' --all
```

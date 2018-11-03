# Build & Run

```
go get -u github.com/hayeah/go-kallax-bug
go-kallax-bug
```

Output would be something like:

```
// ...

obj.Data: abff8e1e90dd27df343ebef523e241326bd6a1490f960d07ee3095df7f6850e0
obj.Data2: c83aa6f5aa1f489df8ccaa55eac71eb293d27637be03a55d777fefcd24796f43

// ...

after update obj.Data: 0001310000000135000045000000090000000000530000000431205748455245
after update obj.Data2: 2432000000440000000753330053000000047637be03a55d777fefcd24796f43
```

We expect `Data` and `Data2` to be the same after update (since only `Counter` is updated).

# Create Database

```
CREATE database kallax_test;
```

Run all migrations:

```
kallax migrate up --dsn 'localhost/kallax_test?sslmode=disable' --all
```

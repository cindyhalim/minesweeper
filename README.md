<div align="center">

# minesweeper

simple minesweeper TUI written in Go

</div>

### get started:

1. run `cd cmd/minesweeper`
2. access the minesweeper TUI by running the program:

```
go run main.go <mode>

```

3. or creating and running the executable:

```
go build
./minesweeper <mode> for mac
minesweeper.exe <mode> for windows
```

### modes:

- `beginner`: 9 x 9 board, 10 mines
- `intermediate`: 16 x 16 board, 40 mines
- `expert`: 16 x 30 board, 99 mines

<div align="center">

# minesweeper

simple minesweeper TUI written in Go

powered by [bubbletea](https://github.com/charmbracelet/bubbletea) 🧋


https://user-images.githubusercontent.com/52426595/234143315-84b4e0b4-0635-40c2-8c53-98da5bb548f3.mov


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

### improvements:
- select mode UI
- keep track of high score for each mode

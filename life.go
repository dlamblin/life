package main

import (
  "fmt"
  "time"
)

const (
  rows          = 32
  cols          = 48
  cellsPerValue = 16 // a uint64 contains 16 nibbles
  deadCell      = `⋅ `
  liveCell      = `○ `
  vt100         = "\033["
  vtUp          = `A`
  interframe    = 200 * time.Millisecond
)

func main() {
  fmt.Printf(fmt.Sprintf("%% %ds\n", (cols+2)), "Life")
  colz := cols / cellsPerValue
  board := make([][]uint64, rows)
  boardone := make([]uint64, rows*colz)
  for i := range board {
    board[i], boardone = boardone[:colz], boardone[colz:]
  }
  start_test := []string{
    "                               x",
    "  x            x                ",
    "  x           x                 ",
    "  x                             ",
    "                x               ",
    "",
    "",
    "  x",
    "  x",
    "  x",
    "",
    "",
    "  x   xx",
    "  x   xx",
    "  x",
    "",
    "",
    "               xx",
    "               xx",
    "                 xx",
    "                 xx",
    "",
    "",
    "",
    "      xxx",
    "      x x",
    "      xxx",
    "",
    "",
    "",
    "                               x",
    "                               x",
  }
  start_gosper := []string{
    "                                      ",
    "                                      ",
    "                                      ",
    "                                      ",
    "                                      ",
    "                                      ",
    "                                      ",
    "                                      ",
    "                                      ",
    "                                      ",
    "                                      ",
    "                         1            ",
    "                       1 1            ",
    "             11      11            11 ",
    "            1   1    11            11 ",
    " 11        1     1   11               ",
    " 11        1   1 11    1 1            ",
    "           1     1       1            ",
    "            1   1                     ",
    "             11                       ",
    "                                      ",
    "                                      ",
    "                                      ",
    "                                      ",
    "                                      ",
    "                                      ",
    "                                      ",
    "                                      ",
    "                                      ",
    "                                      ",
    "                                      ",
    "                                      ",
  }

  _ = start_test
  start := start_gosper

  for i := range start {
    fmt.Println()
    for j, k := range start[i] {
      if k != ' ' {
        set(board, j, i)
      }
    }
  }

  for f := 0; ; f++ {
    out(board, f)
    count(board)
    rules(board)
  }
}

func set(board [][]uint64, x, y int) {
  if x < cols && y < rows {
    j := uint(x / cellsPerValue)
    offset := x % cellsPerValue
    mask := uint64(1) << uint((15-offset)*4)
    board[y][j] |= mask
  }
}

func out(board [][]uint64, f int) {
  fmt.Print(vt100, rows+1, vtUp, "frame:", f, "\n")
  for i := range board {
    for j := range board[i] {
      o := uint64(1) << uint(60)
      for o > 0 {
        if (board[i][j] & o) == o {
          fmt.Print(liveCell)
        } else {
          fmt.Print(deadCell)
        }
        o >>= 4
      }
    }
    fmt.Println()
  }
}

func count(board [][]uint64) {
  for i := range board {
    for j := range board[i] {
      // leftmost bit of uint64 adds to its 3 left and up and down
      m := uint64(1) << 60
      v := board[i][j] & m
      var y int
      if v > 0 {
        if j == 0 {
          y = len(board[i]) - 1
        } else {
          y = j - 1
        }
        if i == 0 {
          board[len(board)-1][y] += 2
          board[len(board)-1][j] += (uint64(2) << 60)
        } else {
          board[i-1][y] += 2
          board[i-1][j] += (uint64(2) << 60)
        }
        board[i][y] += 2
        if i == len(board)-1 {
          board[0][y] += 2
          board[0][j] += (uint64(2) << 60)
        } else {
          board[i+1][y] += 2
          board[i+1][j] += (uint64(2) << 60)
        }
      }
      // 2nd-leftmost bit of uint64 adds to its left
      m = uint64(1) << 56
      v = board[i][j] & m
      if v > 0 {
        if i == 0 {
          board[len(board)-1][j] += (uint64(2) << 60)
        } else {
          board[i-1][j] += (uint64(2) << 60)
        }
        board[i][j] += (uint64(2) << 60)
        if i == len(board)-1 {
          board[0][j] += (uint64(2) << 60)
        } else {
          board[i+1][j] += (uint64(2) << 60)
        }
      }
      // middle bits of uint64
      for k := 52; k >= 0; k -= 4 {
        m := uint64(0x111) << uint(k)
        v := board[i][j] & m
        if v > 0 {
          var u, w uint64
          switch v >> uint(k) {
          case 0x010:
            u = 1
            w = 0
          case 0x001, 0x100:
            u = 1
            w = 1
          case 0x011, 0x110:
            u = 2
            w = 1
          case 0x101:
            u = 2
            w = 2
          case 0x111:
            u = 3
            w = 2
          }
          u <<= uint(k + 5)
          if i == 0 {
            board[len(board)-1][j] += u
          } else {
            board[i-1][j] += u
          }
          board[i][j] += (w << uint(k+5))
          if i == len(board)-1 {
            board[0][j] += u
          } else {
            board[i+1][j] += u
          }
        }
      }
      // rightmost bit of uint64 adds to its 3 right and up and down
      m = uint64(1)
      v = board[i][j] & m
      if v > 0 {
        if j == len(board[i])-1 {
          y = 0
        } else {
          y = j + 1
        }
        if i == 0 {
          board[len(board)-1][y] += (uint64(2) << 60)
          board[len(board)-1][j] += 2
        } else {
          board[i-1][y] += (uint64(2) << 60)
          board[i-1][j] += 2
        }
        board[i][y] += (uint64(2) << 60)
        if i == len(board)-1 {
          board[0][y] += (uint64(2) << 60)
          board[0][j] += 2
        } else {
          board[i+1][y] += (uint64(2) << 60)
          board[i+1][j] += 2
        }
      }
      // 2nd-rightmost bit of uint64 adds to its right
      m = uint64(0x10)
      v = board[i][j] & m
      if v > 0 {
        if i == 0 {
          board[len(board)-1][j] += 2
        } else {
          board[i-1][j] += 2
        }
        board[i][j] += 2
        if i == len(board)-1 {
          board[0][j] += 2
        } else {
          board[i+1][j] += 2
        }
      }
    }
  }
  time.Sleep(interframe)
}

func rules(board [][]uint64) {
  for i := range board {
    for j := range board[i] {
      for k := 60; k >= 0; k -= 4 {
        m := uint64(0xF) << uint(k)
        v := board[i][j] & m
        if v > 0 {
          board[i][j] &= ^m
          switch v >> uint(k) {
          case 5, 6, 7:
            board[i][j] |= (uint64(1) << uint(k))
          }
        }
      }
    }
  }
}

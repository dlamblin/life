Conway's Game of Life bitwise in Go
===

Having thought about using a nibble to both store the currect state of a cell
and accumulate a count of it's neighbor (in the higher three bits) I went
ahead and implemented in Go.  It may not make the most memory/size compact
implementation, but it was a fun excercise.


For now, try using a terminal the size of 96x34 to run the gosper gun.

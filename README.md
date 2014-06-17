Conway's Game of Life
===
Bitwise in Go
==

Having thought about using a nibble to both store the currect state of a cell
and accumulate a count of its neighbor (in the higher three bits) I went
ahead and implemented in Go.  It may not make the most memory/size compact
implementation, but it was a fun excercise.

It renders steps to the terminal.
For now, try using a terminal the size of 96x34 to run the gosper gun.

I later realized that if the higher 3 bits of the nibble can only store 0-7 and
the lowest bit of the nibble stores the current state, then having 8 neighbors
will roll over the bit into the state of the cell to the left. This seems like
a problem, but I need to contrive an example that shows the problem happening
because, EG the Gosper gun seems to work fine as does an isolated 3x3 block.


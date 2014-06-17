#Conway's Game of Life
##Bitwise in Go

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

The thought process here was: if you need two grids to store the current and
next state, then maybe you could do better if you stored the current state and
counts in the same grid. This is only true if a boolean array takes a byte per
boolean, then maybe you could pack a count and the state in less than a byte.

Honestly if I had to do this quickly, or again, I'd either bit pack the
state entirely as bits, and the maintain a second bit pack array for the next
state, or I'd just have a byte array for each; possibly a byte array that
contains state and count and works similarly to this code.

Here I thought rolling over from one uint64 to the next was hard, but once I
had that it was easy to wrap the board edges. It might be harder with a byte
array or not.

I'm thinking of:   
Implementing a couple of versions with different board representations as
above, adding an option to randomly spawn, and an option to skip showing the
first n frames, adding flags for all the consts, making board wrapping and
uint64 alignment optional, making the board more of an object with methods,
and finally seeing if a go routine could update the board faster, which means
profiling this a little.


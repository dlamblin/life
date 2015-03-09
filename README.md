#Conway's Game of Life
##Bitwise in Go

Having thought about using a nibble to both store the current state of a cell
and accumulate a count of its neighbor (in the higher three bits) I went
ahead and implemented in Go.  It may not make the most memory/size compact
implementation, but it was a fun exercise.

It renders steps to the terminal.
For now, try using a terminal the size of 96x34 to run the gosper gun.

###Why as nibbles
The thought process here was: if you need two grids to store the current and
next state, then maybe you could do better if you stored the current state and
counts in the same grid. This is only true if a boolean array takes a byte per
boolean, then maybe you could pack a count and the state in less than a byte.

Honestly if I had to do this quickly, or again, I'd either bit pack the
state entirely as bits, and then maintain a second bit pack array for the next
state, or I'd just have a byte array for each; possibly a byte array that
contains state and count and works similarly to this code.

Here I thought wrapping over from one uint64 to the next was hard, but once I
had that it was easy to wrap the board edges. It might be harder with a byte
array or not.

###Nibble roll over
The higher 3 bits of the nibble can only store 0-7 and the lowest bit of the
nibble stores the current state, then having 8 neighbors will roll over the
bit into the state of the cell to the left. This seems like a problem, but I
need to contrive an example that shows the problem happening because, EG the
Gosper gun seems to work fine as does an isolated 3x3 block.

It seems to me that a 7 rolling to 0 within the cell still indicates the cell
will die regardless, and the 7 to 8 roll over must happen from a live or dead
cell into its left neighbor who is live and who must have at minimum 4
neighbors which indicates death already. That cell will be rolled to the dead
state and its count incremented, possibly even rolling further left. The left
cells will already have been processed, so it will not miss contributing to
the counts as a neighbor, and once the roll over reaches the most significant
bit of the byte it should not affect another cell.

###Futher work
I'm thinking of:   
Implementing a couple of versions with different board representations as
above, adding an option to randomly spawn, and an option to skip showing the
first n frames, adding flags for all the consts, making board wrapping and
uint64 alignment optional, making the board more of an object with methods,
and finally seeing if a go routine could update the board faster, which means
profiling this a little.


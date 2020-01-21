# automata

A command line program that prints elementary cellular automata.  

## usage

Use flags to set the rule (aka Wolfram code), the amount of steps and the initial state.
Defaults are: `-rule=30` `-step=16` `-init="0000000000000001000000000000000"`.
In other words`./automata` is the same as `./automata -rule 30 -step 16 -init "0000000000000001000000000000000"`.

Here are some examples of various ways to set the flags:

* `./automata -rule 110 -step 5 -init "00100"`  
* `./automata -rule=12 --init "00000000"`  
* `./automata --rule=69 --step 1`  
* `./automata -help`  

## what are we printing here

We are printing elementary cellular automata([wikipedia](https://en.wikipedia.org/wiki/Elementary_cellular_automaton)) for no practical purpose at all.

When updating cell values, the program thinks the first and the last cell are connected (wrapping around).

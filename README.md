# Monty Hall Problem

This is a little program to demonstrate the effect of the [Monty Hall Problem](https://en.wikipedia.org/wiki/Monty_Hall_problem)

The result of the run with the default values (`const numberOfDoors = 3` and `const numberOfRuns = 1000000`) is:

```
Won with switch: 66.6883 %
Won without switch: 33.3821 %
``` 

The runs are made separately so the actual number of runs is `numberOfRuns * 2`. This is so a sceptical person will not have to trust expectations like "would have won, if..."

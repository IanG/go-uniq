# go-uniq

This is a solution to [John Crickett](https://www.linkedin.com/in/johncrickett/)'s Coding Challenges [Build Your Own uniq Tool](https://codingchallenges.fyi/challenges/challenge-uniq) in go.

This is a very basic implementation of the `uniq` command used in the shell on most *nix-like operating systems.
It makes use of the `flags` go package for handling command line parameters

## Building Locally

build the code with:

`go build`

this will create the `uniq` binary in your local directory.

## Running

### Getting Help
View command line options with `./uniq -h`.  This will output:

```
Usage: ./uniq [OPTIONS] [inputfile] [outputfile]
  -c	Count occurrences
  -count
    	Count occurrences
  -d	Print only repeated lines
  -repeated
    	Print only repeated lines
  -u	Print only unique lines
  -unique
    	Print only unique lines
  [inputfile] can either be the path to an input file or '-' for stdin
  [outputfile] can either be the path to an output file or '-' for stdout
```

### Finding the unique lines in a file

You can find the unique lines in a file with:

`./uniq ./testdata/test.txt`

This will output:

```
line1
line2
line3
line4
```
You can also pipe input from stdin instead of specifying the filename with:

`cat ./testdata/test.txt | ./uniq -`


### Finding the repeated lines in a file

You can find the repeated lines in a file with:

`./uniq -d ./testdata/test.txt`

This will output:

```
line2
```

You can find out how many times repeated lines occur in the file with:

`./uniq -d -c ./testdata/test.txt`

This will output:

```
2 line2
```
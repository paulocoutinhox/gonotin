# gonotin (Go not in)

This simple scripts compare file A and file B and shows line that NOT IN inside file B.

To execute the golang version:

    go run gonotin.go fileA.txt fileB.txt

or for performance:

    go build gonotin.go
    gonotin fileA.txt fileB.txt

To execute the php version:

    php php-not-in.php fileA.txt fileB.txt

You can execute and send the output to a file, example:

    go run gonotin.go fileA.txt fileB.txt > result.txt

or

    go build gonotin.go
    gonotin fileA.txt fileB.txt > result.txt

To execute sample file generator you can use:

    go run generate.go fileA.txt 100000 120000

> Syntax: go run generate.go [filename] [quantity of rows] [max random number]


Thanks.

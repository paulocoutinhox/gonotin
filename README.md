# gonotin (Go not in)

This simple scripts compare file A and file B and shows line that NOT IN inside file B.

To execute the golang version:

    go run main.go fileA.csv fileB.csv

or

    go build  
    gonotin fileA.csv fileB.csv

To execute the php version:

    php php-not-in.php fileA.csv fileB.csv

You can execute and send the output to a file, example:

    go run main.go fileA.csv fileB.csv > result.txt

or

    go build  
    gonotin fileA.csv fileB.csv > result.txt

To execute sample file generator you can use:

    go run generate.go fileA.csv 100000 120000

> Syntax: go run generate.go [filename] [quantity of rows] [max random number]


Thanks.

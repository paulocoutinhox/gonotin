# gonotin

This simple scripts compare file A and file B and shows line that NOT IN inside file B.

To execute the golang version:

> go run main.go fileA.csv fileB.csv

or

> go build  
> gonotin fileA.csv fileB.csv

To execute the php version:

> php php-not-in.php fileA.csv fileB.csv

You can execute and send the output to a file, example:

> go run main.go fileA.csv fileB.csv > result.txt

or

> go build  
> gonotin fileA.csv fileB.csv > result.txt


Thanks.

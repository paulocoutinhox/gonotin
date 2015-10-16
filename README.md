# gonotin (Go Not In)

This simple scripts compare file A and file B and shows line that not exists in file B, but exists in file A.

# Automated running

    ./run.sh

# Running Go version 

    Syntax: go run go-version/gonotin.go [file A] [file B] [mode]  

```
go run go-version/gonotin.go data/fileA-large.txt data/fileB-large.txt 1
```


or for better performance:

```
go build -o go-version/gonotin go-version/gonotin.go
./go-version/gonotin data/fileA-large.txt data/fileB-large.txt 1
```    

# Running PHP version 

    Syntax: php php-version/gonotin.php [file A] [file B] [mode] 

```
php php-version/gonotin.php data/fileA-large.txt data/fileB-large.txt 1
```

# Running Python version

    Syntax: MODE=[mode] python python-version/gonotin.py [file A] [file B] 

```
MODE=1 python python-version/gonotin.py data/fileA-large.txt data/fileB-large.txt
```

# Running NodeJS version

    Syntax: node nodejs-version/gonotin.js [file A] [file B] [mode] 

```
node nodejs-version/gonotin.js data/fileA-large.txt data/fileB-large.txt 1
```

# Running C++ version

    Buiding:
```
cmake -Bcpp-version/build/ -Hcpp-version/
make -C cpp-version/build/
mv cpp-version/build/gonotin cpp-version/
rm -rf cpp-version/build
```

    Syntax: ./cpp-version/gonotin [file A] [file B] [mode] 

```
./cpp-version/gonotin data/fileA-large.txt data/fileB-large.txt 1
```

# Running options

    Time: You can use "time" unix/osx tool to get the executing time before the command. Every command on benchmark use it.
    
```
time go run go-version/gonotin.go data/fileA-large.txt data/fileB-large.txt 1
```

    Results: You can send results to a file. Every command on benchmark use it.

```
time go run go-version/gonotin.go data/fileA-large.txt data/fileB-large.txt 1 > results.txt
```

# Running the sample file generator

    Syntax: go run generator/generator.go [filename] [quantity of rows] [max random number]

```
go run go-generator/generator.go data/fileA.txt 100000 120000
go run go-generator/generator.go data/fileB.txt 100000 120000
```

# Benchmark results

Check file [RESULTS.md](RESULTS.md)

# Author WebSite

    http://www.pcoutinho.com

# License

    MIT

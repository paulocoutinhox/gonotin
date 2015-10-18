# gonotin (Go Not In)

This simple scripts compare file A and file B and shows line that not exists in file B, but exists in file A.

# Automated running

    SYNTAX: 
    ./run.sh [file size - small, medium or large]

```
./run.sh large
```

# Running Go version 

    SYNTAX: 
    go run go-version/gonotin[version].go [file A] [file B]  

```
go run go-version/gonotin1.go data/fileA-large.txt data/fileB-large.txt
```


    SYNTAX: 
    for better performance (our benchmark use it):

```
mkdir -p temp
go build -o temp/go-gonotin-1 go-version/gonotin1.go
./temp/go-gonotin-1 data/fileA-large.txt data/fileB-large.txt
```    

# Running PHP version (removed until a correct code)

    SYNTAX: 
    php php-version/gonotin[version].php [file A] [file B] 

```
php php-version/gonotin1.php data/fileA-large.txt data/fileB-large.txt
```

# Running Python version

    SYNTAX: 
    python python-version/gonotin[version].py [file A] [file B] 

```
python python-version/gonotin1.py data/fileA-large.txt data/fileB-large.txt
```

# Running NodeJS version (removed until a correct code)

    SYNTAX: 
    node nodejs-version/gonotin[version].js [file A] [file B] 

```
node nodejs-version/gonotin1.js data/fileA-large.txt data/fileB-large.txt
```

# Running C++ version

    BUILDING FOR OSX:
    
```
mkdir -p temp
clang++ -g cpp-version/gonotin[version].cpp -o temp/cpp-gonotin-[version] -lm -std=c++11
```


    BUILDING FOR LINUX:
    
```
mkdir -p temp
g++ -g -pthread cpp-version/gonotin[version].cpp -o temp/cpp-gonotin-[version] -lm -std=c++11
```


    SYNTAX: 
    ./temp/cpp-gonotin-[version] [file A] [file B]

```
./temp/cpp-gonotin-1 data/fileA-large.txt data/fileB-large.txt
```

# Running options

    TIME: 
    You can use "time" unix/osx tool to get the executing time before the command. Every command on benchmark use own time calculation.
    
```
time go run go-version/gonotin.go data/fileA-large.txt data/fileB-large.txt 1
```

    RESULTS: 
    You can send results to a file. Every command on benchmark use it.

```
time go run go-version/gonotin1.go data/fileA-large.txt data/fileB-large.txt > results.txt
```

# Running the sample file generator

    SYNTAX: 
    go run generator/generator.go [filename] [quantity of rows] [max random number]

```
go run go-generator/generator.go data/fileA.txt 100000 120000
go run go-generator/generator.go data/fileB.txt 100000 120000
```

# Running on OSX

    If you plan execute the "run.sh" script on OSX, please install "coreutils" to get the "gdate" command because the "BSD" "date" command remove the "nanoseconds" precision:

```
brew install coreutils
```

# Benchmark results

Check file [RESULTS.md](RESULTS.md)

# Author WebSite

    http://www.pcoutinho.com

# License

    MIT

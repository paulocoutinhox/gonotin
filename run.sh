#! /bin/bash

################################################################################

# functions
startTimer() {
    if [[ "$OSTYPE" == "darwin"* ]]; then
        start=$(gdate +%s.%N)
    else
        start=$(date +%s.%N)
    fi
}

endTimer() {
    if [[ "$OSTYPE" == "darwin"* ]]; then
        dur=$(echo "$(gdate +%s.%N) - $start" | bc);
    else
        dur=$(echo "$(date +%s.%N) - $start" | bc);
    fi

    printf "\n> Execution time: %.6f seconds\n" $dur
}

################################################################################

echo "> Compiling Go file..."
rm -rf go-version/gonotin
go build -o go-version/gonotin go-version/gonotin.go
echo ""

################################################################################

echo "> Compiling C++ file..."

rm -rf cpp-version/gonotin
rm -rf cpp-version/build

# OSX/*nix compiling method
#g++ -g cpp-version/gonotin.cpp -o cpp-version/gonotin -lm

# Cross-Platform compiling method
cmake -Bcpp-version/build/ -Hcpp-version/
make -C cpp-version/build/
mv cpp-version/build/gonotin cpp-version/
rm -rf cpp-version/build
echo ""

################################################################################

echo "> Removing old results..."
rm -rf results
mkdir results
echo ""

################################################################################

echo "> Processing large files..."

CWD=$(pwd)
SIZE=$1
FILE_A=$CWD/data/fileA-$SIZE.txt
FILE_B=$CWD/data/fileB-$SIZE.txt
SUFFIX=.txt
EXECUTABLE=gonotin

echo ""

################################################################################

for i in `seq 1 1`; do
    echo "> Executing C++ (mode = $i, size = $SIZE) version..."
    startTimer
    cpp-version/$EXECUTABLE $FILE_A $FILE_B > results/cpp-$SIZE-mode-$i$SUFFIX
    endTimer
    echo ""
done

################################################################################

for i in `seq 1 7`; do
    echo "> Executing Go (mode = $i, size = $SIZE) version..."
    startTimer
    go-version/$EXECUTABLE $FILE_A $FILE_B $i > results/go-$SIZE-mode-$i$SUFFIX
    endTimer
    echo ""
done


################################################################################

for i in `seq 1 4`; do
    echo "> Executing Python (mode = $i, size = $SIZE) version..."
    startTimer
    MODE=$i python python-version/$EXECUTABLE.py $FILE_A $FILE_B > results/python-$SIZE-mode-$i$SUFFIX
    endTimer
    echo ""
done

################################################################################

for i in `seq 1 2`; do
    echo "> Executing PHP (mode = $i, size = $SIZE) version..."
    startTimer
    php php-version/$EXECUTABLE.php $FILE_A $FILE_B $i > results/php-$SIZE-mode-$i$SUFFIX
    endTimer
    echo ""
done

################################################################################

for i in `seq 1 1`; do
    echo "> Executing Node (mode = $i, size = $SIZE) version..."
    startTimer
    node nodejs-version/$EXECUTABLE.js $FILE_A $FILE_B > results/nodejs-$SIZE-mode-$i$SUFFIX
    endTimer
    echo ""
done

################################################################################

################################################################################

echo "> Finished!"
echo ""
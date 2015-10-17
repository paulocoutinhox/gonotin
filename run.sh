#! /bin/bash

################################################################################

# variables
GO_INSTALLED=0
CPP_INSTALLED=0
PYTHON_INSTALLED=0
PHP_INSTALLED=0
NODEJS_INSTALLED=0

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

    printf "> Execution time: %.6f seconds\n\n" $dur
}

checkGoInstalled() {
    if [ -x "$(command -v go)" ]; then
        GO_INSTALLED=1
    fi
}

checkCppInstalled() {
    if [[ "$OSTYPE" == "darwin"* ]]; then
        if [ -x "$(command -v clang++)" ]; then
            CPP_INSTALLED=1
        fi
    else
        if [ -x "$(command -v g++)" ]; then
            CPP_INSTALLED=1
        fi
    fi
}

checkPythonInstalled() {
    if [ -x "$(command -v python)" ]; then
        PYTHON_INSTALLED=1
    fi
}

checkPhpInstalled() {
    if [ -x "$(command -v php)" ]; then
        PHP_INSTALLED=1
    fi
}

checkNodeJsInstalled() {
    if [ -x "$(command -v node)" ]; then
        NODEJS_INSTALLED=1
    fi
}

################################################################################

# check installed apps

checkGoInstalled
checkCppInstalled
checkPythonInstalled
checkPhpInstalled
checkNodeJsInstalled

################################################################################

if [ $GO_INSTALLED -eq 1 ]; then
    echo "> Compiling Go file..."
    rm -rf go-version/gonotin
    go build -o go-version/gonotin go-version/gonotin.go
    echo ""
else
    echo "> ERROR: Go is not installed!"
    echo ""
fi

################################################################################

if [ $CPP_INSTALLED -eq 1 ]; then
    echo "> Compiling C++ file..."

    rm -rf cpp-version/gonotin
    rm -rf cpp-version/build

    # OSX/*nix compiling method
    if [[ "$OSTYPE" == "darwin"* ]]; then
        clang++ -g cpp-version/gonotin.cpp -o cpp-version/gonotin -lm -std=c++11
    else
        g++ -g -pthread cpp-version/gonotin.cpp -o cpp-version/gonotin -lm -std=c++11
    fi

    # Cross-Platform compiling method
    #cmake -Bcpp-version/build/ -Hcpp-version/
    #make -C cpp-version/build/
    #mv cpp-version/build/gonotin cpp-version/
    #rm -rf cpp-version/build

    echo ""
else
    echo "> ERROR: C++ is not installed!"
    echo ""
fi

################################################################################

echo "> Removing old results..."
rm -rf results
mkdir results
echo ""

################################################################################

CWD=$(pwd)
SIZE=$1
FILE_A=$CWD/data/fileA-$SIZE.txt
FILE_B=$CWD/data/fileB-$SIZE.txt
SUFFIX=.txt
EXECUTABLE=gonotin

echo "> Processing $SIZE files..."
echo ""

################################################################################

if [ $CPP_INSTALLED -eq 1 ]; then
    for i in `seq 1 5`; do
        echo "> Executing C++ (mode = $i, size = $SIZE) version..."
        startTimer
        cpp-version/$EXECUTABLE $FILE_A $FILE_B $i > results/cpp-$SIZE-mode-$i$SUFFIX
        endTimer
        echo ""
    done
else
    echo "> ERROR: G++ is not installed!"
    echo ""
fi

################################################################################

if [ $GO_INSTALLED -eq 1 ]; then
    for i in `seq 1 7`; do
        echo "> Executing Go (mode = $i, size = $SIZE) version..."
        startTimer
        go-version/$EXECUTABLE $FILE_A $FILE_B $i > results/go-$SIZE-mode-$i$SUFFIX
        endTimer
        echo ""
    done
else
    echo "> ERROR: Go is not installed!"
    echo ""
fi

################################################################################

if [ $PYTHON_INSTALLED -eq 1 ]; then
    for i in `seq 1 4`; do
        echo "> Executing Python (mode = $i, size = $SIZE) version..."
        startTimer
        MODE=$i python python-version/$EXECUTABLE.py $FILE_A $FILE_B > results/python-$SIZE-mode-$i$SUFFIX
        endTimer
        echo ""
    done
else
    echo "> ERROR: Python is not installed!"
    echo ""
fi

################################################################################

if [ $PHP_INSTALLED -eq 1 ]; then
    for i in `seq 1 2`; do
        echo "> Executing PHP (mode = $i, size = $SIZE) version..."
        startTimer
        php php-version/$EXECUTABLE.php $FILE_A $FILE_B $i > results/php-$SIZE-mode-$i$SUFFIX
        endTimer
        echo ""
    done
else
    echo "> ERROR: Php is not installed!"
    echo ""
fi

################################################################################

if [ $NODEJS_INSTALLED -eq 1 ]; then
    for i in `seq 1 1`; do
        echo "> Executing Node (mode = $i, size = $SIZE) version..."
        startTimer
        node nodejs-version/$EXECUTABLE.js $FILE_A $FILE_B > results/nodejs-$SIZE-mode-$i$SUFFIX
        endTimer
        echo ""
    done

else
    echo "> ERROR: NodeJs is not installed!"
    echo ""
fi

################################################################################

echo "> Finished!"
echo ""
#! /bin/bash

################################################################################

echo ""

# variables
GO_INSTALLED=0
CPP_INSTALLED=0
PYTHON_INSTALLED=0
PHP_INSTALLED=0
NODEJS_INSTALLED=0
TEMP_FOLDER=temp

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

echo "> Removing old temporary files..."
rm -rf temp
mkdir temp
echo ""

################################################################################

if [ $GO_INSTALLED -eq 1 ]; then
    echo "> Compiling Go files..."

    for i in `seq 1 3`; do
        echo "> Compiling file $i..."

        rm -rf $TEMP_FOLDER/go-gonotin-$i
        go build -o $TEMP_FOLDER/go-gonotin-$i go-version/gonotin$i.go
    done

    echo ""
else
    echo "> ERROR: Go is not installed!"
    echo ""
fi

################################################################################

if [ $CPP_INSTALLED -eq 1 ]; then
    echo "> Compiling C++ files..."

    for i in `seq 1 3`; do
        echo "> Compiling file $i..."

        rm -rf $TEMP_FOLDER/cpp-gonotin-$i

        # OSX and Linux compiling method
        if [[ "$OSTYPE" == "darwin"* ]]; then
            clang++ -g cpp-version/gonotin$i.cpp -o $TEMP_FOLDER/cpp-gonotin-$i -lm -std=c++11
        else
            g++ -g -pthread cpp-version/gonotin$i.cpp -o $TEMP_FOLDER/cpp-gonotin-$i -lm -std=c++11
        fi
    done

    echo ""
else
    echo "> ERROR: C++ is not installed!"
    echo ""
fi

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

if [ $GO_INSTALLED -eq 1 ]; then
    for i in `seq 1 3`; do
        echo "> Executing Go (mode = $i, size = $SIZE) version..."
        startTimer
        $TEMP_FOLDER/go-$EXECUTABLE-$i $FILE_A $FILE_B $i > $TEMP_FOLDER/go-$SIZE-mode-$i$SUFFIX
        endTimer
        echo ""
    done
else
    echo "> ERROR: Go is not installed!"
    echo ""
fi

################################################################################

if [ $CPP_INSTALLED -eq 1 ]; then
    for i in `seq 1 3`; do
        echo "> Executing C++ (mode = $i, size = $SIZE) version..."
        startTimer
        $TEMP_FOLDER/cpp-$EXECUTABLE-$i $FILE_A $FILE_B $i > $TEMP_FOLDER/cpp-$SIZE-mode-$i$SUFFIX
        endTimer
        echo ""
    done
else
    echo "> ERROR: G++ is not installed!"
    echo ""
fi

################################################################################

if [ $PYTHON_INSTALLED -eq 1 ]; then
    for i in `seq 1 3`; do
        echo "> Executing Python (mode = $i, size = $SIZE) version..."
        startTimer
        MODE=$i python python-version/$EXECUTABLE$i.py $FILE_A $FILE_B > $TEMP_FOLDER/python-$SIZE-mode-$i$SUFFIX
        endTimer
        echo ""
    done
else
    echo "> ERROR: Python is not installed!"
    echo ""
fi

################################################################################

#if [ $PHP_INSTALLED -eq 1 ]; then
#    for i in `seq 1 2`; do
#        echo "> Executing PHP (mode = $i, size = $SIZE) version..."
#        startTimer
#        php php-version/$EXECUTABLE$i.php $FILE_A $FILE_B $i > $TEMP_FOLDER/php-$SIZE-mode-$i$SUFFIX
#        endTimer
#        echo ""
#    done
#else
#    echo "> ERROR: Php is not installed!"
#    echo ""
#fi

################################################################################

#if [ $NODEJS_INSTALLED -eq 1 ]; then
#    for i in `seq 1 1`; do
#        echo "> Executing Node (mode = $i, size = $SIZE) version..."
#        startTimer
#        node nodejs-version/$EXECUTABLE$i.js $FILE_A $FILE_B > $TEMP_FOLDER/nodejs-$SIZE-mode-$i$SUFFIX
#        endTimer
#        echo ""
#    done
#
#else
#    echo "> ERROR: NodeJs is not installed!"
#    echo ""
#fi

################################################################################

echo "> Finished!"
echo ""
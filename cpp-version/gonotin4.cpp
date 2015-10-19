/*
AUTHOR: PAULO COUTINHO
DESCRIPTION: Using std::threads
*/

#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
#include <set>
#include <unordered_set>
#include <thread>

using namespace std;

// general
std::unordered_set<string> dataA;
std::unordered_set<string> dataB;

void readerA(char *file)
{
    ifstream stream;
    stream.open(file);

    while(!stream.eof())
    {
        string line;
        stream >> line;
        dataA.insert(line);
    }

    stream.close();
}

void readerB(char *file)
{
    ifstream stream;
    stream.open(file);

    while(!stream.eof())
    {
        string line;
        stream >> line;
        dataB.insert(line);
    }

    stream.close();
}

int main(int argc, char *argv[])
{
    // create the threads
    std::thread t1(readerA, argv[1]);
    std::thread t2(readerB, argv[2]);

    // wait for threads finish
    t1.join();
    t2.join();

    // process data
    for(string value : dataA)
    {
        auto search = dataB.find(value);

        if (search == dataB.end())
        {
            printf("%s\n", value.c_str());
        }
    }

    return 0;
}

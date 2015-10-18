/*
AUTHOR: PAULO COUTINHO
*/

#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
#include <set>
#include <unordered_set>
#include <thread>
#include <future>

using namespace std;

std::unordered_set<string> reader(char *file)
{
    std::unordered_set<string> data;

    ifstream stream;
    stream.open(file);

    while(!stream.eof())
    {
        string line;
        stream >> line;
        data.insert(line);
    }

    stream.close();
    return data;
}

int main(int argc, char *argv[])
{
    // generate future async
    auto futureA = std::async(reader, argv[1]);
    auto futureB = std::async(reader, argv[2]);

    // get data from future
    std::unordered_set<string> dataA = futureA.get();
    std::unordered_set<string> dataB = futureB.get();

    // process data
    for(string value : dataA)
    {
        auto search = dataB.find(value);

        if (search == dataB.end())
        {
            printf("%s\n", value.c_str());
        }
    }
}

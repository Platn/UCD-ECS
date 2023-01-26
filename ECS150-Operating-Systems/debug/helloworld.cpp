#include <iostream>
#include <vector>
#include <string>

using namespace std;

int main()
{
    vector<string> msg;
    msg.push_back("Hello");
    msg.push_back("World");
    
    cout << msg[0] << endl;
    return 0;
}
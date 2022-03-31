#include <iostream>

#include "factorial.cpp"

using namespace std;

int main(){
    cout << Factorial<20>().data << endl;
    return 0;
}
#include <iostream>

using namespace std;

template<int N>
class Factorial{
public:
    Factorial<N>(){
        data = N * Factorial<N-1>().data;
        cout << data << endl;
    }
    long long data;
    

};
template<>
class Factorial<1>{
public:
    Factorial(){
        data = 1; 
        cout << data << endl;
    }
    long long data;
    

};

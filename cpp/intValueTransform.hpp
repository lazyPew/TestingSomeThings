// typedef signed short int16_t;
// typedef signed char int8_t;
#include <iostream>

// typedef unsigned short uint16_t;
// typedef unsigned char uint8_t;

// using namespace std;

class Transformer{
public:
    static void start(int);
    
private:
    static void startTransformation(int);
    static void startWrongTransformation(int);

    void print8BitValues(int8_t*);
    void print32BitValue(int);

    int16_t *transform32BitValueTo16BitValues(int);
    void transform16BitValueTo8BitValues(int16_t);
    void transform8BitValuesTo16BitValue(int8_t);
    void transform16BitValuesTo32BitValue();
};
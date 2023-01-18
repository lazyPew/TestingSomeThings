#include <iostream>

typedef signed short int16_t;
typedef signed char int8_t;

typedef unsigned short uint16_t;
typedef unsigned char uint8_t;

// using namespace std;

class Transformer{
public:
    static void start();
    
private:
    static void startTransformation(int);
    static void startWrongTransformation(int);

    void print8BitValues(int8_t*);
    void print32BitValue(int);

    static uint8_t *transform32BitValueTo16BitValues(int value);
    static uint8_t *transform16BitValueTo8BitValues(uint16_t *values16Bit);
    void transform8BitValuesTo16BitValue(int8_t);
    void transform16BitValuesTo32BitValue();
};
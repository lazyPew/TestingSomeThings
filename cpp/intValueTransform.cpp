#include "intValueTransform.hpp"

void Transformer::start(){
    int valueToTransform;
    std::cout << "Which value to transform?\n";
    std::cin >> valueToTransform;
    if (std::cin.fail()) {
        std:: cout << "ERROR\n";
        return;
    }
    std::cout << "Trying to transform 32bit value = " << valueToTransform << std::endl;

    startTransformation(valueToTransform);

    startWrongTransformation(valueToTransform);
}

void Transformer::startTransformation(int valueToTransform){
    // TODO need to refactor
    int16_t* byte = (int16_t*)(&valueToTransform);

    uint16_t *values16Bit = new uint16_t[2];//transform32BitValueTo16BitValues(valueToTransform);
    uint8_t *values8BitFirst = new uint8_t[2];
    uint8_t *values8BitSecond = new uint8_t[2];
    
    values16Bit[0] = *byte;
    std::cout  << "values16Bit[0] = " << values16Bit[0] << std::endl;
    uint8_t* first8BitValues = (uint8_t*)(&values16Bit[0]);
    values8BitFirst[0] = *first8BitValues;
    values8BitFirst[1] = *(++first8BitValues);
    std::cout << "values8BitFirst[0] = " << unsigned(values8BitFirst[0]) << std::endl;
    std::cout << "values8BitFirst[1] = " << unsigned(values8BitFirst[1]) << std::endl;

    values16Bit[1] = *(++byte);
    std::cout << "values16Bit[1] = " << values16Bit[1] << std::endl;
    uint8_t* second8BitValues = (uint8_t*)(&values16Bit[1]);
    values8BitSecond[0] = *second8BitValues;
    values8BitSecond[1] = *(++second8BitValues);
    std::cout << "values8BitSecond[0] = " << unsigned(values8BitSecond[0]) << std::endl;
    std::cout << "values8BitSecond[1] = " << unsigned(values8BitSecond[1]) << std::endl;

    int result = ((((values8BitSecond[1] << 8) + values8BitSecond[0]) << 16) + ((values8BitFirst[1] << 8) + values8BitFirst[0]));
    uint resultU = ((((values8BitSecond[1] << 8) + values8BitSecond[0]) << 16) + ((values8BitFirst[1] << 8) + values8BitFirst[0]));
    std::cout << result << " " << resultU << std::endl;

    int result2 = ((((values8BitSecond[1] * 256) + values8BitSecond[0]) * 65536) + ((values8BitFirst[1] * 256) + values8BitFirst[0]));
    uint resultU2 = ((((values8BitSecond[1] * 256) + values8BitSecond[0]) * 65536) + ((values8BitFirst[1] * 256) + values8BitFirst[0]));
    std::cout << result2 << " " << resultU2 << std::endl;

    // uint8_t *test = transform32BitValueTo16BitValues(valueToTransform);

    // test = test;
}

uint8_t* Transformer::transform32BitValueTo16BitValues(int value){
    // // int16_t* byte = (int16_t*)(&value);
    // uint16_t *values16Bit = new uint16_t[2];
    // // uint8_t *values8Bit = new uint8_t[4];
    // values16Bit[0] = value;
    // values16Bit[1] = value >> 16;
    // return transform16BitValueTo8BitValues(values16Bit);
    // // return values16Bit;
}

uint8_t* Transformer::transform16BitValueTo8BitValues(uint16_t *values16Bit){
    // uint8_t *values8Bit = new uint8_t[sizeof(values16Bit)];

    // std::cout << "size16 " << sizeof(values16Bit)  << " "  << sizeof(values16Bit) / sizeof(uint16_t *) << values16Bit[0] << values16Bit[1] << std::endl;
    // std::cout << "size8 " << sizeof(values8Bit) << " " << sizeof(values8Bit) / sizeof(uint8_t *)  << std::endl;
    
    // return values8Bit;
}

void Transformer::transform8BitValuesTo16BitValue(int8_t){

}

void Transformer::transform16BitValuesTo32BitValue(){

}

void print8BitValues(int8_t*){
    
}

void print32BitValue(int){

}

void Transformer::startWrongTransformation(int valueToTransform){
    int16_t* byte = (int16_t*)(&valueToTransform);

    int16_t *values16Bit = new int16_t[2];//transform32BitValueTo16BitValues(valueToTransform);
    int8_t *values8BitFirst = new int8_t[2];
    int8_t *values8BitSecond = new int8_t[2];
    
    values16Bit[0] = *byte;
    int8_t* first8BitValues = (int8_t*)(&values16Bit[0]);
    values8BitFirst[0] = *first8BitValues;
    values8BitFirst[1] = *(++first8BitValues);

    values16Bit[1] = *(++byte);
    int8_t* second8BitValues = (int8_t*)(&values16Bit[1]);
    values8BitSecond[0] = *second8BitValues;
    values8BitSecond[1] = *(++second8BitValues);
}
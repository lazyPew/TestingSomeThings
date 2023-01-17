#include "intValueTransform.hpp"

void Transformer::start(int valueToTransform){
    startTransformation(valueToTransform);

    startWrongTransformation(valueToTransform);
}

void Transformer::startTransformation(int valueToTransform){
    int16_t* byte = (int16_t*)(&valueToTransform);

    uint16_t *values16Bit = new uint16_t[2];//transform32BitValueTo16BitValues(valueToTransform);
    uint8_t *values8BitFirst = new uint8_t[2];
    uint8_t *values8BitSecond = new uint8_t[2];
    
    values16Bit[0] = *byte;
    uint8_t* first8BitValues = (uint8_t*)(&values16Bit[0]);
    values8BitFirst[0] = *first8BitValues;
    values8BitFirst[1] = *(++first8BitValues);

    values16Bit[1] = *(++byte);
    uint8_t* second8BitValues = (uint8_t*)(&values16Bit[1]);
    values8BitSecond[0] = *second8BitValues;
    values8BitSecond[1] = *(++second8BitValues);

    uint16_t* end = values16Bit + (sizeof(values16Bit) / sizeof(uint16_t*));

    std::cout << values16Bit << std::endl;
    std::cout << sizeof(values16Bit) << std::endl;
    std::cout << sizeof(uint16_t*) << std::endl;

    for(uint16_t* x = values16Bit; x < end; x++)
    {
        std::cout << x << std::endl; 
    }

}

int16_t *Transformer::transform32BitValueTo16BitValues(int){
    int16_t *twoBytesValues = new int16_t[2];
    

    return twoBytesValues;
}

void Transformer::transform16BitValueTo8BitValues(int16_t){
    
}

void Transformer::transform8BitValuesTo16BitValue(int8_t){

}

void Transformer::transform16BitValuesTo32BitValue(){

}

void print8BitValues(int8_t*){
    
}

void print32BitValue(int){

}

void Transformer::startWrongTransformation(int){
    
}
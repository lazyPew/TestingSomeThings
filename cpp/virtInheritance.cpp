#include <iostream>
#include "virtInheritance.h"

void OldestClass::printString() const{
    std::cout << "string from oldest class, name is " << _name << std::endl;
}

TestingClass::TestingClass()
{

}

void TestingClass::testFunction(){
    printString();
}


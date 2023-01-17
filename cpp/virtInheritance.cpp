#include <iostream>
#include "virtInheritance.hpp"

void OldestClass::printString(){
    std::cout << "string from oldest class, name is " << _name << std::endl;
}

NeededClass::NeededClass()
{

}

void NeededClass::testFunction(){
    printString();
}


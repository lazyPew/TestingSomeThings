#include "includeFile.h"

using namespace std;

int option = -1;

void doOption(int);
void chooseOption(int);

int main(){

    chooseOption(option); 
    return 0;
}


void doOption(int option){
    TestingClass test = TestingClass();
    switch(option){
        case 1:
            test.testFunction();
            break;
        case 2:
            cout << Factorial<20>().data << endl;
            break;
        case 3:
            MapPrinter::checkMaps();
            break;
        case 4:
            
            break;
        case 0:
            break;
    }
}

void chooseOption(int option){
    if (option <= 0){
        do{
            cout << "Choose an option:\n" 
                << "1) Test VirtualInheritance\n"
                << "2) Test Factorial\n"
                << "3) Test std::map and std::unordered_map\n"
                << "4) Test binary numbers\n"
                << "0) Exit\n";
            cin >> option;
            if (!cin.fail()) {
                doOption(option);
            }
            else {
                cout << "Incorrect input. Finishing program." << endl;
                option = 0;
            } 
        } while(option != 0);
    }
    else doOption(option);
}

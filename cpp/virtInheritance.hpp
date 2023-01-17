#include <string>

/**
 * @brief Virtual inheritance is needed in cases of multipal inheritance
 * to avoid errors of ambigious functions.
 * 
 * "Ambigious function error" could happen when needed class inherits
 * two or more classes that inherit same base class. If Oldest class
 * have some non-private non-virtual function then same function will 
 * have both of the parent classes. 
 * 
 * In such situations needed class doesn't understand which version 
 * of needed function it should use and throws error. But in this case 
 * of virtual inheritance parameters that will use its value from first
 * non-virtual parent ???
 */

class OldestClass
{
public:
    void printString();

private:
    std::string _name = "oldest";
};

class ParentClass_1 : virtual public OldestClass
{

private:
    std::string _name = "parent 1";
};

class ParentClass_2 : virtual public OldestClass
{

private:
    std::string _name = "parent 2"; 
};

class NeededClass : public ParentClass_1
                   , public ParentClass_2
{
public:
    NeededClass();

    void testFunction();
};
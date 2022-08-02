#include <string>

/**
 * @brief 
 * 
 */
class OldestClass
{
public:
    void printString() const;

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

class TestingClass : public ParentClass_1
                   , public ParentClass_2
{
public:
    TestingClass();

    void testFunction();
};
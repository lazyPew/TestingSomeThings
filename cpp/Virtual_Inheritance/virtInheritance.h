#include <string>

class OldestClass
{
public:
    void printString() const;
private:
    std::string _name = "oldest";
};

class ParentClass_1 : public OldestClass
{

private:
    std::string _name = "parent 1";
};

class ParentClass_2 : public OldestClass
{

private:
    std::string _name = "parent 2";
};

class TestingClass : public virtual ParentClass_1
                   , public virtual ParentClass_2
{
public:
    TestingClass();

    void testFunction();
};
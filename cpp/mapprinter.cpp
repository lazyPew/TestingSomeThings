#include <iostream>
#include <string>
#include <map>
#include <unordered_map>

/**
 * @brief std::map and std::unordered_map have as its elements std::pair<const type_key,type_value>
 * 
 * Elements in std::map are sorted by its "Key" value but in std::unordered_map 
 * order is defined by its internal hashTable.
 * 
 * The key provided to std::unordered_map is hashed into indicies of hash table
 * that is why the performance of data structure depends on hash function a lot
 * but on an average, the cost of search, insert and delete from hash table is O(1).
 * 
 * Base for std::unordered_map is _Hashtable, dispatched at compile time via 
 * template alias __umap_hashtable.
 * 
 * In cases of creation std::unordered_map with N elements and std::unordered_map
 * with N-1 elements with insertion of 1 additional element or std::unordered_map
 * with N+1 elements with erasion of 1 element their order will be different.
 * Even though they will have all the same elements.
 */

class MapPrinter
{
public:
    static void checkMaps(){

        std::map<int,std::string> map1;
        std::unordered_map<int, std::string> umap1;
        std::unordered_map<int, std::string> umap11;
        map1.insert({1,"1"});
        map1.insert({5,"5"});
        map1.insert({3,"3"});
        map1.insert({4,"4"});
        map1.insert({2,"2"});
        umap1.insert({1,"1"});
        umap1.insert({5,"5"});
        umap1.insert({3,"3"});
        umap1.insert({4,"4"});
        umap1.insert({2,"2"});
        printMaps(1,map1, umap1);
        umap11.insert({1,"1"});
        umap11.insert({5,"5"});
        umap11.insert({3,"3"});
        umap11.insert({4,"4"});
        umap11.insert({2,"2"});
        printMaps(11,map1, umap11);

        std::map<int,std::string> map2 = {{2,"2"},{1,"1"},{0,"0"}};
        std::unordered_map<int, std::string> umap2 = {{2,"2"},{1,"1"},{0,"0"}};
        printMaps(2,map2, umap2);
        map2.insert({3,"3"});
        umap2.insert({3,"3"});
        printMaps(2,map2, umap2);

        std::map<std::string,std::string> map3 = {{"a","a"},{"2","2"},{"f","f"},{"F","F"},{"02","02"}};
        std::unordered_map<std::string, std::string> umap3 = {{"a","a"},{"2","2"},{"f","f"},{"F","F"},{"02","02"}};
        printMaps(3,map3, umap3);
        map3.insert({"3","3"});
        umap3.insert({"3","3"});
        printMaps(3,map3, umap3);

        std::map<std::string,std::string> map4 = {{"a","a"},{"2","2"},{"f","f"},{"F","F"},{"02","02"}};
        std::unordered_map<std::string, std::string> umap4 = {{"a","a"},{"2","2"},{"f","f"},{"F","F"},{"02","02"}};
        printMaps(4,map4, umap4);
        map4.insert({"3","3"});
        umap4.insert({"3","3"});
        map4.erase("3");
        umap4.erase("3");
        printMaps(4,map4, umap4);

}
    
    template<typename T> static void printMaps(int i, std::map<T,std::string> map, std::unordered_map<T, std::string> umap){
        std::cout << "map" << i << ": \n";
        for(auto x= map.begin(); x != map.end(); ++x){
            std::cout << x->first << ": ";
            std::cout << x->second << ", ";
        }
        std::cout << "\n";
        std::cout << "umap" << i << ": \n";
        for(auto x= umap.begin(); x != umap.end(); ++x){
            std::cout << x->first << ": ";
            std::cout << x->second << ", ";
        }
        std::cout << "\n\n";
        // std::cout << 
    }
};
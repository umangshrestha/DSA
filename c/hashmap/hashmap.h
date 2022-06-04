#ifndef __HASH_MAP_H
#define __HASH_MAP_H
// default size of array
#define SIZE 26 
// each array contains linked list


typedef struct HashMap {
    linkedlist data[SIZE];
} dict;

/* memory interpretaion
* data[0] => LinkedList[Node(k1,v1), Node(k2,v2), ...]
* data[1] => LinkedList[Node(k3,v3), Node(k4,v4), ...]
* ...
*/
#endif
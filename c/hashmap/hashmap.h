#ifndef __HASH_MAP_H
#define __HASH_MAP_H
// default size of array
#define SIZE 26 
#include "../linkedlist/linkedlist.h"
#include "../prime/prime.h"

// each array contains linked list

typedef struct HashMap hm;

struct HashMap {
    linkedlist data[SIZE];
};

/* memory interpretaion
* data[0] => LinkedList[Node(k1,v1), Node(k2,v2), ...]
* data[1] => LinkedList[Node(k3,v3), Node(k4,v4), ...]
* ...
*/
hm* hmNew();
void hmFree(hm *self);
void hmAdd(hm *self, int key, int value);
void hmRemove(hm *self, int key);
void hmPrint(hm *self);
int  hmIsEmpty(hm *self);
int  hmIsKey(hm *self, int key);
int  hmSize(hm *self);

int  getIndex(int key) {
    return getNextPrime(key) % SIZE;
}

#endif
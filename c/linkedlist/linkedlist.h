/*
This linked list implementation is specifically defined for hashmap a.k.a Dict.
The major changes are:
*  - Each node contains int value for key and pair
*  - If the key exist in linked list it will update the value
*/

#ifndef __LINKED_LIST_H
#define __LINKED_LIST_H

// node is the body of linked list
// each node points to next node in memory
// once next node is NULL then the node is terminated
typedef struct Node node;
struct Node {
    int key;
    int value;
    node *next;
};

// linked list the primary struct
// it has a head pointing to the node
typedef struct LinkedList linkedlist;
struct LinkedList {
    node *head;
};

linkedlist* llNew();
void llFree(linkedlist *self);
void llAdd(linkedlist *self, int key, int value);
void llRemove(linkedlist *self, int key);
void llPrint(linkedlist *self);
int  llIsEmpty(linkedlist *self);
int  llIsKey(linkedlist *self, int key);
int llSize(linkedlist *self);


#endif
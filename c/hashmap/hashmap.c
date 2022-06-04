#include<stdio.h>
#include<stdlib.h>
#define SIZE 10



typedef struct HashMap {
    linkedlist *data[SIZE];
} dict;



dict* NewDict() {
    // allocating memory for dict
    dict *d = (dict *) malloc(sizeof(dict));
    // allocating memory for linked list inside dict
    for(int i=0; i<SIZE; i++) 
        d->data[i] = NewLinkedList();
}


void removeDict(dict *d) {
    // removinf data in linked list
    for (int i=0; i< SIZE; i++) 
        //going to end of list
        freeLinkedList(d->data[i]);
        
    // removing dict
    free(d);
    d = NULL;
}


void addToDict(dict *d, int key, int value) {
    int i = getHashMod(key);
    node *n = (node *) malloc(sizeof(n));
    n->key = key;
    n->value = value;

    if (d->data[i]->head == NULL) {
        d->data[i]->head = n; 
    } 
}


void printDict(dict *d) {
    printf("=============");
    for (int i=0; i<SIZE; i++) {
        printf("[%d]", i);
        printLinkedList(d->data[i]);
    }
}


int main() 
{
    linkedlist *l = NewLinkedList();
    addToLinkedList(l, 1, 2);
    addToLinkedList(l, 2, 2);
    addToLinkedList(l, 2, 21);
    printLinkedList(l);
    freeLinkedList(l);
    // removeDict(d);


    return 0;
}
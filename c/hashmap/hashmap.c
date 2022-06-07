#include <stdio.h>
#include <stdlib.h>
#include "hashmap.h"


hm* hmNew() {
    // allocating memory for hm
    hm *d = (hm*) malloc(sizeof(hm));
    // allocating memory for linked list inside hm
}

void hmFree(hm *d) {
    // removinf data in linked list
    for (int i=0; i< SIZE; i++) 
        llFree(&d->data[i]);
    // removing hm
    d = NULL;
}

int hmIsEmpty(hm *d) {
    for(int i=0; i<SIZE; i++)  {
       if (!llIsEmpty(&d->data[i]))
            return 0;
    }
    return 1;
}

int hmSize(hm *d) {
    int s = 0;
     for(int i=0; i<SIZE; i++)  {
       s += llSize(&d->data[i]);
    }
    return s;

}

void hmAdd(hm *d, int key, int value) {
    int i = getIndex(key);
    llAdd(&d->data[i], key, value);  
}

void hmRemove(hm *d, int key) {
    int i = getIndex(key);
    llRemove(&d->data[i], key);  
}



void hmPrint(hm *d) {
    for (int i=0; i<SIZE; i++) {
        printf("[%d]", i);
        llPrint(&d->data[i]);
    }
}


int main() 
{
    hm* d = hmNew();

    // for(int i=0; i<SIZE; i++) { 
    //     int a = llIsEmpty(d->data[i]);
    //     p   d->data = (linkedlist**) malloc(sizeof(linkedlist) * SIZE);
    for(int i=0; i<2*SIZE; i++) {
        hmAdd(d, i, i);
        printf("[%d]:=>%d\n", i, hmSize(d));
    }
    hmRemove(d, 1);
    hmPrint(d);

    // hmAdd(h, 1,2);
    // hmPrint(h);
    
}
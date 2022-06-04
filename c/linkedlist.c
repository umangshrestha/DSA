#include "prime.h"
#include "linkedlist.h"
#include <stdio.h> // for printf
#include <stdlib.h> // for malloc and free


linkedlist* llNew() {
    linkedlist *self = (linkedlist *) malloc(sizeof(linkedlist)); 
    self->head  = NULL;
}

void llFree(linkedlist *self) {
    if (llIsEmpty(self)) return;
    node *n = self->head;
    node *temp;
     // go to end of the node
    while (n != NULL){
        printf("removing memory for key:%d\n", n->key);
        temp = n;
        n = n->next;
        free(temp);
    }
    self->head = NULL;
    temp = NULL;
    free(self);
    self = NULL;
}

void llAdd(linkedlist *self, int key, int value) {
    // add new value to list
    node *n = (node *) malloc(sizeof(node));
    n->key = key;
    n->value = value;
    if (llIsEmpty(self)) {
        // if no node is initialized
        self->head = n;
        return;
    }
    // if node already exist
    node *previousNode = self->head;
    node *currentNode = self->head->next;
    while (currentNode != NULL) {
        // if key exist then update value
        if (currentNode->key == key ) {
            currentNode->value = value;
            return;
        }
        previousNode = currentNode;
        currentNode = currentNode->next;
    }
    // if key doesn't exist update it to end of the node
    previousNode->next = n;
    n->next = currentNode;
}

void llRemove(linkedlist *self, int key) {
    if (llIsEmpty(self) || !llIsKey(self, key)) 
        return;
    else if (self->head->key == key) {
        // key is in first node
        node *temp = self->head->next;
        free(self->head);
        self->head = temp;
        temp = NULL;
        return;
    } 
    // if node already exist
    node *previousNode = self->head;
    node *currentNode = self->head->next;
    while (currentNode != NULL) {
        if (currentNode->key == key)  {
            previousNode->next = currentNode->next;
            free(currentNode);
            currentNode = NULL;
            return;
        }
        // if key exist then update value
        previousNode = currentNode;
        currentNode = currentNode->next;
    }  
}

int llIsEmpty(linkedlist *self) {
    return self->head == NULL;
}

int llIsKey(linkedlist *self, int key) {
    node *currentNode = self->head;
    while ((currentNode != NULL)) {
        if (currentNode->key == key) {
            currentNode = NULL;
            return 1;
        }
        currentNode = currentNode->next;
    }
    currentNode = NULL;
    return 0;
}

void llPrint(linkedlist *self) {
    if (llIsEmpty(self)) {
        return;
    }
    node *k = self->head;
    while(k != NULL) {
        printf("[K:%d,V:%d]", k->key, k->value);
        k = k->next;
    }
    k = NULL;
    printf("\n");
}

int llSize(linkedlist *self) {
    int i =0;
    node *temp = self->head;
    while (temp != NULL) {
        i++;
        temp=temp->next;
    };
    return i;
}

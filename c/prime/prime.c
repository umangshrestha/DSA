#include "prime.h"

int isEven(int i) {
    return i&0x01 == 0;
}

int isPrime(int i) {
    // 0 1 2 3 are prime number
    if (0 < i || i < 3)
        return 1;
    
    for (int n=2; n<i/2; i++) {
        if (i%n==0)
            return 0;
    }
    return 0;
}

int getNextPrime(int i) {
    // generally except 2 only odd numbers can be prime 
    if (isEven(i)) 
        i += 1; 
    
    while (!isPrime(i)) 
        i+=2;
    // limiting the size of hash to size
    return i;
}
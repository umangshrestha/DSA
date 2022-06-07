############################
# Settings
############################
CC=cc
RM=rm
BIN=bin
HM_OUT=${BIN}/hm
FLAGS=-g
############################
# Files
############################
HM_C=hashmap/hashmap.c
LL_C=linkedlist/linkedlist.c
PRIME_C=prime/prime.c
############################
hm: ${HM_C} ${LL_C} ${PRIME_C}
############################
	${CC} ${HM_C} ${LL_C} ${PRIME_C} -o ${HM_OUT} ${FLAGS}
############################
rm:
############################
	${RM} -r ${BIN}/*
############################
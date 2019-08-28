#!/bin/csh

rm *.o
gcc -c *.c
#ar -rv libx17r.a *.o

#gcc main.c -I./ -static -L./ -lx17r -o main.exe

gcc -o main.exe *.o

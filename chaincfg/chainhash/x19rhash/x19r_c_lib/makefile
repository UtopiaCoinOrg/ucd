SRCS =blake.c blake2b.c blake2s.c bmw.c cubehash.c echo.c fugue.c groestl.c hamsi.c haval.c jh.c keccak.c luffa.c ripemd.c sha2big.c shabal.c shavite.c simd.c skein.c sph_sha2.c streebog.c whirlpool.c tiger.c x19r.c
#OBJS = $(SRCS:.c= .o)
OBJS = blake.o blake2b.o blake2s.o bmw.o cubehash.o echo.o fugue.o groestl.o hamsi.o haval.o jh.o keccak.o luffa.o ripemd.o sha2big.o shabal.o shavite.o simd.o skein.o sph_sha2.o streebog.o whirlpool.o tiger.o x19r.o

CC = gcc
#CC = arm-linux-gcc #for android
CCFLAGS =

a : $(OBJS)
	ar -rv libx19r.a  $(OBJS)

%.o : %.c
	$(CC) -c $< $(CCFLAGS)

clean:
	rm -rf test *.o
	rm -rf libx19r.a
	rm -rf *.exe
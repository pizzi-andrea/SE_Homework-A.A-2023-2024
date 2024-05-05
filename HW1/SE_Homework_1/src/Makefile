

#IDIR=../hdr
ODIR=../obj
BDIR=../bin

CC=gcc
CFLAGS=-z execstack -O0 -fno-stack-protector 

SRC = $(wildcard *.c)
OBJ = $(patsubst %.c, $(ODIR)/%.o, $(SRC))

all:	$(BDIR)/main

$(ODIR)/%.o: %.c
	$(CC) -c -o $@ $< $(CFLAGS)

$(BDIR)/main: $(OBJ)
	$(CC) -o $@ $(OBJ) $(CFLAGS)




.PHONY: clean

clean:
	rm -f $(ODIR)/*.o *~ $(BDIR)/*~ core


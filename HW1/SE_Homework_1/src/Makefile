
# directory di compilazione
ODIR=../obj
BDIR=../bin

# compilatore e flags di compilazione 
CC=gcc
CFLAGS=-z execstack -O0 -fno-stack-protector -no-pie -z norelro

# lista dei file da compilare
SRC = $(wildcard *.c)
OBJ = $(patsubst %.c, $(ODIR)/%.o, $(SRC))

# ruoli
all:	$(BDIR)/main

$(ODIR)/%.o: %.c
	$(CC) -c -o $@ $< $(CFLAGS)

$(BDIR)/main: $(OBJ)
	$(CC) -o $@ $(OBJ) $(CFLAGS)




.PHONY: clean

clean:
	rm -f $(ODIR)/*.o *~ $(BDIR)/*~ core


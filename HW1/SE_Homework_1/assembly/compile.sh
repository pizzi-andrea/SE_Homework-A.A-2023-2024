#!/bin/bash

# Compilazione ed linksing del codice

nasm -f elf64 $1 -o out.o   # Compila il codice assemply senza linking
ld out.o -o out             # Effettua il linking
rm out.o                    # rimuove il file oggetto

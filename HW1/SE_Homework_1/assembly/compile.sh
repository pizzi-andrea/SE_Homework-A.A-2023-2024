nasm -f elf64 $1 -o out.o
ld out.o -o out
rm out.o

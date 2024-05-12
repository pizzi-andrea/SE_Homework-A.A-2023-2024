nasm -f elf64 $1 -o out.o
ld out.o -o ./bin/out
rm out.o

## Step 1 .: disabilitare le misure di sicurezza standard

+ `sysctl kernel/randomize_va_space` 

## Step 1: Ottenere il valore del registro `RIP` (`EIP` in x32)



1. Per analizzare il codice vulnerabile utilizzare gdb-peda
2. Generare un file di traboccamento con `pattern create <size> <file>` 
	msf-pattern_create -l 1000 > trash.txt



3. Verificare l'offset con gdb-patten `pattern offset <pattern>` 
	valore offset `121`. Il valore di ritorno si trova in posizione 121 dall'inizio dello stack frame. Da 121 in poi si può inserire lo shellcode



4. Calcolo il numpero di NOP da inserire prima della shell code (offset - dimensione_shell code)
5. Posizionari prima dello shell code l'indirizzo da inserire nel RIP per eseguire il 'salto' al codice della shellcode
[](https://www.vividmachines.com/shellcode/shellcode.html#linex3)

nasm -f elf shellex.asm
ld -o shellex shellex.o
objdump -d shellex

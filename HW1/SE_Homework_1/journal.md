## Step 1 .: disabilitare le misure di sicurezza standard

+ `sysctl kernel/randomize_va_space` 

## Step 1: Ottenere il valore del registro `RIP` (`EIP` in x32)



1. Per analizzare il codice vulnerabile utilizzare gdb-peda
2. Generare un file di traboccamento con `pattern create <size> <file>` 
	msf-pattern_create -l 4000 > trash.txt 



3. Verificare l'offset con gdb-patten `pattern offset <pattern>` 
	valore offset `249`. Il valore di ritorno si trova in posizione 249 dall'inizio dello stack frame. Da 249 in poi si può inserire lo shellcode





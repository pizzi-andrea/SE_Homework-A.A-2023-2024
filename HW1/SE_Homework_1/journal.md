## Step 1 .: disabilitare le misure di sicurezza standard

+ `sysctl kernel/randomize_va_space` 

## Step 1: Ottenere il valore del registro `RSP` (`EIP` in x32)



1. Per analizzare il codice vulnerabile utilizzare gdb-peda
2. Generare un file di traboccamento con il frame work metaexploit

Inserendo i valori direttamente generati non succede niente poichè i valori ottenuto sono considerati indirizzi non validi in una architettura a 64bit (un istruzione può usare al massimo 48bit) quindi il RIP non viene sovrascritto.

E' necessario mascherare i dati in modo che sembrino degli indirizzo canonici validi a 48 bit.



3. Verificare l'offset con metaexploit (``)


msf-pattern_create -l 4000 > trash.txt 

msf-pattern_offset
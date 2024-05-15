## Step 1 .: disabilitare le misure di sicurezza standard

Disabilitazzione della randomizzazione degli indirizzi (per il sistema host)
 `sysctl kernel/randomize_va_space` 

Disabilitare i meccanismi di protezione in fase di compilazione

`-z execstack`: permetti l'esecuzione di codice arbitario nello stack
`-fno-stack-protector`: diabilita i meccanismi di protezione dello stack


## Step 2: Calcolare l'offset dello stack (posizione dello stack pointer `RIP`/`EIP` in x32)

Per l'analisi viene utilizzata un estensione del tradizionale debbuger C gdb con funzionalità avanzate di exploit **gdb-peda**

1. Tramite il comando del debbuger `pattern create` viene generato un file contenente una sequenza di valori generati secondo un pattern
di dimensione maggiore di quella prevista del buffer

2. Si imposta il debbugger per tracciare l'esecuzione del programma target 

3. Si immette come input al programma infetto il file pattern generato (vedere pattern.txt)

4. Tramite il comando `pattern search` viene estratto offset dello stack pointer (nel caso corrente pari a 152 byte)

## Step 3: Assumere il controllo del RIP 

Per testare il controllo del contenuto del registro RIP sono stati scritti script python *(versione python3)* 
che generano la stringa da somministrare. 

La prima versione di script python, *rip.py* e' un semplice script che genera partendo dal valore dell'offset
calcolato al passo precedente una stringa per scrivere all'interno del rip un valore casuale
(massimo 6 byte per limitazioni sull'indirizzamento effettivo)

## Step 4: Preparare lo shellcode

## Step 5: Preparare la stringa malevola

## Step 6: Passare la stringa in input al programma senza incovenienti di codifica

## Step 7: Lancia una shell con permessi di root

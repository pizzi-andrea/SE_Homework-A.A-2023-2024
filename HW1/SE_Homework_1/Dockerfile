# Compile program with old gcc version without security fectures
FROM docker.io/gcc:6.2 AS compiler

WORKDIR /homework/

COPY . .

RUN make --directory=src clean
RUN make --directory=src

FROM debian:latest

# isntall useful tools
RUN apt update && apt install git python3 -y
RUN apt update && apt install gdb -y
RUN apt install nano -y
# create not-root user (threads)
RUN useradd -ms /bin/bash user_01
USER user_01
WORKDIR /home/user_01

RUN git clone https://github.com/longld/peda.git ~/peda; echo "source ~/peda/peda.py" >> ~/.gdbinit;  
COPY --from=compiler ./homework/bin ./homework/data .










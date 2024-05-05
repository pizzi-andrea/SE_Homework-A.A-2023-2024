# Compile program with old gcc version without security fectures
FROM gcc:6.2 AS compiler

WORKDIR /homework/

COPY . .

RUN make --directory=src clean
RUN make --directory=src

FROM debian:latest

WORKDIR /homework/
COPY --from=compiler ./homework/bin .

# isntall useful tools
RUN apt update && apt install git python3 -y
RUN git clone https://github.com/longld/peda.git ~/peda; echo "source ~/peda/peda.py" >> ~/.gdbinit;  
RUN apt update && apt install gdb -y






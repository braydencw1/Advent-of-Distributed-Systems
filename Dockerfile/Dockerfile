FROM openjdk:23-oraclelinux7


RUN yum upgrade -y && \
    yum install graphviz gnuplot wget bzip2 git glibc -y
WORKDIR /app


RUN wget -O maelstrom.tar.bz2 https://github.com/jepsen-io/maelstrom/releases/download/v0.2.3/maelstrom.tar.bz2 && \
    tar -xvjf maelstrom.tar.bz2 && \ 
    rm maelstrom.tar.bz2

CMD ["./maelstrom/maelstrom", "test", "-w", "echo", "--bin", "/cmd/maelstrom-echo", "--node-count", "1", "--time-limit", "10"]

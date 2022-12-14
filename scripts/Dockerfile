FROM --platform=linux/amd64 python:3.8.12-buster

RUN apt-get update && apt-get install -y git wget
RUN wget https://go.dev/dl/go1.18.9.linux-amd64.tar.gz -q
RUN tar -C /usr/local -xzf go1.18.9.linux-amd64.tar.gz

ENV PATH=${PATH}:/usr/local/go/bin

RUN git clone --recurse-submodules https://github.com/COVESA/vehicle_signal_specification

WORKDIR /vehicle_signal_specification

RUN make install && pip install graphql-core==3.2.3 pytest && make all

RUN mkdir /out && cp vss_rel_$(cat VERSION).* /out/
RUN echo "vss_rel_$(cat VERSION).yaml" > /out/filename

WORKDIR /out/
COPY . .

RUN go mod tidy
RUN go run main.go $(cat filename)
RUN rm filename 

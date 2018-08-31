all: build

build: udp_server
udp_server:
	go build -o bnw_udp
clean:

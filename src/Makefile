CFLAGS=-std=c++11 -Wno-unknown-pragmas -Wall

# Added for mxnet, change to yourself's mxnet root
export MXNET_ROOT=/root/MXNet/mxnet

CFLAGS+= -I$(MXNET_ROOT)/include
LDFLAGS+=$(MXNET_ROOT)/lib/libmxnet.so

libmxnetwrap.so: c_api_wrap.o
	g++ -fPIC --shared -o libmxnetwrap.so c_api_wrap.o $(LDFLAGS)
	ln -s $(GOPATH)/src/github.com/songtianyi/go-mxnet/src/libmxnetwrap.so /usr/lib/libmxnetwrap.so
c_api_wrap.o: c_api_wrap.cxx
	g++ -fPIC -c c_api_wrap.cxx $(CFLAGS)
clean:
	rm -f *.o *.so /usr/lib/libmxnetwrap.so

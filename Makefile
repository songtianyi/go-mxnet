MODULE_NAME=mxnet
build_wrapper:
	swig -go -gccgo -c++ -intgosize 64 cpp2go.swig
clean:
	rm -f  $(MODULE_NAME).go *_wrap.cxx
	

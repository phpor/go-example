swig -intgosize 64 -c++ -cgo -go caid.i

g++ -c -fpic -I . -L . -l caid caid_alg.h  caid_wrap.cxx

ar -r libcaid_wrap.a caid_wrap.o



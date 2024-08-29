	
# 获取系统类型
UNAME_S := $(shell uname -s)

# 根据系统类型设置不同的变量或执行不同的命令
ifeq ($(UNAME_S),Linux)
    PLATFORM = LINUX
    # Linux特定的编译选项或者规则可以放在这里
endif

ifeq ($(UNAME_S),Darwin)
    PLATFORM = MACOS
    # macOS特定的编译选项或者规则可以放在这里
endif


build:
ifeq ($(PLATFORM),LINUX)
	cc -g -o say_linux.o -c say.c && ar rcs libsay_linux.a say_linux.o
else ifeq ($(PLATFORM),MACOS)
	cc -g -o say_darwin.o -c say.c && ar rcs libsay_darwin.a say_darwin.o
endif
	cp libsay_* ../sdk/

clean:
	rm -rf *.o *.a
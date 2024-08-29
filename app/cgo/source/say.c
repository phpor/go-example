#include <stdio.h>


int say(void) {
    printf("haha\n"); // 加个换行比较好，否则可能不会输出的哦，没有换行的话，mac上能输出，linux上不输出
//    fflush(stdout); // 手动刷新stdout缓冲区,没有换行的话，fflush也不行
    return 1;
}
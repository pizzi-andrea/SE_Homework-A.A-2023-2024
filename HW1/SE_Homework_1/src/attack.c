#include <stdlib.h>
#include <stdio.h>
#include <string.h>

int overflow(const char *str)
{
    char buffer[256];
    memcpy(buffer, str, 1000);
    return EXIT_SUCCESS;
}

int main(int argc, char const *argv[])
{
    overflow(argv[1]);
    printf("%s", "Hello World");
    return EXIT_SUCCESS;
}



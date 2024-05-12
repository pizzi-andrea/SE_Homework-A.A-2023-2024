#include "exploit.h"
#include "bad.h"

/*
 * @brief Main del programma
*/
int main(int argc, char const *argv[])
{
    overflow(argv[1]);
    return EXIT_SUCCESS;
}

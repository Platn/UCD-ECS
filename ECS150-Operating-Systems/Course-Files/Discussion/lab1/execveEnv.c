#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main(void) {
    char *cmd[] = { "ls", "-l", NULL };
    char *env[] = { "HOME=/usr/home", NULL };
    execve ("/bin/ls", cmd, env);
    return 0;
}
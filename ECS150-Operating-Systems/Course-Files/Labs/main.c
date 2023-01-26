#include <sys/types.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>
#include <string.h>

int main(int argc, char *argv[]) {
    char userInput[256];
    scanf("%s", userInput);
    printf("%s",userInput);

    pid_t pid = fork();

    

    if(pid < 0) {
        // Print error, might be perror(), or fprintf
    }

    else if (pid > 0) {
        // Child Process
        wait(NULL);
    }

    else {
        // Child process
        
    }

    

    return 0;
}
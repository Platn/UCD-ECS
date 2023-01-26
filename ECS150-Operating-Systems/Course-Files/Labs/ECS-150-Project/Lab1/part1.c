#define _POSIX_C_SOURCE  200809L
#include <stdio.h>
#include <unistd.h>
#include <signal.h>
#include <time.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>

void catch(int snum) {
  int pid;
  int status;
  pid = wait(&status);
  kill(pid, SIGTERM);
}

char isNumber(char *val) {
  for(int i = strlen(val) - 1; i >= 0; i--) {
    if(val[i] >= '0' && val[i] <= '9')
            continue;
    return 0;
  }
  return 1;

  // https://stackoverflow.com/questions/16644906/how-to-check-if-a-string-is-a-number
}

void command_handler(int argc, char *argv[], char *envp[]) {
  char *command = strdup(argv[2]);
  int cArgC = argc - 3;  // Number of command args
  char *myargs[cArgC + 2];
  myargs[0] = command;
  
  for (int i = 1; i < cArgC + 1; i++) {
    myargs[i] = strdup(argv[i + 2]);
  }
  myargs[cArgC + 1] = NULL;  // end of arr
  // printf("%s", envp);
  // ./main.o 6 ./loop.o /home/runner/TimeOut-Practice/
  // ./main.o 6 /bin/ls
  if(execve(command, myargs, envp) == -1) {
    perror("execve");
  }
}

int main(int argc, char*argv[], char*envp[]){
  if(argc < 3) {
    printf("Usage: timeout sec command [args ...]\n");
    return 0;
  }
  
  else if(isNumber(argv[1]) == 0) {
    printf("%s is not a positive integer\n", argv[1]);
    return 0;
  }

  sigset_t  mask;
  siginfo_t info;

  // printf("Arg1: %s", argv[0]);
  // printf("Arg2: %s", argv[1]);
  // printf("Arg3: %s", argv[2]);

  sigemptyset(&mask);
  signal(SIGCHLD, catch);
  
  pid_t newProc = fork();
  int signum;

  if (newProc < 0) {
    printf("Fork Failed");
      return 1;
  }

  else if (newProc == 0) {
    command_handler(argc, argv, envp);
    return EXIT_SUCCESS;
  }

  else {
    struct timespec t = {0};
    long int time = atoi(argv[1]);
    t.tv_sec = time;

    signum = sigtimedwait(&mask, &info, &t);

    if(signum != SIGCHLD) {
      kill(newProc, SIGTERM);
    }  
  }
  return 0;
}
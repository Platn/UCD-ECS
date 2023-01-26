#include <stdio.h>
#include <unistd.h>
#include <signal.h>
#include <time.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>
#include <dirent.h>

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

int hasSlash(char* arg2) {
  int pos = 0;
  while(arg2[pos] != '\0') {
    if(arg2[pos] == '/') {
      return 1;
    }
    pos++;
  }

  return 0;
}

int searchDir(char* nextDir, char* arg) {
  char *dir = nextDir;
  struct dirent *d;
	DIR *dh = opendir(dir);
	if (!dh) {
		if (errno = ENOENT) {
			perror("Directory doesn't exist");
		}
		else {
			perror("Unable to read directory");
		}
		// exit(EXIT_FAILURE);
    return 0;
	}

	while ((d = readdir(dh)) != NULL) {
		if(strcmp(d->d_name, arg) == 0) {

      return 1;
    }
	}
  return 0;
}

// https://iq.opengenus.org/ls-command-in-c/

char* findFile(char *arg2) {
  char *filePath = malloc(sizeof(char) * 256);
  const char* path = getenv("PATH");
  char* pathParts[100];

  char* token;
  char src[50];
  char dest[50];
  strcpy(src, arg2);
  strcpy(dest,"/");
  strcat(dest,src);

  char delimit[] = ":\0";

  token = strtok((char*)path, delimit);

  pathParts[0] = token;
  int i = 1;

  while(token != NULL) {
    token = strtok(NULL, delimit);
    if(token == NULL) {
      break;
    }

    pathParts[i] = token;
    i++;
  }

  token = NULL;
  free(token);
  path = NULL;
  free((char*)path);
  int matched = 0;

  for(int j = 0; j < i; j++) {
    matched = searchDir(pathParts[j], arg2);
    if(matched == 1) {
      strcpy(filePath,pathParts[j]);
      strcat(filePath, dest);
      break;
    }
  }

  if(matched == 0) {
    return NULL;
  }

  free((char*)path);
  return filePath;
}

void command_handler(int argc, char *argv[], char *envp[], char* fileName) {

  char *command = fileName != NULL ? fileName : strdup(argv[2]);
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

  /*if(fileName != NULL) {
    execve(fileName, myargs, NULL);
  }

  else {
    for (int i = 1; i < cArgC + 1; i++) {
      myargs[i] = strdup(argv[i + 2]);
    }
    myargs[cArgC + 1] = NULL;  // end of arr

    if(execve(command, myargs, envp) == -1) {
      perror("execve");
    }
  }*/

  // https://stackoverflow.com/questions/26597977/split-string-with-multiple-delimiters-using-strtok-in-c
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
  char *file = NULL;
  if(!hasSlash(argv[2])) {
    file = findFile(argv[2]);
  }

  sigset_t  mask;
  siginfo_t info;

  sigemptyset(&mask);
  signal(SIGCHLD, catch);

  pid_t newProc = fork();
  int signum;

  if (newProc < 0) {
    printf("Fork Failed");
      return 1;
  }

  else if (newProc == 0) {
    command_handler(argc, argv, envp, file);
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

  free(file);
  return 0;
}

// SIG CHILD
// https://stackoverflow.com/questions/21762208/notify-parent-process-when-child-process-dies#:~:text=In%20case%20you%20only%20want,spawned%20by%20this%20process%20dies.

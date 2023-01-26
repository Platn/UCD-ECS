#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <signal.h>
#include <time.h>

// TODO
// Format output

// TODO: Find a better way to do this
char* signal_list[] = {"SIGHUP",  "SIGINT",   "SIGQUIT", "SIGILL",  "SIGTRAP",
                       "SIGABRT", "SIGBUS",   "SIGFPE",  "SIGKILL", "SIGBUS",
                       "SIGSEGV", "SIGSYS",   "SIGALRM", "SIGTERM", "SIGUSR1",
                       "SIGUSR2", "SIGCHLD",  "SIGCONT", "SIGTSTP", "SIGTTIN",
                       "SIGTTOU", "SIGSTOP",  "SIGXCPU", "SIGXFSZ", "SIGVTALRM",
                       "SIGPROF", "SIGWINCH", "SIGPOLL", "SIGUSR1", "SIGUSR2"};

static void sig_handler(int signo) {
  time_t current_time = time(NULL);
  const char* time_string = strtok(ctime(&current_time), "\n");
  printf("%s Recieved signal %d (%s)\n", time_string, signo,
         signal_list[signo - 1]);
}

int main(int argc, char* argv[]) {
  printf("%d\n", (int)getpid());

  for (int i = 1; i <= 31; i++)
    if (signal(i, sig_handler) == SIG_ERR) printf("Can't catch SIG_ERR\n");

  while (1) sleep(1);
  return 0;
}

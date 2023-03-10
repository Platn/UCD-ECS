#define _POSIX_C_SOURCE  200809L
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <signal.h>
#include <string.h>
#include <stdio.h>
#include <errno.h>

static inline const char *signal_name(const int signum)
{
    switch (signum) {
    case SIGINT:  return "SIGINT";
    case SIGHUP:  return "SIGHUP";
    case SIGTERM: return "SIGTERM";
    case SIGQUIT: return "SIGQUIT";
    case SIGUSR1: return "SIGUSR1";
    case SIGUSR2: return "SIGUSR2";
    default:      return "(unnamed)";
    }    
}

int main(void)
{
    sigset_t  mask;
    siginfo_t info;
    pid_t     child, p;
    int       signum;    

    sigemptyset(&mask);
    sigaddset(&mask, SIGINT);
    sigaddset(&mask, SIGHUP);
    sigaddset(&mask, SIGTERM);
    sigaddset(&mask, SIGQUIT);
    sigaddset(&mask, SIGUSR1);
    sigaddset(&mask, SIGUSR2);
    if (sigprocmask(SIG_BLOCK, &mask, NULL) == -1) {
        fprintf(stderr, "Cannot block SIGUSR1: %s.\n", strerror(errno));
        return EXIT_FAILURE;
    }

    child = fork();
    if (child == -1) {
        fprintf(stderr, "Cannot fork a child process: %s.\n", strerror(errno));
        return EXIT_FAILURE;
    } 
    
    else if (!child) {
        /* This is the child process. */
        printf("Child process %d sleeping for 3 seconds ...\n",         (int)getpid());
        fflush(stdout);
        sleep(3);

        printf("Child process %d sending SIGUSR1 to parent process (%d) ...\n", (int)getpid(), (int)getppid());
        fflush(stdout);
        kill(getppid(), SIGUSR1);

        printf("Child process %d exiting.\n", (int)getpid());
        return EXIT_SUCCESS;
    }

    /* This is the parent process. */
    printf("Parent process %d is waiting for signals.\n", (int)getpid());
    fflush(stdout);

    while (1) {

        signum = sigtimedwait(&mask, &info, 6);
        if (signum == -1) {

            /* If some other signal was delivered to a handler installed
               without SA_RESTART in sigaction flags, it will interrupt
               slow calls like sigwaitinfo() with EINTR error. So, those
               are not really errors. */
            if (errno == EINTR)
                continue;

            printf("Parent process: sigwaitinfo() failed: %s.\n", strerror(errno));
            return EXIT_FAILURE;
        }

        if (info.si_pid == child)
            printf("Parent process: Received signal %d (%s) from child process %d.\n", signum, signal_name(signum), (int)child);
        else
        if (info.si_pid)
            printf("Parent process: Received signal %d (%s) from process %d.\n", signum, signal_name(signum), (int)info.si_pid);
        else
            printf("Parent process: Received signal %d (%s).\n", signum, signal_name(signum));
        fflush(stdout);

        /* Exit when SIGUSR1 received from child process. */
        if (signum == SIGUSR1 && info.si_pid == child) {
            printf("Parent process: Received SIGUSR1 from child.\n");
            break;
        }

        /* Also exit if Ctrl+C pressed in terminal (SIGINT). */
        if (signum == SIGINT && !info.si_pid) {
            printf("Parent process: Ctrl+C pressed.\n");
            break;
        }
    }

    printf("Reaping child process...\n");
    fflush(stdout);

    do {
        p = waitpid(child, NULL, 0);
        if (p == -1) {
            if (errno == EINTR)
                continue;
            printf("Parent process: waitpid() failed: %s.\n", strerror(errno));
            return EXIT_FAILURE;
        }
    } while (p != child);

    printf("Done.\n");
    return EXIT_SUCCESS;
}
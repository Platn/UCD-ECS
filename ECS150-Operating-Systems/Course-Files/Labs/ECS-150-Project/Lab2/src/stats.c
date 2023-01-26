#include <stdio.h>

typedef struct {
  int finished_time;
} Sys_Stats;

typedef struct {
  int busy_time;
  int idle_time;
  int cpu_dispatches;
} CPU_Stats;

typedef struct {
  int busy_time;
  int idle_time;
  int io_dispatches;
} IO_Stats;

Sys_Stats Sys_stats = {0};
CPU_Stats CPU_stats = {0,0,0};
IO_Stats IO_stats = {0,0,0};

void Print_All_Job_Stats(Node* jobs)
{
  printf("Processes:\n");
  printf("Name	CPU time	When done	# Dispatches	# Blocked for I/O	I/O time\n");
  Node* job = jobs;
  do{
    printf("\n");
    Print_Single_Job_Stats(job->data);
    job = job->next;
  } while(job);
}

void Print_Single_Job_Stats(Job* js)
{
  printf("%s	%d %d	%d	%d	%d\n", js->name, js->runtime, js->stats.completion_time, js->stats.cpu_count, js->stats.blocked_count, js->stats.io_time);
}

void Print_Sys_Stats()
{
  printf("System:\n");
  printf("The wall clock time at which the simulation finished: %d\n",Sys_stats.finished_time);
}

void Print_CPU_Stats()
{
  printf("CPU:\n");
  printf("Total time spent busy: %d\n",CPU_stats.busy_time);
  printf("Total time spent idle: %d\n",CPU_stats.idle_time);
  printf("CPU utilization: %.2f\n",CPU_stats.busy_time/(CPU_stats.busy_time+(float)CPU_stats.idle_time));
  printf("Number of dispatches: %d\n",CPU_stats.cpu_dispatches);
  printf("Overall throughput: %.2f\n",CPU_stats.cpu_dispatches/(CPU_stats.busy_time+(float)CPU_stats.idle_time));
}

void Print_IO_Stats()
{
  printf("I/O Device:\n");
  printf("Total time spent busy: %d\n",IO_stats.busy_time);
  printf("Total time spent idle: %d\n",IO_stats.idle_time);
  printf("I/O device utilization: %.2f\n",IO_stats.busy_time/(IO_stats.busy_time+(float)IO_stats.idle_time));
  printf("Number of time I/O was started: %d\n",IO_stats.io_dispatches);
  printf("Overall throughput: %.2f\n",IO_stats.io_dispatches/(IO_stats.busy_time+(float)IO_stats.idle_time));
}

void Print_Stats(Node* jobs)
{
  Print_All_Job_Stats(jobs);
  Print_Sys_Stats();
  printf("\n");
  Print_CPU_Stats();
  printf("\n");
  Print_IO_Stats();
}

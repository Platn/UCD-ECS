#include <stdlib.h>
#include <stdio.h>
#include <string.h>

typedef enum { LOW, MEDIUM, HIGH } priority;

typedef struct {
  int completion_time;
  int count;
  int blocked_count;
  int io_time;
  int cpu_count;
} Job_Stats;

Job_Stats* Job_Stats_New() {
  Job_Stats* newStats = malloc(sizeof(Job_Stats));
  newStats->blocked_count = 0;
  newStats->completion_time = 0;
  newStats->cpu_count = 0;
  newStats->io_time = 0;
  return newStats;
}

typedef struct {
  const char* name;
  unsigned int time_elapsed;
  unsigned int runtime;
  float blocking_probablity;
  priority prior;
  Job_Stats* stats;
} Job;

typedef struct Node {
  Job* data;
  struct Node* next;
} Node;

typedef struct {
  int finished_time;
} Sys_Stats;

Sys_Stats* Sys_Stats_New() {
  Sys_Stats* newSysStats = malloc(sizeof(Sys_Stats));
  newSysStats->finished_time = 0;
  return newSysStats;
}

typedef struct {
  int busy_time;
  int idle_time;
  int cpu_dispatches;
} CPU_Stats;

CPU_Stats* CPU_Stats_New() {
  CPU_Stats* newCPUStats = malloc(sizeof(CPU_Stats));
  newCPUStats->busy_time = 0;
  newCPUStats->cpu_dispatches = 0;
  newCPUStats->idle_time = 0;
  return newCPUStats;
}

typedef struct {
  int busy_time;
  int idle_time;
  int io_dispatches;
} IO_Stats;

IO_Stats* IO_Stats_New() {
  IO_Stats* newIOStats = malloc(sizeof(IO_Stats));
  newIOStats->busy_time = 0;
  newIOStats->idle_time = 0;
  newIOStats->io_dispatches = 0;
  return newIOStats;
}

/* Contructor */
Job* Job_New(char* name, unsigned int runtime, float blocking_probablity,
             priority prior) {
  Job* job = malloc(sizeof(Job));
  job->name = malloc(sizeof(char) * (strlen(name) + 1));
  strcpy(job->name, name);
  job->runtime = runtime;
  job->blocking_probablity = blocking_probablity;
  job->time_elapsed = 0;
  job->prior = prior;
  Job_Stats* js = Job_Stats_New();
  job->stats = js;

  return job;
}

int GetRemainingTime(Job* job) { return job->runtime - job->time_elapsed; }

/*int CheckIfFinished(Job* job) { return job->time_elapsed >= job->runtime; }*/

Node* Sort_Jobs_Finish_First(Node* jobs) {
  // We are sorting by going through the list and swapping data points
  // We just create nodes that move around the data? So basically 
  Node* head = jobs; // Keep this the same?
  Node* traverser;
  traverser = head;
  int swapped = 0;
  while(!swapped) {
    while(traverser != NULL){
      swapped = 1;
      if(traverser->next != NULL && 
        traverser->data->stats->completion_time > traverser->next->data->stats->completion_time){ // Compare data
        Job* dataHold = traverser->data;
        traverser->data = traverser->next->data;
        traverser->next->data = dataHold;
        traverser = head;
        swapped = 0;
      }
      traverser = traverser->next;
    }
    
  }

  return head;
  
}

/* Using void pointers cuz I'm too lazy to add another parameter */
void Print_Job_Stat_Formatted(void* value, int max_width, int is_int) {
  if (is_int) {
    printf("%*d", max_width, *(int*)value);
  } else {
    printf("%-*s", max_width, (char*)value);
  }
}

void Print_Single_Job_Stats(Job* js) {
  Print_Job_Stat_Formatted((void*)js->name, 10, 0);
  // printf("\t");
  Print_Job_Stat_Formatted((void*)&js->runtime, strlen("CPU time")-1, 1);
  // printf("\t");
  Print_Job_Stat_Formatted((void*)&js->stats->completion_time,
                           strlen("when done")+2, 1);
  // printf("\t");
  Print_Job_Stat_Formatted((void*)&js->stats->cpu_count, strlen("cpu disp")+2,
                           1);
  // printf("\t");
  Print_Job_Stat_Formatted((void*)&js->stats->blocked_count,
                           strlen("i/o disp")+2, 1);
  // printf("\t");
  Print_Job_Stat_Formatted((void*)&js->stats->io_time, strlen("i/o time")+2, 1);
  printf("\n");
}

void Print_All_Job_Stats(Node* jobs) {
  printf("Processes:\n\n");
  char* padding = "";
  printf("%3sname%5sCPU time%2swhen done%2scpu disp%2si/o disp%2si/o time\n", padding, padding, padding, padding, padding, padding);

  Node* job = Sort_Jobs_Finish_First(jobs);
  
  Job* jobData;
  do {
    jobData = job->data;
    Print_Single_Job_Stats(jobData);
    job = job->next;
  } while (job);
}

void Print_Sys_Stats(Sys_Stats* Sys_stats) {
  printf("\nSystem:\n");
  printf("The wall clock time at which the simulation finished: %d\n",
         Sys_stats->finished_time);
}

void Print_CPU_Stats(CPU_Stats* CPU_stats, int jobcount) {
  printf("CPU:\n");
  printf("Total time spent busy: %d\n", CPU_stats->busy_time);
  printf("Total time spent idle: %d\n", CPU_stats->idle_time);
  printf("CPU utilization: %.2f\n",
         CPU_stats->busy_time /
             (CPU_stats->busy_time + (float)CPU_stats->idle_time));
  printf("Number of dispatches: %d\n", CPU_stats->cpu_dispatches);
  printf("Overall throughput: %.2f\n",
         jobcount /
             (CPU_stats->busy_time + (float)CPU_stats->idle_time));
}

void Print_IO_Stats(IO_Stats* IO_stats, int jobcount) {
  printf("I/O device:\n");
  printf("Total time spent busy: %d\n", IO_stats->busy_time);
  printf("Total time spent idle: %d\n", IO_stats->idle_time);
  printf(
      "I/O utilization: %.2f\n",
      IO_stats->busy_time / (IO_stats->busy_time + (float)IO_stats->idle_time));
  printf("Number of dispatches: %d\n", IO_stats->io_dispatches);
  printf("Overall throughput: %.2f\n",
         jobcount /
             (IO_stats->busy_time + (float)IO_stats->idle_time));
}

void Print_Stats(Node* jobs, Sys_Stats* Sys_stats, CPU_Stats* CPU_stats,
                 IO_Stats* IO_stats) {
  int jobcount = 0;
  Node* travel = jobs;
  while(travel != NULL) {
    jobcount++;
    travel = travel->next;
  }
  Print_All_Job_Stats(jobs);
  Print_Sys_Stats(Sys_stats);
  printf("\n");
  Print_CPU_Stats(CPU_stats, jobcount);
  printf("\n");
  Print_IO_Stats(IO_stats, jobcount);
}
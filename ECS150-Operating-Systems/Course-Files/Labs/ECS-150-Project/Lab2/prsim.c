#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "queue.c"
#include "helper.c"
/* #include "src/stats.c"*/

Job* Run_CPU(Queue*,Queue*,int, CPU_Stats*);

int main(int argc, char* argv[]) {
  if (argc < 3) {
    fprintf(stderr, "Usage: %s [-r | -f] file\n", argv[0]);
    return 1;
  }
  (void)srandom(12345);
  char* flag = argv[1];
  char* file = argv[2];

  Sys_Stats* Sys_stats = Sys_Stats_New();
  CPU_Stats* CPU_stats = CPU_Stats_New();
  IO_Stats* IO_stats = IO_Stats_New();

  if (strcmp(flag, "-r") != 0 && strcmp(flag, "-f") != 0) {
    fprintf(stderr, "Usage: %s [-r | -f] file\n", argv[0]);
    return 1;
  }

  /*printf("Running file %s with flag %s\n", file, flag);*/

  /* Load file */
  unsigned int num_jobs = 0;
  Node* jobs = LoadFile(file, &num_jobs);

  Queue* ready_queue = Queue_New();
  Queue* IO_queue = Queue_New();

  /* Load the ready queue */
  Node* node = jobs;
  while (node != NULL) {
    Job* job = node->data;
    Queue_Push(ready_queue, job);
    node = node->next;
  }

  Job* cpu = NULL;
  Job* io = NULL;
  Job* sameTickJob = NULL;
  unsigned int current_io_service_time_remaining = 0;
  int cpu_requires_no_blocking = 0;
  int blocked=0;
  int runtime=0;
  /*unsigned int blocking_runtime=0;*/

  unsigned int tick = 0;
  /* Run in FCFS mode */
  if (strcmp(flag, "-f") == 0) {
    /* Run the simulation */
    while (ready_queue->size > 0 || IO_queue->size > 0 || cpu || io) {
      /* Tick the clock */
      tick++;

      if (cpu == NULL) { /* If nothing in the cpu*/
        blocked=0;
        cpu = Queue_Pop(ready_queue);
        if (cpu) { // Make sure we popped somthing
          cpu->stats->cpu_count++;
          CPU_stats->cpu_dispatches++;
          if(cpu->runtime - cpu->time_elapsed==0) /* If we popped a 
            job with 0 time left, finish the job */
          {
            cpu->stats->completion_time = tick;
            cpu = NULL; // NULL it so idle time is incremented
          } else { // Calculate runtime
            runtime = cpu->runtime - cpu->time_elapsed;
            if(cpu->runtime - cpu->time_elapsed>1){
              float IO_transfer_chance = RandFloat(); /* Get a chance*/
              float cpu_block_chance =
                  cpu->blocking_probablity; /* Multiply by probability*/
              if (IO_transfer_chance < cpu_block_chance)
              {
                blocked = 1;
                runtime = RandInt(cpu->runtime - cpu->time_elapsed);
              }
            }
          }
        }
      }

      if(cpu != NULL) // Handle a job that finishes this cycle
      {
        if(cpu->runtime - cpu->time_elapsed==0)
        {
          cpu->stats->completion_time = tick ;
          //Print_Single_Job_Stats(cpu);
          cpu = NULL; // When job finishes we do nothing else        
        }
      }
      /* Check for blocking */
      if (cpu != NULL) { /* If the cpu is not null*/
        CPU_stats->busy_time++;
        cpu->time_elapsed++;
        runtime--;
        if(runtime==0) // Job is almost done in cpu
        {
          if(blocked) // Push to IO if we block
          {
            blocked=0;
            //printf("Pushed to IO q\n");
            Queue_Push(IO_queue, cpu);
            sameTickJob = cpu; // Remember the job we push this cycle
            cpu=NULL;
          } /* Do nothing if we do not block
            Job will finish next cycle */
        }
      } else {
        CPU_stats->idle_time++;
      }

      if (io == NULL) { /* If nothing stored in io*/
        if(sameTickJob==Queue_Peek(IO_queue) && sameTickJob != NULL) // 
        {
          //printf("Same tick job, ignoring io");
        } else // Only pop if we will not pop the job pushed this cycle
        io = Queue_Pop(IO_queue); /* Get something from queue*/
        if (io != NULL) {
          io->stats->blocked_count++;
          IO_stats->io_dispatches++;
          int time_left = io->runtime - io->time_elapsed;
          if(time_left == 0) // If time left 0, service for 1 time unit
            current_io_service_time_remaining = 1;
          else
            current_io_service_time_remaining = 
              RandInt(30); /* Set time allotment*/
        }
      }

      /* Check for IO */
      if (io != NULL) {
        IO_stats->busy_time++;
        io->stats->io_time++;
        current_io_service_time_remaining--;
        if (current_io_service_time_remaining <= 0) {
          Queue_Push(ready_queue, io);
          io = NULL;
        }
      } else {
        IO_stats->idle_time++;
      }

      sameTickJob=NULL;
      
    }
  
  
  
  
  
  
  } else {
    /* Run in RR mode */
    int max_quant = 0;
    while (ready_queue->size > 0 || IO_queue->size > 0 || cpu || io) {
      /* Tick the clock */
      tick++;

      if (cpu == NULL) { /* If nothing in the cpu*/
        blocked=0;
        cpu = Queue_Pop(ready_queue);
        if (cpu) { // Make sure we popped somthing
          cpu->stats->cpu_count++;
          CPU_stats->cpu_dispatches++;
          if(cpu->runtime - cpu->time_elapsed==0) /* If we popped a 
            job with 0 time left, finish the job */
          {
            cpu->stats->completion_time = tick;
            cpu = NULL; // NULL it so idle time is incremented
          } else { // Calculate runtime
            // runtime = cpu->runtime - cpu->time_elapsed; // Here is the problem
            if(cpu->runtime - cpu->time_elapsed>1){
              float IO_transfer_chance = RandFloat(); /* Get a chance*/
              float cpu_block_chance =
                  cpu->blocking_probablity; /* Multiply by probability*/
              if (IO_transfer_chance < cpu_block_chance)
              {
                blocked = 1;
                runtime = RandInt(5);
                if(cpu->runtime - cpu->time_elapsed < runtime) {
                  runtime = cpu->runtime - cpu->time_elapsed;
                }
              }
            }
          }
        }
      }

      if(cpu != NULL) // Handle a job that finishes this cycle
      {
        if(cpu->runtime - cpu->time_elapsed==0)
        {
          cpu->stats->completion_time = tick ;
          //Print_Single_Job_Stats(cpu);
          cpu = NULL; // When job finishes we do nothing else        
        }
      }
      /* Check for blocking */
      if (cpu != NULL) { /* If the cpu is not null*/
        CPU_stats->busy_time++;
        cpu->time_elapsed++;
        runtime--;
        if(runtime==0) // Job is almost done in cpu
        {
          if(blocked) // Push to IO if we block
          {
            blocked=0;
            //printf("Pushed to IO q\n");
            Queue_Push(IO_queue, cpu);
            sameTickJob = cpu; // Remember the job we push this cycle
            cpu=NULL;
          } /* Do nothing if we do not block
            Job will finish next cycle */
          else{ // This is if blocked is not created

          }
        }
      } else {
        CPU_stats->idle_time++;
      }

      if (io == NULL) { /* If nothing stored in io*/
        if(sameTickJob==Queue_Peek(IO_queue) && sameTickJob != NULL) // 
        {
          //printf("Same tick job, ignoring io");
        } else // Only pop if we will not pop the job pushed this cycle
        io = Queue_Pop(IO_queue); /* Get something from queue*/
        if (io != NULL) {
          io->stats->blocked_count++;
          IO_stats->io_dispatches++;
          int time_left = io->runtime - io->time_elapsed;
          if(time_left == 0) // If time left 0, service for 1 time unit
            current_io_service_time_remaining = 1;
          else
            current_io_service_time_remaining = 
              RandInt(30); /* Set time allotment*/
        }
      }

      /* Check for IO */
      if (io != NULL) {
        IO_stats->busy_time++;
        io->stats->io_time++;
        current_io_service_time_remaining--;
        if (current_io_service_time_remaining <= 0) {
          Queue_Push(ready_queue, io);
          io = NULL;
        }
      } else {
        IO_stats->idle_time++;
      }

      sameTickJob=NULL;
      
    }
  }
  Sys_stats->finished_time = tick;
  /*Print_All_Job_Stats(jobs);*/
  Print_Stats(jobs, Sys_stats, CPU_stats, IO_stats);
  /* Free Jobs */
  // printf("Freeing Jobs\n");
  node = jobs;
  while (node->next != NULL) {
    Node* self = node;
    node = node->next;
    free(self->data);
    free(self);
  }

  return 0;
}

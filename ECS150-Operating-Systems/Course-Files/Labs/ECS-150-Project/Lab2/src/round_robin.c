#include <stdio.h>
#include <stddef.h>
#include <stdlib.h>

#include "job.c"
#include "queue.c"
#include "helper.c"

typedef struct {
    Queue* cpu_queue;
    Queue* io_queue;
    int quant_left;
    bool isLocked;
} RR;

RR* RR_New() {
    RR *robin = malloc(sizeof(RR));
    robin->cpu_queue = Queue_New();
    robin->io_queue = Queue_New(); // This might be converted to FCFS
    robin->quant_left = -1;
    robin->isLocked = false;
}

bool Will_Block_CPU(RR* robin) {
    // Here we determine if it needs to block, so we get its probability and 
    robin->cpu_queue->front;

}

bool Will_Block_IO(RR* robin) {

}

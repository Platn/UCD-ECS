#include <stdio.h>
#include <stddef.h>
#include <stdlib.h>

#include "job.c"

typedef struct {
  unsigned int size;
  struct Node* front;
  struct Node* back;
} Queue;

/* Contructor */
Queue* Queue_New() {
  Queue* queue = malloc(sizeof(Queue));
  queue->size = 0;
  queue->front = NULL;
  queue->back = NULL;
  return queue;
}

/* Destructor */
void Queue_Destroy(Queue* queue) {
  Node* node = queue->front;
  while (node != NULL) {
    Node* next = node->next;
    free(node);
    node = next;
  }
  free(queue);
}

/* Push a job to the back of the queue */
void Queue_Push(Queue* queue, Job* job) {
  Node* node = malloc(sizeof(Node));
  node->data = job;
  node->next = NULL;
  if (queue->size == 0) {
    queue->front = node;
    queue->back = node;
  } else {
    queue->back->next = node;
    queue->back = node;
  }
  queue->size++;
}

/* Return the job at the front of the queue without changing the queue */
Job* Queue_Peek(Queue* queue) {
  if (queue->size == 0) {
    return NULL;
  } else {
    return queue->front->data;
  }
}

/* Pop a job from the front of the queue */
Job* Queue_Pop(Queue* queue) {
  if (queue->size == 0) {
    return NULL;
  }
  Node* node = queue->front;
  queue->front = node->next;
  queue->size--;
  return node->data;
}

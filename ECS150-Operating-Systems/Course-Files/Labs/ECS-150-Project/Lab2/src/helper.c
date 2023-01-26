#include <errno.h>
#include <stdio.h>
#include <stdlib.h>

static int RandInt(int max) {
  int r = random();
  int n = (r % max) + 1;
  // printf("%10d i %7d Input: %3d ", r, n, max);
  return n;
}

static float RandFloat() {
  int r = random();
  float n = (r/(float)RAND_MAX);
  // printf("%10d f %6f ", r, n);
  return n;
}

static Node* LoadFile(char* file_path, unsigned int* num_jobs) {
  FILE* file = fopen(file_path, "r");
  if (file == NULL) {
    fprintf(stderr, "%s: No such file or directory\n", file_path);
    /*errno = ENOENT;
	  perror(file_path);*/
    exit(1);
  }

  /* Read number of jobs */
  Node* node = malloc(sizeof(Node));
  Node* head = node;
  char* buffer = NULL;
  size_t buffer_size;
  int lines_read = 0;

  while (getline(&buffer, &buffer_size, file) != EOF) {
    char name_buffer[100];
    int runtime;
    float blocking_probablity;

    int input_read = sscanf(buffer, "%s %d %f", name_buffer, &runtime, &blocking_probablity);

    /* Error Checking */
    if (input_read != 3) {  
      fprintf(stderr, "Malformed line %s(%d)\n", file_path, lines_read + 1);
      exit(1);
    }

    if (strlen(name_buffer) > 10) {
      fprintf(stderr, "name is too long %s(%d)\n", file_path, input_read);
      exit(1);
    }

    if (runtime < 0) {
      fprintf(stderr, "runtime is not positive interger %s(%d)\n", file_path, lines_read);
      exit(1);
    }

    if (blocking_probablity < 0 || blocking_probablity > 1) {
      fprintf(stderr, "probability < 0 or > 1 %s(%d)\n", file_path, lines_read);
      exit(1);
    }

    /* Create node */
    if (node->data == NULL) {
      node->data = Job_New(name_buffer, runtime, blocking_probablity, LOW);
    } else {
      node->next = malloc(sizeof(Node));
      node = node->next;
      node->next = NULL;
      node->data = Job_New(name_buffer, runtime, blocking_probablity, LOW);
    }
    (*num_jobs)++;
    lines_read++;
  }

  fclose(file);
  return head;
}

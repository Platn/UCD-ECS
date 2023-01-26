#include <stdio.h>

int main(int argc, char *argv[]) {
  if (argc < 2) {
    fprintf(stderr, "Not enough arguments\n");
    return 1;
  }
  while (1) {
    printf("Hello, world!\n");
  }
  return 1;
}
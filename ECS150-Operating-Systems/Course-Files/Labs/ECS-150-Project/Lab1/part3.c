#include <stdio.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>

#ifdef DEBUG
#define DEBUG_PRINT(...) fprintf( stderr, __VA_ARGS__ );
#else
#define DEBUG_PRINT(...)
#endif

void debug_dump(char *filename, struct stat *file_stat) {
      DEBUG_PRINT("---------------------------\n");
    DEBUG_PRINT("Information for %s\n", filename);
    DEBUG_PRINT("---------------------------\n");
    DEBUG_PRINT("File Size: \t\t%d bytes\n", file_stat->st_size);
    DEBUG_PRINT("Number of Links: \t%d\n", file_stat->st_nlink);
    DEBUG_PRINT("File inode: \t\t%d\n", file_stat->st_ino);

    DEBUG_PRINT("File Permissions: \t");
    DEBUG_PRINT( (S_ISLNK(file_stat->st_mode)) ? "s " : "- ");
    DEBUG_PRINT( (S_ISDIR(file_stat->st_mode)) ? "d" : "-");
    DEBUG_PRINT( (file_stat->st_mode & S_IRUSR) ? "r" : "-");
    DEBUG_PRINT( (file_stat->st_mode & S_IWUSR) ? "w" : "-");
    DEBUG_PRINT( (file_stat->st_mode & S_IXUSR) ? "x" : "-");
    DEBUG_PRINT( (file_stat->st_mode & S_IRGRP) ? "r" : "-");
    DEBUG_PRINT( (file_stat->st_mode & S_IWGRP) ? "w" : "-");
    DEBUG_PRINT( (file_stat->st_mode & S_IXGRP) ? "x" : "-");
    DEBUG_PRINT( (file_stat->st_mode & S_IROTH) ? "r" : "-");
    DEBUG_PRINT( (file_stat->st_mode & S_IWOTH) ? "w" : "-");
    DEBUG_PRINT( (file_stat->st_mode & S_IXOTH) ? "x" : "-");
    DEBUG_PRINT("\n\n");
}

int main(int argc, char *argv[]) {
  if (argc < 3) {
    printf("Usage: %s <file1> <file2>\n", argv[0]);
    return 1;
  }

  struct stat file1_stat, file2_stat;
  int file_1_ret;
  file_1_ret = lstat(argv[1], &file1_stat);
  debug_dump(argv[1], &file1_stat);
  if (file_1_ret <0 ) {
    perror(argv[1]);
    return 1;
  }

  int file_2_ret;
  file_2_ret = lstat(argv[2], &file2_stat);
  debug_dump(argv[2], &file2_stat);
  if (file_2_ret <0 ) {
    perror(argv[2]);
    return 1;
  }

  ssize_t readlink_ret_1;
  char symlink_1_contents [4096];
  char filename_1_buffer [4096];
  symlink_1_contents[0] = '\0';
  strcpy(filename_1_buffer, argv[1]);
  int file1_stat_ret;
  if S_ISLNK(file1_stat.st_mode) {
    readlink_ret_1 = readlink(filename_1_buffer, symlink_1_contents, 4096);
    /*its important to set bufsize
    (3rd arg above) explicitly,
    because otherwise readlink() relies on
    st_size of the stat which may be 0
    for empty or dynamically generated files. */
    if ((int)readlink_ret_1 < 0) {
      perror(argv[1]);
      return 1;
    }
    symlink_1_contents[(int)readlink_ret_1] = '\0';
    file1_stat_ret = lstat(symlink_1_contents,&file1_stat);
    if (file1_stat_ret < 0) {
      perror(symlink_1_contents);
      return 1;
    }
    DEBUG_PRINT("Following %s ... \n", filename_1_buffer);
    debug_dump(symlink_1_contents, &file1_stat);
    strcpy(filename_1_buffer,symlink_1_contents);
  }


  if ((symlink_1_contents[0] != '\0') && (strcmp(symlink_1_contents, argv[2]) == 0) && !S_ISLNK(file2_stat.st_mode)) {
    printf("%s is a symbolic link to %s\n",argv[1],argv[2]);
    return 0;
  }
  file_1_ret = lstat(argv[1], &file1_stat);
  debug_dump(argv[1], &file1_stat);
  if (file_1_ret <0 ) {
    perror(argv[1]);
    return 1;
  }
  ssize_t readlink_ret_2;
  char symlink_2_contents [4096];
  char filename_2_buffer [4096];
  symlink_2_contents[0] = '\0';
  int file2_stat_ret;
  strcpy(filename_2_buffer, argv[2]);
  if S_ISLNK(file2_stat.st_mode) {
    readlink_ret_2 = readlink(filename_2_buffer, symlink_2_contents, 4096);
    if ((int)readlink_ret_2 < 0) {
      DEBUG_PRINT("%s\n", filename_2_buffer);
      perror(argv[2]);
      return 1;
    }
    symlink_2_contents[(int)readlink_ret_2] = '\0';
    file2_stat_ret = lstat(symlink_2_contents,&file2_stat);
    if (file2_stat_ret < 0) {
      perror(symlink_2_contents);
      return 1;
    }
    DEBUG_PRINT("Following %s ... \n", filename_2_buffer);
    debug_dump(symlink_2_contents, &file2_stat);
    strcpy(filename_2_buffer,symlink_2_contents);
  }
  if ((symlink_2_contents[0] != '\0') && (strcmp(symlink_2_contents, argv[1]) == 0) && !S_ISLNK(file1_stat.st_mode)) {
    printf("%s is a symbolic link to %s\n",argv[2],argv[1]);
    return 0;
  }

  char* real_file1,real_file2;
  char buf1[4096],buf2[4096];
  real_file1=realpath(argv[1],buf1);
  real_file2=realpath(argv[2],buf2);
  if(real_file1 == NULL)
  {
    perror(argv[1]);
    return 1;
  }
  if(real_file2 == NULL)
  {
    perror(argv[2]);
    return 1;
  }

  file_1_ret = lstat(buf1, &file1_stat);
  debug_dump(argv[1], &file1_stat);
  if (file_1_ret <0 ) {
    perror(argv[1]);
    return 1;
  }

  file_2_ret = lstat(buf2, &file2_stat);
  debug_dump(argv[2], &file2_stat);
  if (file_2_ret <0 ) {
    perror(argv[2]);
    return 1;
  }

  if ((file1_stat.st_dev == file2_stat.st_dev) && (file1_stat.st_ino == file2_stat.st_ino)) {
    printf("These files are linked.\n");
    return 0;
  }
  printf("The files are not linked.\n");
  return 0;
}

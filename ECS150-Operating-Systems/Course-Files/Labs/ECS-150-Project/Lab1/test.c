#include  <stdlib.h>
#include  <stdio.h>
 
/* Where the environment variable 'PATH' is set to a value. */
 
int main(void)
{
   char *pathvar;
 
   pathvar = getenv("PATH");
   printf("pathvar=%s",pathvar);
}

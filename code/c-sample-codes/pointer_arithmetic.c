# include <stdio.h>
# include <stdlib.h>


int main(int argc, char **argv) {

  int sample_var = 100;
  int *ptr;
  ptr = &sample_var;
  printf("Address of: %x\n", ptr);
  ptr++;
  printf("Address of: %x\n", ptr);


  return 0;
}

#include <errno.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, const char* argv[const]) {
  (void)(argc);
  const char* fileName = argv[0];
  FILE* file = fopen(fileName, "r");
  if (file == NULL) {
    perror("Error: Could not open file");
    return EXIT_FAILURE;
  }

  int existStatus = EXIT_SUCCESS;

  while (true) {
    const char c = (char)fgetc(file);
    if (feof(file)) {
      break;
    }
    if (ferror(file)) {
      perror("Error: Could not read file");
      existStatus = EXIT_FAILURE;
      break;
    }
    printf("%c", c);
  }
  // Close the file
  fclose(file);
  return existStatus;
}

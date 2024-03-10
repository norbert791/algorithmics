#include <errno.h>
#include <stdio.h>
#include <string.h>

static size_t matchPattern(const char pattern[static 1],
                           const char text[static 1]);

int main(void) {
  const char* pattern = "Wor";
  const char* text = "Hello World";
  errno = 0;
  size_t result = matchPattern(pattern, text);
  if (errno != 0) {
    printf("Not found\n");
  } else {
    printf("Index: %zu\n", result);
  }

  return 0;
}

static size_t matchPattern(const char pattern[const static 1],
                           const char text[const static 1]) {
  register const size_t n = strlen(pattern);
  register const size_t m = strlen(text);

  for (size_t i = 0; i + n < m; i++) {
    if (memcmp(pattern, &text[i], n) == 0) {
      errno = 0;
      return i;
    }
  }

  errno = 1;
  return m;
}
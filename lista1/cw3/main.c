#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Find the largest index i such that x[0..m-1] = y[i..i+m-1]. result must point
// to a valid memory location.
static bool find(const char x[static 1], const char y[static 1],
                 size_t* result);

int main(void) {
  const char* x = "Hello";
  const char* y = "WorldHello";
  size_t result = 0;
  if (find(x, y, &result)) {
    printf("Index: %zu\n", result);
  } else {
    printf("Not found\n");
  }
  return 0;
}

static bool find(const char x[const static 1], const char y[const static 1],
                 size_t* const result) {
  register const size_t n = strlen(x);
  register const size_t m = strlen(y);

  bool foundOnce = false;
  *result = 0;

  for (size_t i = 0; i < n && i < m; i++) {
    size_t l = 0;
    size_t k = m - i - 1;

    bool found = true;
    while (l < i && k < m) {
      if (x[l] != y[k]) {
        found = false;
        break;
      }
      l++;
      k++;
    }
    if (found) {
      foundOnce = true;
      *result = i;
    }
  }

  return foundOnce;
}
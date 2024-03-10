#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static bool isPrefix(const char str[static 1], const char target[static 1]);
static bool isSufix(const char str[static 1], const char target[static 1]);

int main() {
  const char* str = "Hello";
  const char* target = "Hello World";
  printf("Is prefix: %d\n", isPrefix(str, target));
  assert(isPrefix(str, target));

  const char* str2 = "World";
  const char* target2 = "Hello World";
  printf("Is sufix: %d\n", isSufix(str2, target2));
  assert(isSufix(str2, target2));

  const char* str3 = "Helloo";
  const char* target3 = "Hello";
  printf("Is prefix: %d\n", isPrefix(str3, target3));
  assert(!isPrefix(str3, target3));

  const char* str4 = "Worldd";
  const char* target4 = "World";
  printf("Is sufix: %d\n", isSufix(str4, target4));
  assert(!isSufix(str4, target4));

  const char* str5 = "";
  const char* target5 = "Hello";
  printf("Is prefix: %d\n", isPrefix(str5, target5));
  printf("Is sufix: %d\n", isSufix(str5, target5));
  assert(isPrefix(str5, target5));
  assert(isSufix(str5, target5));

  const char* str6 = "Hello";
  const char* target6 = "";
  printf("Is prefix: %d\n", isPrefix(str6, target6));
  printf("Is sufix: %d\n", isSufix(str6, target6));
  assert(!isPrefix(str6, target6));
  assert(!isSufix(str6, target6));

  const char* str7 = "";
  const char* target7 = "";
  printf("Is prefix: %d\n", isPrefix(str7, target7));
  printf("Is sufix: %d\n", isSufix(str7, target7));
  assert(isPrefix(str7, target7));
  assert(isSufix(str7, target7));

  return 0;
}

static bool isPrefix(const char str[const static 1],
                     const char target[const static 1]) {
  // Contract check
  if (str == NULL || target == NULL) {
    perror("isPrefix: str or target is NULL");
    exit(EXIT_FAILURE);
  }
  size_t strIndex = 0;
  size_t targetIndex = 0;

  while (str[strIndex] != '\0' && target[targetIndex] != '\0') {
    if (str[strIndex] != target[targetIndex]) {
      return false;
    }
    strIndex++;
    targetIndex++;
  }
  if (str[strIndex] == '\0') {
    return true;
  }

  return false;
}

static bool isSufix(const char str[static 1], const char target[static 1]) {
  // Contract check
  if (str == NULL || target == NULL) {
    perror("isSufix: str or target is NULL");
    exit(EXIT_FAILURE);
  }
  size_t strLen = strlen(str);
  size_t targetLen = strlen(target);

  while (strLen > 0 && targetLen > 0) {
    if (str[strLen - 1] != target[targetLen - 1]) {
      return false;
    }
    strLen--;
    targetLen--;
  }
  if (strLen == 0 && str[strLen] == target[targetLen]) {
    return true;
  }

  return false;
}

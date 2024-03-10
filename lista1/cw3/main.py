#!/usr/bin/env python3


def find(x: str, y: str) -> int:
    maxI = -1
    for i in range(min(len(x), len(y))):
        yIndex = len(y) - i - 1
        print(x[: i + 1], y[yIndex:])
        if x[: i + 1] == y[yIndex:]:
            maxI = i
    return maxI


if __name__ == "__main__":
    str1 = "World"
    str2 = "Hello World"
    n = find(str1, str2)
    if n == -1:
        print("not found")
    else:
        print(f"found at index {n}")

#!/usr/bin/env python3

from sys import argv


def rabin_karp_matcher(text: str, pattern: str) -> bool:
    n = len(text)
    m = len(pattern)
    d = 256
    q = 71993
    h = pow(d, m - 1) % q
    p = 0
    t = 0
    for i in range(m):
        p = (d * p + ord(pattern[i])) % q
        t = (d * t + ord(text[i])) % q
    for s in range(n - m + 1):
        if p == t:
            if pattern == text[s : s + m]:
                return True
        if s < n - m:
            t = (d * (t - ord(text[s]) * h) + ord(text[s + m])) % q
    return False


if __name__ == "__main__":
    str1 = argv[1]
    str2 = argv[2]
    if rabin_karp_matcher(str2, str1):
        print("found")
    else:
        print("not found")

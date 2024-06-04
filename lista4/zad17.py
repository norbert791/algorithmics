#!/usr/bin/env python3

from random import random
import matplotlib.pyplot as plt

if __name__ == "__main__":
    maxNumOfStations = int(1e3)
    results = [0] * maxNumOfStations
    for i in range(1, int(maxNumOfStations) + 1):
        if i <= 400:
            results[i - 1] = i
        else:
            randStations = list(map(lambda _: random(), range(i)))
            randStations.sort()
            results[i - 1] = 399 / randStations[399]

    for i in range(maxNumOfStations):
        exactResult = i + 1
        results[i] = results[i] / exactResult
    plt.plot(range(1, len(results) + 1), results)
    plt.xlabel("Num of stations")
    plt.ylabel("Result relative err")
    plt.title("Relative error of the result")
    plt.show()

#!/usr/bin/env python3

from random import randrange
import matplotlib.pyplot as plt

if __name__ == "__main__":
    numOfBalls = int(1e6)
    numOfBins = int(1e3)
    bins = [0] * numOfBins
    for i in range(numOfBalls):
        bins[randrange(numOfBins)] += 1

    plt.bar(range(numOfBins), bins)
    plt.xlabel("Bin")
    plt.ylabel("Number of Balls")
    plt.title("Distribution of Balls in Bins")
    plt.savefig("uniform.png")
    plt.close()

    bins = [0] * numOfBins
    for i in range(numOfBalls):
        a = randrange(numOfBins)
        b = randrange(numOfBins)
        if bins[a] < bins[b]:
            bins[a] += 1
        else:
            bins[b] += 1

    plt.bar(range(numOfBins), bins)
    plt.xlabel("Bin")
    plt.ylabel("Number of Balls")
    plt.title("Distribution of Balls in Bins")
    plt.savefig("two choices.png")
    plt.close()

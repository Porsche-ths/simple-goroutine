import random

file = open("./files/data.csv", "w")

for i in range(10000000):
    file.write("{:.2f}".format(random.random() * 100) + "\n")

file.close()
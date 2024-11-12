file = open("./files/file.csv", "w")

for i in range(1000):
    file.write(str(i) + "," + str(i) + "\n")

file.close()
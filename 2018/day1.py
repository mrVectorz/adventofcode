# https://adventofcode.com/2018/day/1

# part 1
freq=0
with open("day_1_input.txt", "r") as file:
  for line in file.readlines():
    freq += int(line.split()[0])

print(freq)


# part 2

freq = 0
freq_list = [0]
counter = 0
while True:
  print("Exec number: {}".format(counter))
  with open("day_1_input.txt", "r") as file:
    for line in file.readlines():
      freq += int(line.split()[0])
      if freq in freq_list:
        print("First dupe: {}".format(freq))
        exit()
      freq_list.append(freq)
  counter += 1

# slow
import numpy as np
freq = 0
freq_list = [0]
counter = 0
while True:
  print("Exec number: {}".format(counter))
  with open("day_1_input.txt", "r") as file:
    for line in file.readlines():
      freq += int(line.split()[0])
      freq_list.append(freq)
      if (np.unique(freq_list, return_counts=True)[1] >= 2).all():
        print("First dupe: {}".format(freq))
        exit()
  counter += 1

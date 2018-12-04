# https://adventofcode.com/2018/day/2
# Author: Marc Methot
# Email: mb.methot@gmail.com

# part 1
multi = {"doubles":0, "triples":0}

def checker(ID, double=False, triple=False):
  for l in ID:
    if not double and ID.count(l) == 2:
        multi['doubles'] = multi['doubles']+1
        double = True
    elif not triple and ID.count(l) == 3:
        multi['triples'] = multi['triples']+1
        triple = True
    elif double and triple:
      break

with open("day_2_input.txt", "r") as f:
  for line in f.readlines():
    checker(line.strip())

print(multi)
print("Total: {}\n".format(multi['doubles']*multi['triples']))

# part 2
### FUGLY ###
from difflib import SequenceMatcher

with open("day_2_input.txt", "r") as f:
  IDs = f.readlines()

string_dist_list = []
for i in IDs:
  for m in IDs:
    if m != i:
      string_dist_list.append((i, m, SequenceMatcher(None, i, m).ratio()))

sorted_list = sorted(string_dist_list, key=lambda tup: tup[2])

for count, letter in enumerate(sorted_list[len(sorted_list)-1][0]):
  if letter != sorted_list[len(sorted_list)-1][1][count]:
    uniq=count

print("Common letters: {}{}".format(sorted_list[len(sorted_list)-1][0][:uniq],
                sorted_list[len(sorted_list)-1][0][uniq+1:].strip()))


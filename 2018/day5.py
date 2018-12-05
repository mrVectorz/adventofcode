# https://adventofcode.com/2018/day/2
# Author: Marc Methot
# Email: mb.methot@gmail.com

# part 1
import string
f = open("day_5_input.txt", "r").read().strip()

letters = {}
for i in string.ascii_lowercase:
  letters[i.lower()] = i.upper()
  letters[i.upper()] = i.lower()

poly = []

for char in f:
  if poly and char == letters[poly[-1]]:
    poly.pop()
  else:
    poly.append(char)

print(len(poly))

# part 2
lenght = 1e6
for letter in string.ascii_lowercase:
  temp = [char for char in f if char != letter.lower() and char != letter.upper()]
  poly = []
  for i in temp:
    if poly and i == letters[poly[-1]]:
      poly.pop()
    else:
      poly.append(i)
  lenght = min(lenght, len(poly))
print(lenght)

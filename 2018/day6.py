# https://adventofcode.com/2018/day/6
# Author: Marc Methot
# Email: mb.methot@gmail.com

from collections import Counter

# Initial thought was to grow each point one by one
# Not a success story
def grow(coord,count):
  x = coord[1]
  y = coord[0]
  e = True
  steps = 0
  while e:
    if gmap[x+steps][y] == c[2]:
      print("Step inc")
      steps += 1
    else:
      e = False
  for Y in range(x-steps,1+x+steps):
    for X in range(y-steps,1+y+steps):
      print((X,Y))
      gmap[Y][X] = c[2]

# manhattan distance
def m_dist(v1, data):
  return min(abs(v1[0]-v2[0]) + abs(v1[1]-v2[1]) for v2 in data)

with open('day_6_input.txt') as f:
  coords = [i.strip().split(", ") for i in f.readlines()]
coords = [(int(i[0]), int(i[1])) for i in coords]

max_x = max(list(zip(*coords))[0])
max_y = max(list(zip(*coords))[1])

gmap = {}

for x in range(max_x):
  for y in range(max_y):
    md = m_dist((x,y), coords)
    for i, (X, Y) in enumerate(coords):
      if abs(x-X)+abs(y-Y) == md:
        if gmap.get((x, y), -1) != -1:
          gmap[x,y] = -1
          break
        gmap[x,y] = i

s = set([-1])
s = s | set(gmap[x, max_y-1] for x in range(max_x))
s = s | set(gmap[x, 0] for x in range(max_x))
s = s | set(gmap[max_x-1, y] for y in range(max_y))
s = s | set(gmap[0, y] for y in range(max_y))

print(f"Part 1: {[i[1] for i in Counter(gmap.values()).most_common() if i[0] not in s][0]}")

p2 = []
for x in range(max_x):
  for y in range(max_y):
    #get total dist
    a = sum([abs(x-k)+abs(y-l) for k, l in coords])
    if a < 1e4:
      p2.append(a)

print(f"Part 2: {len(p2)}")

#for i in range(3):
#  grow(coords[49],1)

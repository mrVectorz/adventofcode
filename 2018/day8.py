# https://adventofcode.com/2018/day/8
# Author: Marc Methot
# Email: mb.methot@gmail.com

f = open("day_8_input.txt", "r").read() 
test = [int(i) for i in f.split()]

T = []
switch = False
p = 0
n = 0
for node in test:
  if p-n >= len(test):
    break
  m_count = test[p+1]
  if not switch:
    switch=True
    n -= m_count
    metadata = test[n:][:m_count]
    T.append({"childs": test[p],
              "metadata count": m_count,
              "metadata": metadata})
    p += 2
  elif switch:
    switch=False
    metadata = test[p+2:p+2+m_count]
    T.append({"childs": test[p], "metadata count": m_count,
              "metadata": metadata})
    p += abs(2+m_count)

print(f"lenght of array: {len(test)}\nCounter: {p-n}\tp:{p}\tn:{n}")
print(sum([sum(i['metadata']) for i in T]))

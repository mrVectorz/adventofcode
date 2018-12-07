with open('day_7_input.txt') as f:
  steps = [i.strip() for i in f.readlines()]

steps = [i.split()[1:8:6] for i in steps]

workflow = []
while len(steps) >= 1:
  steps.sort()
  for step in steps:
    if step[0] in map(lambda i: (i[1]), steps):
      continue
    else:
      workflow.append(step[0])
      if len(steps) == 1:
        workflow.append(step[1])
      while step[0] in map(lambda i: (i[0]), steps):
        steps.pop(list(map(lambda i: (i[0]), steps)).index(step[0]))
      break

for i in workflow:
  print(i, end="")
print()

# day 3
# part 1

material = (1000, 1000)
fabric = {i:[0]*material[1] for i in range(material[0])}

claims = []
with open("day_3_input.txt", "r") as f:
  for claim in f.readlines():
    claims.append(claim.split())

# ['#123', '@', '3,2:', '5x4']
for claim in claims:
  l = int(claim[2].strip(":").split(",")[0])
  h = int(claim[2].strip(":").split(",")[1])
  L = int(claim[3].split("x")[0])
  H = int(claim[3].split("x")[1])
  for y in range(h,h+H):
    for x in range(l,l+L):
      fabric[y][x] = fabric[y][x]+1

overlap = 0
for i in fabric:
  overlap += sum(n > 1 for n in fabric[i])
print(overlap)


# part 2

def check_mah_claim(claim, overlap=False):
  l = int(claim[2].strip(":").split(",")[0])
  h = int(claim[2].strip(":").split(",")[1])
  L = int(claim[3].split("x")[0])
  H = int(claim[3].split("x")[1])
  for y in range(h,h+H):
    for x in range(l,l+L):
      if fabric[y][x] != 1:
        return True
  return False

for claim in claims:
  if not check_mah_claim(claim):
    print(claim[0].strip("#"))

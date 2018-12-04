# https://adventofcode.com/2018/day/4
# Author: Marc Methot
# Email: mb.methot@gmail.com

# part 1
# Collected way too much data for no reason
from datetime import datetime

with open("day_4_input.txt", "r") as f:
  logs = f.readlines()

logs.sort(key=lambda x: datetime.strptime(x.split("]")[0][1:], '%Y-%m-%d %H:%M'))

guard_id = 0
sleep_stats = []
for log in logs:
  time = datetime.strptime(log.split("]")[0][1:], '%Y-%m-%d %H:%M')
  if "Guard" in log:
    guard_id = int(log.split()[3].strip("#"))
    if not any(d.get("guard", None) == guard_id for d in sleep_stats):
      sleep_stats.append({"guard":guard_id, "sleep_total":0, "lsleep":0, "minutes":[]})
      index = len(sleep_stats)-1
    else:
      index = [c for c,i in enumerate(sleep_stats) if sleep_stats[c]['guard'] == guard_id][0]
  elif "falls" in log:
    asleep = time.minute
  elif "wakes" in log:
    tsleep = sleep_stats[index]["sleep_total"]
    nsleep = time.minute-asleep
    lsleep = sleep_stats[index]["lsleep"]
    lsleep = lsleep if lsleep >= nsleep else nsleep

    nminutes = []
    minutes = sleep_stats[index]["minutes"]
    for m in range(asleep, time.minute):
      minutes.append(m)
    sleep_stats[index] = {"guard":guard_id, "sleep_total": tsleep+nsleep, "lsleep":lsleep, "minutes":minutes}


sleep_stats.sort(key=lambda x: x['sleep_total'])

ID,tsleep,sleep,minutes = sleep_stats[len(sleep_stats)-1].values()
cminutes = [(i, minutes.count(i)) for i in minutes]
cminutes.sort(key=lambda x: x[1], reverse=True)
print("Part1 -- ID: {} - minute: {} - Result: {}".format(ID, cminutes[0][0], ID*cminutes[0][0]))


# part 2
s = []
for i in sleep_stats:
  ID = i["guard"]
  minutes = [(x, i['minutes'].count(x)) for x in i['minutes']]
  minutes.sort(key=lambda x: x[1], reverse=True)
  if len(minutes) > 1:
    minute = minutes[0]
    s.append((ID, minute))

s.sort(key=lambda x: x[1][1], reverse=True)

print("Part2 -- ID: {} - minute: {} - Result: {}".format(s[0][0], s[0][1][0], s[0][0]*s[0][1][0]))

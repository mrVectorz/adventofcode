inputs="day_10_input.txt"; a=0; for i in `sort -n $inputs` $(echo $(sort -n ${inputs} | tail -1)+3|bc); do echo $a-$i | bc; a=$i; done | sort -n | uniq -c | awk 'BEGIN {a=1}; {a=a*$1}; END {print a}'

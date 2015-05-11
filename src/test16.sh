# bash src/test16.sh data/hightemp.txt [outdir] [n]
# lcount = hightempの行数 = 24
lcount=`cat $1 | wc -l`
# lcount + n - 1 = 24 + 4 - 1 = 27
lcount=`expr $lcount + $3 - 1`
# num = lcount / n = = 27 / 4 = 6
num=`expr $lcount / $3` 
split -l $num $1 $2/split_
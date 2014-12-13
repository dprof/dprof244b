#!/bin/sh

socatopt="-t 1 -T 1 -b 16384"

val=`echo abc`
val=`printf "%s" "${val}"`
vallen=`printf "%s" "${val}" | wc -c`
set_command=""
set_commands=""

# build
for i in `seq 1 1`; do
    if [ `expr $i % 2` -eq "0" ]; then
        key="foo"
    else
        key="bar"
    fi
    key=`printf "%s%d" "${key}" "${i}"`

    set_command="set ${key} 0 0 ${vallen}\r\n${val}\r\n"
    set_commands=`printf "%s%s" "${set_commands}" "${set_command}"`
done

printf "%b" "$set_commands" > /tmp/socat.input

# write
#    cat /tmp/socat.input | socat ${socatopt} - TCP:corn21.stanford.edu:9102,nodelay,shut-down,nonblock=1 &
#    printf 'set a b' | socat ${socatopt} - TCP:corn21.stanford.edu:9102,nodelay,shut-down,nonblock=1 &

#printf '\nsetbit\n'
#printf '*2\r\n$3\r\ndel\r\n$3\r\nfoo\r\n' | socat ${debug} ${timeout} - TCP:corn21.stanford.edu:9102,shut-close
while true; do
printf '*3\r\n$3\r\nset\r\n$3\r\nfoo\r\n$3\r\noof\r\n' | socat ${debug} ${timeout} - TCP:10.36.44.50:9002,shut-close
sleep 1
printf '*3\r\n$3\r\nset\r\n$3\r\nfoo\r\n$3\r\noof\r\n' | socat ${debug} ${timeout} - TCP:10.36.44.50:9002,shut-close
sleep 1
printf '*3\r\n$3\r\nset\r\n$3\r\nfoo\r\n$3\r\noof\r\n' | socat ${debug} ${timeout} - TCP:10.36.44.50:9002,shut-close
sleep 1
printf '*3\r\n$3\r\nset\r\n$3\r\nfoo\r\n$3\r\noof\r\n' | socat ${debug} ${timeout} - TCP:10.36.44.50:9002,shut-close
sleep 1
printf '*3\r\n$3\r\nset\r\n$3\r\nfoo\r\n$3\r\noof\r\n' | socat ${debug} ${timeout} - TCP:10.36.44.50:9002,shut-close
sleep 1
printf '*3\r\n$3\r\nset\r\n$3\r\nfoo\r\n$3\r\noof\r\n' | socat ${debug} ${timeout} - TCP:10.36.44.50:9002,shut-close
sleep 1
printf '*3\r\n$3\r\nset\r\n$3\r\nfoo\r\n$3\r\noof\r\n' | socat ${debug} ${timeout} - TCP:10.36.44.50:9002,shut-close
sleep 1
printf '*3\r\n$3\r\nset\r\n$3\r\nfoo\r\n$3\r\noof\r\n' | socat ${debug} ${timeout} - TCP:10.36.44.50:9002,shut-close
sleep 1
printf '*3\r\n$3\r\nset\r\n$3\r\nfoo\r\n$3\r\noof\r\n' | socat ${debug} ${timeout} - TCP:10.36.44.50:9002,shut-close
sleep 1
printf '*3\r\n$3\r\nset\r\n$3\r\nfoo\r\n$3\r\noof\r\n' | socat ${debug} ${timeout} - TCP:10.36.44.50:9002,shut-close
sleep 1
#
#
#
#
#
#
#
#
#
#printf '*3\r\n$3\r\nset\r\n$3\r\nbar\r\n$3\r\nrab\r\n' | socat ${debug} ${timeout} - TCP:corn21.stanford.edu:9102,shut-close
#printf '*4\r\n$6\r\nsetbit\r\n$3\r\nfoo\r\n$1\r\n1\r\n$1\r\n1\r\n' | socat ${debug} ${timeout} - TCP:corn21.stanford.edu:9102,shut-close
#printf '*2\r\n$3\r\nget\r\n$3\r\nfoo\r\n' | socat ${debug} ${timeout} - TCP:corn21.stanford.edu:9102,shut-close

done

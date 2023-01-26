pid=$1

if [ $# -eq 0 ]
  then
    echo "Add process ID as an argument."
		exit 1
fi

for (( i=1;i<=31;i++ ))
do
	if (( $i!=9 && $i!=19 ))
	then
		kill -$i $pid
		sleep .01
	fi
done

sleep .5

kill -9 $pid

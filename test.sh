#### -- Random Forest Test -- ####
go build

#### -- Config -- ####
CASES=10


#### -- Print Header -- ####
RESET="\x1b[0m"
BOLD="\x1b[1m"
ITALIC="\x1b[3m"

printf "\E[H\E[2J" ## Clear screen
printf $BOLD
printf $ITALIC
echo "Launching Random Forest Test...$RESET\n"

best_depth=0
best_accuracy=0

#### -- Test Function -- ####
unit_test()
{
	## Initialize
	DEPTH=$1
	cmd="./Random_Forest -d $DEPTH"
	accuracyTrainTotal=0
	accuracyTestTotal=0
	depthMeanTotal=0
	depthMax=0
	timeTotal=0
	test=0

	## Test Loop
	while [ $test -lt $CASES ]
	do
		printf "|\x1b[1m %-2d \x1b[0m| test %d / %d\r" $DEPTH $test $CASES

		output=$(eval "$cmd")

		## Depth
		depthMean=$(echo "$output" | grep Depth | cut -d "|" -f 4)
		depthMeanTotal=$(echo "$depthMeanTotal + $depthMean" | bc)

		depthMaxCase=$(echo "$output" | grep Depth | cut -d "|" -f 5)
		if [ $depthMaxCase -gt $depthMax ]
		then
			depthMax=$depthMaxCase
		fi

		## Accuracy
		accuracyTrain=$(echo "$output" | grep Accuracy | cut -d "|" -f 3)
		accuracyTest=$(echo "$output" | grep Accuracy | cut -d "|" -f 4)

		accuracyTrainTotal=$(echo "$accuracyTrainTotal + $accuracyTrain" | bc)
		accuracyTestTotal=$(echo "$accuracyTestTotal + $accuracyTest" | bc)

		## Time
		time=$(echo "$output" | grep time | cut -d \n -f 203 | cut -d ":" -f 2)
		prefix=$(echo "$time" | rev | cut -c-1-8 | rev | cut -c-1-1)
		if [ "$prefix" = "m" ]
		then
			time_cut=$(echo "$time" | rev | cut -c9-42 | rev)
			time_cut=$(echo "scale = 9; ($time_cut / 1000)" | bc)
		elif [ "$prefix" = "µ" ]
		then
			time_cut=$(echo "$time" | rev | cut -c9-42 | rev)
			time_cut=$(echo "scale = 9; ($time_cut / 1000000)" | bc)
		elif [ "$prefix" = "n" ]
		then
			time_cut=$(echo "$time" | rev | cut -c9-42 | rev)
			time_cut=$(echo "scale = 9; ($time_cut / 1000000000)" | bc)
		else
			time_cut=$(echo "$time" | rev | cut -c9-42 | rev)
		fi
		timeTotal=$(echo "$timeTotal + $time_cut" | bc)

		test=$(($test + 1))
	done

	## Mean
	accuracyTrainMean=$(echo "scale = 7; $accuracyTrainTotal / $CASES" | bc)
	accuracyTestMean=$(echo "scale = 7; $accuracyTestTotal / $CASES" | bc)
	depthMeanTotal=$(echo "scale = 1; $depthMeanTotal / $CASES" | bc)
	timeMean=$(echo "scale = 9; $timeTotal / $CASES" | bc)
	if (( $(echo "$accuracyTestMean > $best_accuracy" | bc -l) ))
	then
		best_accuracy=$accuracyTestMean
		best_depth=$DEPTH
	fi

	printf "|\x1b[1m %-2d \x1b[0m| %-3d | %-4.1f | %-12f | %-12f | %-9f |\n" $DEPTH $depthMax $depthMeanTotal $accuracyTrainMean $accuracyTestMean $timeMean
}


#### -- Print Table -- ####
printf "Test Cases per Depth: %d\n\n" $CASES
echo "+-----------------+-----------------------------+-----------+"
echo "|\x1b[1m Depth           \x1b[0m|\x1b[1m Accuracy Mean               \x1b[0m|\x1b[1m Time Mean \x1b[0m|"
echo "+-----------------+-----------------------------+           |"
echo "|\x1b[1m -d \x1b[0m|\x1b[1m Max \x1b[0m|\x1b[1m Mean \x1b[0m|\x1b[1m Training Set \x1b[0m|\x1b[1m Test Set     \x1b[0m| (Seconds) |"
echo "+----+-----+------+--------------+--------------+-----------+"

#### -- Depth -- ####
depth=1
while [ $depth -lt 11 ]
do
	unit_test $depth
	depth=$(($depth + 1))
done

echo "+----+-----+------+--------------+--------------+-----------+"
echo
printf "Best Depth for Test Set Accuracy: %d\n\n" $best_depth


#### -- Cleanup -- ####
rm Random_Forest

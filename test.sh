#### -- Random Forest Test -- ####
go build

#### -- Print Header -- ####
RESET="\x1b[0m"
BOLD="\x1b[1m"
ITALIC="\x1b[3m"

printf "\E[H\E[2J" ## Clear screen
printf $BOLD
printf $ITALIC
echo "Launching Random Forest Test...$RESET\n"


#### -- Config -- ####
CASES=10


#### -- Test Function -- ####
unit_test()
{
	DEPTH=$1
	# echo "Oh hi!" ##########
	echo $DEPTH
	cmd="./Random_Forest -d $DEPTH"
	accuracyTrainTotal=0
	accuracyTestTotal=0

	test=0
	while [ $test -lt $CASES ]
	do
		# echo $test
		output=$(eval "$cmd")
		accuracyTrain=$(echo "$output" | grep Accuracy | cut -d "|" -f 3)
		accuracyTest=$(echo "$output" | grep Accuracy | cut -d "|" -f 4)
		# accuracy=$(echo "$output" | grep Accuracy) ##########
		# echo $output ##########
		# echo $accuracy ##########
		accuracyTrainTotal=$(echo "$accuracyTrainTotal + $accuracyTrain" | bc)
		accuracyTestTotal=$(echo "$accuracyTestTotal + $accuracyTest" | bc)
		# echo $accuracyTrain
		# echo $accuracyTrainTotal

		# echo $accuracyTest
		test=$(($test + 1))
	done
	accuracyTrainMean=$(echo "scale = 7; $accuracyTrainTotal / $CASES" | bc)
	accuracyTestMean=$(echo "scale = 7; $accuracyTestTotal / $CASES" | bc)
	echo $accuracyTrainMean
	echo $accuracyTestMean
	echo ##########
}

# unit_test 1

#### -- Depth -- ####
depth=1
while [ $depth -lt 10 ]
do
	unit_test $depth
	depth=$(($depth + 1))
done

#### -- Cleanup -- ####
rm Random_Forest

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
	printf "|\x1b[1m "
	printf $DEPTH
	printf "     \x1b[0m| "
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
	printf $accuracyTrainMean
	printf "     | "
	printf $accuracyTestMean
	printf "     |\n"
}

# unit_test 1
echo "        +-----------------------------+"
echo "        |\x1b[1m Accuracy                    \x1b[0m|"
echo "+-------+--------------+--------------+"
printf "|\x1b[1m Depth \x1b[0m|\x1b[1m Training Set \x1b[0m|\x1b[1m Test Set     \x1b[0m|\n"
echo "+-------+--------------+--------------+"
# echo "|%v Metric          %v|%v Training Set  %v|%v Test Set      %v|\n", BOLD, RESET, BOLD, RESET, BOLD, RESET)

#### -- Depth -- ####
depth=1
while [ $depth -lt 10 ]
do
	unit_test $depth
	depth=$(($depth + 1))
done
echo "+-------+--------------+--------------+"
echo

#### -- Cleanup -- ####
rm Random_Forest

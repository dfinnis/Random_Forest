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

best_depth=0
best_accuracy=0

#### -- Test Function -- ####
unit_test()
{
	DEPTH=$1
	cmd="./Random_Forest -d $DEPTH"

	accuracyTrainTotal=0
	accuracyTestTotal=0
	test=0

	while [ $test -lt $CASES ]
	do
		printf "|\x1b[1m "
		printf $DEPTH
		printf "     \x1b[0m| "
		printf "test "
		printf $test
		printf " / "
		printf $CASES
		printf "\r"

		output=$(eval "$cmd")
		accuracyTrain=$(echo "$output" | grep Accuracy | cut -d "|" -f 3)
		accuracyTest=$(echo "$output" | grep Accuracy | cut -d "|" -f 4)

		accuracyTrainTotal=$(echo "$accuracyTrainTotal + $accuracyTrain" | bc)
		accuracyTestTotal=$(echo "$accuracyTestTotal + $accuracyTest" | bc)

		test=$(($test + 1))
	done

	accuracyTrainMean=$(echo "scale = 7; $accuracyTrainTotal / $CASES" | bc)
	accuracyTestMean=$(echo "scale = 7; $accuracyTestTotal / $CASES" | bc)

	if (( $(echo "$accuracyTestMean > $best_accuracy" | bc -l) ))
	then
		best_accuracy=$accuracyTestMean
		best_depth=$DEPTH
	fi

	printf "|\x1b[1m "
	printf $DEPTH
	printf "     \x1b[0m| "
	printf $accuracyTrainMean
	printf "     | "
	printf $accuracyTestMean
	printf "     |\n"
}

#### -- Print Table -- ####
echo "        +-----------------------------+"
echo "        |\x1b[1m Accuracy Mean               \x1b[0m|"
echo "+-------+--------------+--------------+"
printf "|\x1b[1m Depth \x1b[0m|\x1b[1m Training Set \x1b[0m|\x1b[1m Test Set     \x1b[0m|\n"
echo "+-------+--------------+--------------+"

#### -- Depth -- ####
depth=1
while [ $depth -lt 10 ]
do
	unit_test $depth
	depth=$(($depth + 1))
done

echo "+-------+--------------+--------------+"
echo
printf "Best Depth for Test Set Accuracy: "
printf $best_depth
echo
echo


#### -- Cleanup -- ####
rm Random_Forest

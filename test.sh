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


#### -- Test Function -- ####
unit_test()
{
	DEPTH=$1
	# echo "Oh hi!" ##########
	echo $DEPTH
	cmd="./Random_Forest -d $DEPTH"
	output=$(eval "$cmd")
	accuracyTrain=$(echo "$output" | grep Accuracy | cut -d "|" -f 3)
	accuracyTest=$(echo "$output" | grep Accuracy | cut -d "|" -f 4)
	# accuracy=$(echo "$output" | grep Accuracy) ##########
	# echo $output ##########
	# echo $accuracy ##########
	echo $accuracyTrain
	echo $accuracyTest
	echo ##########
}


unit_test 1

depth=1
while [ $depth -lt 10 ]
do
	depth=$(($depth + 1))
	unit_test $depth
done

rm Random_Forest

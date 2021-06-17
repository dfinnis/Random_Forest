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

cmd="./Random_Forest"
output=$(eval "$cmd")
accuracy=$(echo "$output" | grep Accuracy)
# echo $output
echo $accuracy

rm Random_Forest

set -euo pipefail

if [ $# -eq 0 ]
  then
    printf "No arguments supplied. Usage: testpart3.sh [path/to/bishop-sample/ecs150-lab1/linked]\n"
    exit 1
fi

compare_output () {
  MY_OUTPUT=$(part3.out $1 $2  2>&1)
  BISH_OUTPUT=$($BISH_SAMPLE $1 $2  2>&1)
  if [[ "$MY_OUTPUT" == "$BISH_OUTPUT" ]]; then
  printf "\U2714 Passed! Args: $1, $2\n"
  printf "\t+ Expected: $BISH_OUTPUT\n"
  printf "\t+ Saw: $MY_OUTPUT"
  else
  printf "\U2716 Failed! Args: $1, $2\n"
  printf "\t+ Expected: $BISH_OUTPUT\n"
  printf "\t+ Saw: $MY_OUTPUT"
  fi
}

BISH_SAMPLE=$1

printf "Creating testing files..."
make part3
chmod +x part3.out
touch tests/fileA
ln tests/fileA tests/fileB
ln -s tests/fileA tests/fileC
touch tests/fileD
ln -s tests/fileA tests/fileF
ln -s tests/fileF tests/fileG
# generate new files for test cases here

printf "Test part 3\n"

declare -a LISTOFTESTARGS=(
  "tests/fileA,tests/fileC"
  "tests/fileC,tests/fileA"
  "tests/fileA,tests/fileB"
  "tests/fileB,tests/fileA"
  "tests/fileA,tests/fileD"
  "tests/fileA,tests/fileE"
  "tests/fileC,tests/fileF"
  "tests/fileC,tests/fileG"
  #add the file paths here
)
COUNTER=0
for i in "${LISTOFTESTARGS[@]}"
do
  COUNTER=$((COUNTER+1))
  printf "\n=======Test $COUNTER\n"
   arg1=$(echo "$i" | awk -F',' '{print $1}')
   arg2=$(echo "$i" | awk -F',' '{print $2}')
   echo "$(compare_output $arg1 $arg2)"
done

printf "Cleaning up...\n"
rm -f tests/file* part3.out
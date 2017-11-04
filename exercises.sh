option=$1

echo "------------- Getting example $option."


git checkout .
#git pull

  case $option in
	1)
		echo "------------- SUBJECT: Declaring variables, data types."
		git checkout exercise-$option
		;;
	2)
		git checkout exercise-$option
		;;
	*)
		echo "1 2 degil"
		;;
  esac
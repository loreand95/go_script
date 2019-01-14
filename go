#!/bin/bash

# *************************    GO - V.01    ***************************	#
# Script to save and navigate on your list of favourite path		#
#									#
# Tested on:  macOS Mojave 10.14.2					#
#									#
#			      INSTRUCTIONS				#
# 1) Set environment variable $GO_PATH with path of go_path.bak.txt	#
# 	 Example:							#
#    export GO_PATH=/Users/loreand/go/go_path.bak			#
#									#
# 2) Go to directory, cd /.../go					#
# 3) Run command with ". go" or "source go"				#
#									#
#				TIPS					#
# Add script folder to $PATH and add ' alias go=". go" ' to env		#
#									#
# 						Lorenzo Andreoli	#
# ********************************************************************* #


#Function to check configuration script
check_configuration(){
	#Is $GO_PATH env?
	if [ -z $GO_PATH ]
		#NO
		then
		echo '-ERROR: $GO_PATH NOT FOUND!'
		echo -e "-TIP: Check if $GO_PATH is add on environment variable :)\n"
		
		return 1
		#Is $GO_PATH valid?
		elif [ ! -e $GO_PATH ]
			then
			echo -e "-ERROR: $GO_PATH NOT EXIST!"
			return 1
	fi
}

#Function to move path by input number
#$1 = number of row selected
movetopath(){
	path=$( awk "NR==$1" $GO_PATH )
	cd $path
}


#Function to remove path by input number
#$1 = number of row selected
removepath(){
	sed -i '' $1d $GO_PATH

	# sed command Linux, try:
	# sed -i $1d $GO_PATH
}

go_path(){
	# 1 parameter is empty?
	if [ -z $1 ]
		#Yes, see favourite list
		then
			#Print and enumerate
			cat $GO_PATH | nl
			#Get input choise
			read -p '> ' choise
			
			if [[ $choise == *"rm"* ]]
				then
					read -p 'Remove line: ' select_line
					removepath $select_line
				
				elif [[ $choise =~ ^-?[0-9]+$ ]]
				then
					movetopath $choise
			fi


		#No, the parameter is 'now', save path
		elif [ $1 = "now" ]
			then

			#Get number lines
			let number_line=$(( $( cat $GO_PATH | wc -l ) +1))
			#save list
			echo "$( pwd )">> $GO_PATH
			#print path saved
			echo "Save path: $( pwd ) as $number_line" 


		#No, the parameter is a number option
		elif [[ $1 =~ ^-?[0-9]+$ ]]
			then
			movetopath $1 #go to path
		
		#No, input not recognized
		else
			echo "some is wrong :("
	fi
}


# Main function
main(){
	#Check $GO_PATH exist
	check_configuration

	if [ $? -eq 0 ]
		#Right configuration :)
		then
			go_path "$@"
		else
		#Bad configuration :(
		echo -e "EXIT - How to run the script? More details on www.github.com/loreand95/Go_script\n"	
	fi
}

main "$@"

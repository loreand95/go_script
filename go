#!/bin/bash

# *************************    GO - V.01    ************************* #
# Script to save and navigate on your list of favourite path 		  #
#																	  #
# Tested on:  macOS Mojave 10.14.2 			   						  #
#																	  #
#							  INSTRUCTIONS							  #
# 1) Set environment variable $GO_PATH with path of go_path.bak.txt   #
# 	 Example:														  #
#    export GO_PATH=/Users/loreand/go/go_path.bak					  #
# 	  																  #
# 2) Go to directory, cd /.../go 									  #
# 3) Run command with ". go" or "source go"							  #
#					  												  #
# 								  TIPS       						  #
# Add script folder to $PATH and add ' alias go=". go" ' to env	      #
#																	  #
# 												 Lorenzo Andreoli     #
# ******************************************************************* #


#Function to check configuration script
check_configuration(){

	if [ -z $GO_PATH ]
		then
		echo '-ERROR: $GO_PATH NOT FOUND!'
		echo -e '-TIP: Check if $GO_PATH is add on environment variable :)\n'
		
		return 1

		elif [ ! -e $GO_PATH ]
			then
			echo -e "-ERROR: $GO_PATH NOT EXIST!"
			return 1
	fi
}

#Function to move path by input number
movetopath(){
	path=$( awk "NR==$1" $GO_PATH )
	cd $path
}


#Function to remove path by input number
removepath(){
	sed -i '' $1d $GO_PATH

	# sed command Linux, try:
	# sed -i $1d $GO_PATH
}

go(){

	#Check exist first parameter
	if [ -z $1 ]
		#No parameter, see favourite list
		then
			#TODO check if exist favourite.txt
			cat $GO_PATH | nl
			read -p '> ' choise
			
			if [[ $choise == *"rm"* ]]
				then
					read -p 'Remove line: ' select_line
					removepath $select_line
				
				elif [[ $choise =~ ^-?[0-9]+$ ]]
				then
					movetopath $choise
			fi


		#The parameter 'now', save path
		elif [ $1 = "now" ]
			then

			#Get number lines
			let number_line=$(( $( cat $GO_PATH | wc -l ) +1))
		
			echo "$( pwd )">> $GO_PATH
			echo "Save path: $( pwd ) as $number_line" 


		#The parameter is a number option
		elif [[ $1 =~ ^-?[0-9]+$ ]]
			then
			movetopath $1 #go to path

		else
			echo "some is wrong :("
	fi
}


# Main function
main(){
	#Check configuration
	check_configuration

	if [ $? -eq 0 ]
		then
			go "$@"
		else
		echo -e "EXIT - How to run the script? More details on www.github.com/loreand95/Go_script\n"
		
	fi
}

main "$@"
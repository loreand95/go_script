#!/bin/bash

# *************************    GO - V.01    ***************************	#
# Script to save and navigate on your list of favourite path            #
#                                                                       #
# Tested on:  macOS Mojave 10.14.2                                      #
#                                                                       #
#                                           Lorenzo Andreoli            #
# ********************************************************************* #

PATH_GO=$(which go)
DIR_GO=$(dirname "$PATH_GO")
export PATH_GO_LIST="$DIR_GO/go_list.bak"



#Function to check configuration script
check_configuration(){
    #Is $PATH_GO_LIST env?
    if [ -z $PATH_GO_LIST ]
    #NO
    then
        echo '-ERROR: $PATH_GO_LIST NOT FOUND!'
        echo -e "-TIP: Check if $PATH_GO_LIST is add on environment variable :)\n"
        
        return 1
        #Is $PATH_GO_LIST valid?
    elif [ ! -e $PATH_GO_LIST ]
    then
        echo -e "-ERROR: $PATH_GO_LIST NOT EXIST!\n"
        return 1
    fi
}

#Function to move path by input number
#$1 = number of row selected
movetopath(){
    path=$(awk "NR==$1" $PATH_GO_LIST | awk '{print $2}') # select first column
    cd $path
}


#Function to remove path by input number
#$1 = number of row selected
removepath(){
    if [[ $1 =~ ^-?[0-9]+$ ]]
    then
        sed -i '' $1d $PATH_GO_LIST
        echo -e "The line n.$1 has been removed\n"
    else
        echo -e "No lines have been removed\n"
    fi
}

# Function to find line number by input
# $1: path name to find
findPath(){

    line_number=$( cut -d ' ' -f 1 $PATH_GO_LIST | grep -w -n $1 $PATH_GO_LIST | cut -d : -f 1 | head -n 1 )
    echo "$line_number"
}

go_list(){
    # 1 parameter is empty?
    if [ -z $1 ]
    #Yes, see favourite list
    then
        #Print and enumerate
        cat $PATH_GO_LIST | nl
        #Get input choise
        read -p "> " choise opt1
        
        if [[ $choise == "rm" ]]
        then
            removepath $opt1
            
        elif [[ $choise =~ ^-?[0-9]+$ ]]
        then
            movetopath $choise
            
        fi
        
        
    #No, the parameter is 'now', save path
    elif [ $1 = "save" ]
    then
        if [ -z $2 ]
        then
            read -p "Save path as: " title
        else
            title=$2
        fi

        exist=$( findPath $title )
        
        while [ ! -z $exist ]
        do
            read -p "Title \"$title\" exist! Choose another title:" title
            exist=$( findPath $title )
        done
        
        #Get number lines
        let number_line=$(( $( cat $PATH_GO_LIST | wc -l ) +1))
        #Save list
        echo -e "$title\t\t$( pwd )">> $PATH_GO_LIST
        #Print path saved
        echo "Save path: $( pwd ) as $number_line"

        
        
    #No, the parameter is a number option
    elif [[ $1 =~ ^-?[0-9]+$ ]]
    then
        movetopath $1 #Go to path
        
    #No, the parameter is a string, find path
    else
        #get line of path by title
        line=$( findPath $1 )

        if [ -z $line ]
        then
            echo "\"$1\" - Title not found!"
        else
        movetopath $line
        fi
    fi
}


# Main function
main(){
    
    #Check $PATH_GO_LIST exist
    check_configuration
    
    if [ $? -eq 0 ]
    #Right configuration :)
    then
        go_list "$@"
    else
        #Bad configuration :(
        echo -e "EXIT - How to run the script? More details on www.github.com/loreand95/Go_script\n"
    fi
}

main "$@"

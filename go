#!/bin/bash

# *************************    GO - V.01    ***************************	#
# Script to save and navigate on your list of favourite path            #
#                                                                       #
# Tested on:  macOS Mojave 10.14.2                                      #
#                                                                       #
#INSTRUCTIONS                                                           #
# 1) Go to directory, cd /go/to/path/Go_script                          #
# 2) Run command with ". go" or "source go"                             #
#                                                                       #
#TIPS                                                                   #
# Add 'alias go=". go"' to env to run "go" everywhere                   #
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
        echo '\n-ERROR: $PATH_GO_LIST NOT FOUND!'
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
    path=$( awk "NR==$1" $PATH_GO_LIST )
    cd $path
}


#Function to remove path by input number
#$1 = number of row selected
removepath(){
    sed -i '' $1d $PATH_GO_LIST
    
    # sed command Linux, try:
    # sed -i $1d $PATH_GO_LIST
}

PATH_GO_LIST(){
    # 1 parameter is empty?
    if [ -z $1 ]
    #Yes, see favourite list
    then
        #Print and enumerate
        cat $PATH_GO_LIST | nl
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
        let number_line=$(( $( cat $PATH_GO_LIST | wc -l ) +1))
        #save list
        echo "$( pwd )">> $PATH_GO_LIST
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
    
    #Check $PATH_GO_LIST exist
    check_configuration
    
    if [ $? -eq 0 ]
    #Right configuration :)
    then
        PATH_GO_LIST "$@"
    else
        #Bad configuration :(
        echo -e "EXIT - How to run the script? More details on www.github.com/loreand95/Go_script\n"
    fi
}

main "$@"

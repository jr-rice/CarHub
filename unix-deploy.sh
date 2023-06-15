#!/bin/sh

# Constants
arg1=$1
arg2=$2
arg_count=$#

deploy() {
    read -p "Database Username: " user
    read -p "Database Password: " pass
    read -p "Database Hostname (Leave empty if SQL database is hosted locally): " host
    echo "Deploying application..."

    if [ "x$arg2" = "x" ]
    then
        echo "Building in default directory: /var/www/html"
        vite build --out-dir /var/www/html
    else
        echo "Building in selected directory: $arg2"
        vite build --out-dir $arg2
    fi

    if [ "x$host" = "x" ]
    then
        echo "Attempting to run Gin-Gonic Back End, press Ctrl + C to exit"
        go run src/data/main.go $user $pass "localhost"
    else
        echo "Attempting to run Gin-Gonic Back End, press Ctrl + C to exit"
        go run src/data/main.go $user $pass $host
    fi
}

eval_num_of_args() {
    case "$arg_count" in
        1)
            read_arg
        ;;
        2)
            read_arg
        ;;
        *)
            if [ $arg_count -gt 2 ]
            then
                echo "ERROR: Too many arguments provided. Exiting UNIX Setup Script."
            else
                echo "ERROR: Too little arguments provided. Exiting UNIX Setup Script."
            fi
            exit 1
        ;;
    esac
}

read_arg() {
    case "$arg1" in
        "deploy")
            deploy
        ;;
        "help")
            echo "`basename ${0}` - usage:"
            echo "deploy - Builds and deploys the app on your web server"
            echo "help - Brings up this help menu again"
            exit 0
        ;;
        *)
            echo "ERROR: Invalid argument ${arg1} provided. Exiting UNIX Setup Script."
            exit 1
        ;;
    esac
}
#!/bin/sh

pkgs="typescript react react-dom react-router-dom vite @vitejs/plugin-react @types/react @types/react-dom @types/react-router-dom"

eval_num_of_args() {
    case "$#" in
        1)
            read_arg
        ;;
        *) 
            echo "ERROR: Too many or too little arguments provided. Exiting UNIX Setup Script."
            exit 1
            ;;
    esac
}

go_mod_tidy() {
    if [ -x go ]
    then
        go mod tidy
    else
        echo "Could not find Go."
        echo "Please install Go then re-run the UNIX Setup Script again."
        exit 1
    fi
}

install_npm() {
    if [ -x npm ]
    then
        npm install --save-dev $pkgs
        go_mod_tidy
    else
        echo "Could not find NPM."
        echo "Please install NPM then re-run the UNIX Setup Script again."
        exit 1
    fi
}

install_yarn() {
    if [ -x yarn ]
    then
        yarn add $pkgs --dev
        go_mod_tidy
    else
        echo "Could not find Yarn."
        echo "Please install Yarn then re-run the UNIX Setup Script again."
        exit 1
    fi
}

read_arg() {
    case "$1" in
        "npm")
            install_npm
        ;;
        "yarn")
            install_yarn
        ;;
        "help")
            echo "`basename ${0}`:usage:"
            echo "npm - Install the dependencies with NPM"
            echo "yarn - Install the dependencies with Yarn"
            echo "help - Brings up this help menu again"
            exit 0
        ;;
        *)
            echo "ERROR: Invalid argument provided! Exiting UNIX Setup Script."
            exit 1
            ;;
    esac
}

eval_num_of_args
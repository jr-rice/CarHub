# CarHub Project
This guide has been written to better explain the setup of the project, project deployment, and the structure of the application

## Requirements
- node.js >=18.14.0
- npm >=9.7.1
- go >=1.20
- python >=3.9 (Windows Only)

## Deploying
The project was built with both a Gin-Gonic Back End and a ReactJS/TypeScript Front End, so we will first need to build the Front End, this will require setting up our project file with our necessary packages. This project has provided easy-to-use shell scripts in order to streamline a fast and simple setup of the necessary packages and programs necessary to run the application.

The npm packages and Go modules can be installed on Windows using the windows-setup Python script for the respective package manager you wish to use for the project, they can be either executed from the GUI itself or ran from Command Prompt/Windows Terminal like so:
```
python windows-setup.py
```

The npm packages and Go modules can be installed on MacOS/Linux using the unix-setup script with the following syntax:
```
./unix-setup.sh [package manager]
```

The setup scripts should install the Vite build tool with it, which allows for building the app to be deployed.

App can be deployed on Windows with windows-deploy Python script either from the GUI or within Command Prompt/Windows Terminal like so:
```
python windows-deploy.py
```

App can be deployed on MacOS/Linux using the unix-deploy script with the following syntax:
```
./unix-deploy.sh [command] [argument]
```
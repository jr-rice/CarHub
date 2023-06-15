import os
import subprocess
import sys

def main():
    if os.name != "nt":
        print("Error: This script is intended for Windows use only!")
        sys.exit(1)
    else:
        username = input("Database Username: ")
        password = input("Database Password: ")
        hostname = input("Database Hostname (Leave empty if SQL database is being hosted locally): ")
        directory = input("Build Directory: ")
        print("Deploying application")

        try:
            if directory == "":
                print("Error: No build directory specified! Please enter the directory that your web server uses for hosting web applications and re-run the deploy script!")
                os.system("pause")
                sys.exit(1)
            else:
                subprocess.call(["vite", "build", "--out-dir", directory])
                if hostname == "":
                    print("Attempting to run Gin-Gonic Back End, press Ctrl + C to exit")
                    subprocess.call(["go", "run", "src\data\main.go", username, password, "localhost"])
                else:
                    print("Attempting to run Gin-Gonic Back End, press Ctrl + C to exit")
                    subprocess.call(["go", "run", "src\data\main.go", username, password, hostname])
                sys.exit(0)
        except FileNotFoundError:
            print("Error: Go is not installed! Please install Go and re-run the deploy script!")
            os.system("pause")
            sys.exit(1)

if __name__ == "__main__":
    main()
import os
import subprocess
import sys

def main():
    pkgs = "typescript react react-dom react-router-dom vite @vitejs/plugin-react @types/react @types/react-dom @types/react-router-dom"

    if os.name != "nt":
        print("Error: This script is intended for Windows use only!")
        sys.exit(1)
    else:
        try:
            subprocess.call(["yarn", "add", pkgs, "--dev"])
            print("Success: yarn install complete! Collecting Go modules...")
        except FileNotFoundError:
            print("Warning: yarn was not found, defaulting to npm install")
            try:
                subprocess.call(["npm", "install", "--save-dev", pkgs])
                print("Success: npm install complete! Collecting Go modules...")
            except FileNotFoundError:
                print("Error: No package manager found! Please install npm or yarn then re-run the script!")
                os.system("pause")
                sys.exit(1)

        try:
            subprocess.call(["go", "mod", "tidy"])
            print("Success: Go modules collected! Project is ready for deployment!")
            os.system("pause")
            sys.exit(0)
        except FileNotFoundError:
            print("Error: Go is not installed! Please install Go and re-run the setup script!")
            os.system("pause")
            sys.exit(1)

if __name__ == "__main__":
    main()
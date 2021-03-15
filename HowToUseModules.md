## Work with modules
0. inside the folder of your project run those commands
1. initialization
    ```bash 
    go mod init github.com/YourUsername/title-of-the-module
    ```
    - it creates a go.mod file that is similar to package.json
2.  adding and downloading dependencies
    ```bash 
    go get github.com/YourUsername/title-of-the-module
    ```
    - it will be download in $GOPATH/pkg folder
    - the dependencies informations are added in go.mod
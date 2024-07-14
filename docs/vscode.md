# visual studio code

* go from go team at google
* go test explorer from prem parihar

# install code tools

1. type shift ctrl p
2. filter on shell
3. click install 'code' command in PATH

# give it a test drive

1. open terminal
2. navigate browser to `https://github.com/cloudacademy/godemo`
3. in the terminal 
4. cd sample_repo
5. `git clone git@github.com:cloudacademy/godemo.git`

# common variables

* ${workspaceFolder} the path of the folder opened in VS Code
* ${file} the current opened file
* ${relativeFile} the current opened file relative to workspaceRoot
* ${fileBasename} the current opened file's basename
* ${fileDirname} the current opened file's dirname
* ${fileExtname} the current opened file's extension
* ${cwd} the task runner's current working directory on startup


# Notes on debugger and building 

workflow from [here](https://stackoverflow.com/questions/70594682/go-build-with-debug-mode-in-visaul-studio-code)

1. install Delve (debugger for golang)
```
go install github.com/go-delve/delve/cmd/dlv@latest
```
2. install dDelve (gui for delve)
```
go install github.com/aarzilli/gdlv@latest
```
3. Restart vs code


```
$ ~/go/bin/dlv debug
```

## vim extension

```
defaults write com.microsoft.VSCode ApplePressAndHoldEnabled -bool false
defaults write com.microsoft.VSCodeInsiders ApplePressAndHoldEnabled -bool false 
defaults write com.vscodium ApplePressAndHoldEnabled -bool false 
defaults write com.microsoft.VSCodeExploration ApplePressAndHoldEnabled -bool false
defaults delete -g ApplePressAndHoldEnabled 
```
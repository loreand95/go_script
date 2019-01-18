# Go_script
Script to save and navigate on your list of favourite path

## Install on OSX
Set environment variable, add folder to $PATH
```
$ cd ~
$ echo 'export PATH=$PATH: [path] /Go_script' >> .bash_profile
```
Add alias to script
```
$ echo 'alias go=". go"' >> .bash_profile
```
## Run

### Run and view favourite path list
```
$ go
```
![](https://github.com/loreand95/Go_script/blob/master/images/go-move.gif)

### Save favourite path
```
$ cd [path]
$ go now
```
![](https://github.com/loreand95/Go_script/blob/master/images/go-save.gif)

### Remove path
```
$ go
$ rm
$ number of path
```
![](https://github.com/loreand95/Go_script/blob/master/images/go-remove.gif)

## Author
Lorenzo Andreoli

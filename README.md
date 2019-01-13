# Go_script
Script to save and navigate on your list of favourite path

## Install
Set environment variable $GO_PATH with path of go_path.bak

```
> cd /Users/username
> echo 'export GO_PATH=/...custom path.../go_path.bak' >> .bash_profile
```
Add folder to $PATH
```
> echo 'export PATH=$PATH:/.../Go_script-master' >> .bash_profile
```
Add alias to script	
```
> echo alias go=". go" >> .bash_profile
```
## Run

### Run and view favourite path list
```
> go
```
![](https://github.com/loreand95/Go_script/blob/master/images/go-move.gif)

### Save favourite path
```
> cd / ... / path
> go now
```
![](https://github.com/loreand95/Go_script/blob/master/images/go-save.gif)

### Remove path
```
> go
> rm
> number of path
```
![](https://github.com/loreand95/Go_script/blob/master/images/go-remove.gif)

## Author
Lorenzo Andreoli

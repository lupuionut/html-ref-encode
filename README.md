Simple tool to encode your text using numeric character references.
### Installation
```
go build main.go
```

### Usage
Run the binary specifying the text to encode. 
For example:
```
$ ./main "Do you like Äripäev?"
```
it will output
```
$ Do you like &#xC4rip&#xE4ev?
```

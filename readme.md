# CSV Splitter

## How to Build

```powershell
go get github.com/estenrye/csvsplitter
cd "$env:GOPATH/src/github.com/estenrye/csvsplitter"
go build
```

## How to Use

### Splitting a CSV with a header row on the first column.
```powershell
cd "$env:GOPATH/src/github.com/estenrye/csvsplitter"
csvsplitter.exe -hasHeaderRow -splitColumn=1 .\testFiles\test.csv
```

### Splitting a CSV with no header row on the first column.
```powershell
cd "$env:GOPATH/src/github.com/estenrye/csvsplitter"
csvsplitter.exe -splitColumn=1 .\testFiles\test_noheaderrow.csv
```
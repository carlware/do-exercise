# Command line tool to extract gps latitude and longitude from images

### Usage:
```
--input:  filename or directory name 
--output: name of the output file
```

### Example
```shell
./exif --input images
```

## Development
New tags processor can be added and new writer can be added implementing `TagHandleFunc` and  `RecordWriter` respectively.  

```shell
# build
make build

# test
make test
```

## TODO
1. increase test coverage
2. add lazy loader
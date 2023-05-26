# Command line tool to extract gps latitude and longitude from images

### Usage:
```
--input:  filename or directory name 
--output: name of the output file
```

### Example
```shell
# build
make build

# process images folder
./exif --input images

# process one file
./exif --input images/bird.jpeg

# process directory with html format
./exif --input images --format html
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
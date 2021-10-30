# Howto

Copy any TTF font into the working directory and call the file `font.ttf`. Then
run the program:

    $ go run cmd/verzettler.go

The output `example.pdf` will be written to the current working directory.

# TODO

- accept command line arguments
    - `rows` and `cols`
    - `font` (TTF file)
    - argument for data file (JSON) containing the pairs

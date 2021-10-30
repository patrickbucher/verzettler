# Verzettler

Verzettler turns key-value pairs provided in the JSON format into flash cards
ready for duplex printing.

## Howto

Build the program:

    $ go build cmd/verzettler.go

Copy any TTF font into the working directory and call the file `font.ttf`. Then
run the program:

    $ ./verzettler pairs.json

The output `flashcards.pdf` will be written to the current working directory.

Get help on the command line options:

    $ ./verzettler -h

Use a special font:

    $ ./verzettler -font /usr/share/fonts/TTF/DejaVuSans.ttf pairs.json

Use a bigger font size:

    $ ./verzettler -size 18 pairs.json

Use a different row/column layout in the output:

    $ ./verzettler -rows 5 -cols 3 pairs.json

Use a different output path:

    $ ./verzettler -out vocabulary.pdf pairs.json

# Verzettler

Verzettler turns key-value pairs provided in the JSON format into flash cards
ready for duplex printing.

## Zettelserver

Run it:

    $ go run server/zettelserver.go

Then browse [localhost:3771/index.html](http://localhost:3771/index.html).

Alternatively, use `curl` or the like:

    $ curl -X POST -F pairs=@pairs.json -F cols=3 -F rows=5 -F size=14 \
      "http://localhost:3771/zettel" --output flashcards.pdf

## Command Line Tool

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

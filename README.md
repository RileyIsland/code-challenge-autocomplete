# Code Challenge: Autocomplete Service

> A web service using Golang that can return autocomplete suggestions based on word fragments, using a local text file
> as the data source.

## Instructions for Mac

- Clone the repository using Git to your local filesystem.
- Using the command line, navigate to the code folder.
- Run `go run main.go`.
- Navigate to http://localhost:9000/autocomplete?term=TERM.
- Replace "TERM" in the url above with any other matches you would like to autocomplete.

## Documentation

- See the `examples` folder for sample JSON results
- Sources:
    - Text file parsing by word: http://networkbit.ch/read-text-file-in-go/
    - Map sorting: https://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values
    - Substrings: https://www.dotnetperls.com/substring-go

## Known Issues/TODOs

- Create a module for matches and move code from main.go into the module for cleanliness purposes
- Clean more words from the initial word results
    - Visit http://localhost:9000/autocomplete?term=brother: "brother", "brother-" and "brother--" are separate results)
    - Get a list of words for which `isNonWord` returns true and identify any other caveats that could be added to
      properly include/exclude characters/words

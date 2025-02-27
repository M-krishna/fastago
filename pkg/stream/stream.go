package stream

import (
	"bufio"
	"io"
)

// Line represents a line in a fasta file
// IsName is true if the first character of the line is '>'
type Line struct {
	Line   string
	IsName bool
}

// StreamFasta return Line objects for each line of the fasta in a channel
func StreamFasta(input io.Reader, output chan Line, errs chan error) {
	defer close(output)
	defer close(errs)

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) != 0 && line[0] == '>' {
			output <- Line{Line: line, IsName: true}
		} else {
			output <- Line{Line: line, IsName: false}
		}

	}

	if err := scanner.Err(); err != nil {
		errs <- err
	}

	errs <- nil
}

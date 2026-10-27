GoPulse

GoPulse is a lightweight command-line website status checker built with Go.

It checks multiple websites concurrently and displays their HTTP status code, availability, and response time.

Features

- Check multiple websites at once
- Concurrent requests using Go goroutines (limited to 10 at a time)
- HTTP status code detection
- Response time measurement
- Detect unavailable websites
- Accept URLs directly from the command line
- Load websites from a text file
- Automatic HTTPS handling

Requirements

- Go 1.20 or newer

Installation

Clone the repository:

git clone https://github.com/jimmmzzzzz/GoPulse.git

Move into the project directory:

cd GoPulse

Usage

Check websites directly

go run main.go google.com github.com openai.com

Example output:

GoPulse
------------------------------
[200] https://google.com             ONLINE  320ms
[200] https://github.com             ONLINE  451ms
[200] https://openai.com             ONLINE  610ms

Use a URL file

Add websites to "urls.txt":

google.com
github.com
openai.com
example.com

Then run:

go run main.go

Build

Compile GoPulse into an executable:

go build -o gopulse

Run it:

./gopulse google.com github.com

On Windows:

gopulse.exe google.com github.com

How It Works

GoPulse creates a goroutine for each website being checked.

Each goroutine sends an HTTP request independently, allowing multiple websites to be checked concurrently instead of waiting for each request to finish one by one.

The results are passed back through a Go channel and displayed in the terminal.

Technologies

- Go
- "net/http"
- Goroutines
- Channels
- WaitGroups

License

MIT License

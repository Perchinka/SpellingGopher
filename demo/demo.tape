Output demo/demo.gif

Require go

Set Shell "bash"
Set FontSize 20
Set Width 1400
Set Height 600
Set Padding 20
Set Theme "Dracula"
Set TypingSpeed 75ms

Hide
Type "go build -o /tmp/spelling-gopher ./cmd/spelling-gopher"
Enter
Sleep 3s
Type "clear"
Enter
Type "/tmp/spelling-gopher --csv demo/quotes.csv"
Enter
Sleep 1500ms
Show

Sleep 800ms
Type "Programs must be written for people to rea"
Sleep 150ms
Type "b"
Sleep 450ms
Backspace
Type "d, and only incidentally for machines to execute."

Sleep 2500ms
Escape
Sleep 500ms

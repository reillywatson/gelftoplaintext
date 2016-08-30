gelftoplaintext is a small utility to extract the messages from GELF logs, because double-encoded JSON is a terrible thing to look at.

Installation:
go get github.com/reillywatson/gelftoplaintext

Usage:
`gelftoplaintext file.gelf`

Or:
`cat file.gelf | gelftoplaintext`

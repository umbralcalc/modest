package renderer

import "fmt"

// RenderDoc renders to valid html doc.
func RenderDoc(code []fmt.Stringer) string {
	htmlDoc := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <script src="https://cdn.plot.ly/plotly-latest.min.js"></script>
</head>
<body>
`
	for _, c := range code {
		htmlDoc += c.String() + `
		`
	}
	htmlDoc += `
</body>
</html>`
	return htmlDoc
}

# templates

HTML, oas3 templates.

### Example

#### Load template example:

```go
package pkg

import (
    "html/template"
    "log"
)

func loadTemplate(tempDir string) *template.Template {
    templatePath := tempDir + "file-form.html"
    t, err := template.ParseFiles(templatePath)
    if err != nil {
        // handle error
    }
    return t
}

```
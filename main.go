cat > main.go << 'EOF'
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "OK")
    })
    fmt.Println("AppDex backend rodando na porta 8081")
    http.ListenAndServe(":8081", nil)
}
EOF


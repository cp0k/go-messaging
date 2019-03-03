http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if pusher, ok := w.(http.Pusher); ok {
            // Push is supported.
            if err := pusher.Push("/app.js", nil); err != nil {
                log.Printf("Failed to push: %v", err)
            }
        }
    })

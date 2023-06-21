package server

// Init initialize server
func Init() error {
    r, err := NewRouter()
    if err != nil {
        return err
    }
    r.Run(":8080")
    return nil
}
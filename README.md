### Project Structure
```
/habit-tracker
    /api                 # API route handlers and middleware  
    /cmd                 # Main applications for this project  
        /habittracker    # Main package for the habit tracker application  
            main.go      # Entry point for the application  
    /config              # Configuration files and logic  
    /internal            # Private application and library code  
        /habit           # Business logic for habit-related operations  
        /user            # Business logic for user-related operations  
    /migrations          # Database migrations  
    /models              # Data models and GORM structs  
    /repository          # Database interactions and GORM integration  
    /util                # Utility functions and common helpers  
```

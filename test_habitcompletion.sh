curl http://localhost:8080/api/habits/completions \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"ID": 2,"HabitID": 90,"Date": "2024-03-12T00:00:00Z"}'

curl http://localhost:8080/api/habits/1/completions \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"ID": 12,"HabitID": 10,"Date": "2024-03-10T00:00:00Z"}'

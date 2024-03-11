curl http://localhost:8080/api/habits \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"ID": 10,"Name": "test habit","UserID": "1","StartDate": "2024-03-10T00:00:00Z","EndDate": "2024-03-10T00:00:00Z","Completions": []}'

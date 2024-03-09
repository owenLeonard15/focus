curl http://localhost:8080/api/habits \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"ID": 5,"Title": "focus","Description": "workout"}'

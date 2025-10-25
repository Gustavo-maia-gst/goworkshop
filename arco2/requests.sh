#!/bin/bash
NUM_REQUESTS="$1"
PORT="${2:-8000}"
URL="http://localhost:$PORT"

uris=("chuck" "pokemon" "rickandmorty")
tmpfile="results"

start_total=$(date +%s.%N)

for i in $(seq 0 $((NUM_REQUESTS - 1))); do
    index=$(( i % 3 ))
    uri="${uris[$index]}"

    (
        curl -s "$URL/$uri" > /dev/null
        echo "Request $i -> $uri completa"
    ) &
done

wait

end_total=$(date +%s.%N)
total=$(echo "$end_total - $start_total" | bc)

echo "------------------------------"
echo "Todas as requisições concluídas!"
echo "Tempo total: ${total}s"

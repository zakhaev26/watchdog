#!/bin/bash

echo "Firing 4 threads."

# Run frequent_consumer in the background
(cd frequent_consumer/_elastic && go run main.go) &

# Run critical_consumer in the background
(cd critical_consumer/_elastic && go run main.go) &

# Run critical_producer in the background
(cd critical_producer && go run main.go) &

# Run frequent_producer in the background
(cd frequent_producer && go run main.go) &


# Wait for both processes to finish
wait

echo "Both processes have finished."

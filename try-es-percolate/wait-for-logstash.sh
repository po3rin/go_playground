#!/bin/bash

## Or whatever command is used for checking logstash availability
until curl 'http://kibana:5601' > /dev/null; do
  echo "Waiting for kibana..."
  sleep 3;
done

# Start your server
fresh

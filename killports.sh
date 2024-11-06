#!/bin/bash

function killport() {
    lsof -i TCP:$1 | grep LISTEN | awk '{print $2}' | xargs kill -9 
}

killport 4000
killport 5001
killport 50051
killport 8080


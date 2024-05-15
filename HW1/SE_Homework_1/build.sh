#!/bin/bash

# Controlla se la randomizzazione degli indirizzi e' attiva se si viene disabilitata
if [ $(cat /proc/sys/kernel/randomize_va_space) == 2 ]; then
	echo "Randomization address disabled"
	echo 0 | sudo tee /proc/sys/kernel/randomize_va_space
fi

# costruisce un container
podman build --tag debian:homework -f Dockerfile

#!/bin/bash
sudo docker compose down --remove-orphans
sudo docker compose build
sudo docker compose up -d
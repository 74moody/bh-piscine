#! /bin/bash


curl -s https://learn.reboot01.com/assets/superhero/all.json | jq '.[] | select(.id == 170).name, .powerstats.power, .apperance.gender'

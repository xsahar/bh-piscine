 curl -s https://learn.reboot01.com/assets/superhero/all.json | jq ' . [] | select (.id == 70) | .name'

#!/bin/bash

if [[ "$#" -ne 1 ]]; then
  echo "Provide 1 arg"
  exit 1 
fi

target="$1"

if [[ "${target}" != "dev" && "${target}" != "prod" ]]; then
  echo "Arg should be dev/prod"
  exit 2
fi

if [[ "${target}" == "dev" ]]; then
  #  doesn't work, issues with
  #  - tailwind watch
  #  - templ watch
  npx tailwindcss -i static/css/style.css -o static/css/tailwind.css --watch & \
  TEMPL_EXPERIMENT=rawgo templ generate -watch -proxy=http://localhost:8080 -open-browser=false & \
  air
elif [[ "${target}" == "prod" ]]; then
  npx tailwindcss build -i static/css/style.css -o static/css/tailwind.css -m
  TEMPL_EXPERIMENT=rawgo templ generate
  go build -o bin/main cmd/goshop/main.go && ./bin/main
fi

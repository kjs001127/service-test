#!/bin/sh

if [ "$CH_ENTRY" = "admin" ]; then
  exec /ch-app-store/admin
elif [ "$CH_ENTRY" = "general" ]; then
  exec /ch-app-store/general
else
  echo "INVALID CH_ENTRY VALUE : $CH_ENTRY"
  exit 1
fi

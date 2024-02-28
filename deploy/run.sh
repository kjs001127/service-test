#!/bin/sh

if [ "$CH_ENTRY" = "desk" ]; then
  exec /ch-app-store/desk
elif [ "$CH_ENTRY" = "admin" ]; then
  exec /ch-app-store/admin
elif [ "$CH_ENTRY" = "front" ]; then
  exec /ch-app-store/front
elif [ "$CH_ENTRY" = "general" ]; then
  exec /ch-app-store/general
else
  echo "INVALID CH_ENTRY VALUE : $CH_ENTRY"
  exit 1
fi

#!/bin/sh

if [ "$CH_ENTRY" = "desk" ]; then
  exec /chx-backend/desk
elif [ "$CH_ENTRY" = "admin" ]; then
  exec /chx-backend/admin
elif [ "$CH_ENTRY" = "front" ]; then
  exec /chx-backend/front
elif [ "$CH_ENTRY" = "general" ]; then
  exec /chx-backend/general
else
  echo "INVALID CH_ENTRY VALUE : $CH_ENTRY"
  exit 1
fi

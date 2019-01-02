#!/bin/bash
cd webapp
npm run build
cp -r ./dist/static/   ../static/
cp ./dist/index.html   ../views/
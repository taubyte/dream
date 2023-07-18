#!/bin/bash

mkdir static || true
cd static
if test -d odval
then
    (
        cd odval
        git pull
        cd -
    )
else
    git clone git@bitbucket.org:taubyte/odval.git
fi
cd -

# build
(
    cd static/odval
    yarn install
    yarn build
    cd -
)

cd static/odval/
zip -9 -r ../../ui.zip dist
cd -

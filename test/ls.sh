#!/bin/bash

echo "Launch it from the directory where is encrypted.arafs or when your encrypt fs is"
echo "use: ./ls.sh <path>"
# example cd arfs && ../ls.sh test/mytestfolder
darkenpath=$(adret darkenpath -key="toto" $1)
ubac_result=$(ubac ls -path=encrypted.arafs $darkenpath)
adret decryptls -key="toto" $ubac_result
#!/bin/bash

echo "Launch it from the directory where is encrypted.arafs"
darkenpath=$(adret darkenpath -key="toto" $1)
ubac_result=$(ubac cat -path=encrypted.arafs $darkenpath)
adret decryptcat -key="toto" $ubac_result
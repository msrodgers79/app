#!/bin/bash
sshcmd="ssh -t mrodgers@app.msrodgers.co.uk"
$sshcmd screen -S "deployment" /home/mrodgers/app/prod_deploy.sh

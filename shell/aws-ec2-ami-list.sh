#!/bin/bash
export aws_region="$1"
export ami_name="$2"
export days_backwards="$2"

# get the day 20 days before today
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
START_DATE=`$DIR/get-start-date.sh $days_backwards`

printf "%s\n" "Listing AMIs for region ${aws_region} since ${START_DATE}" 

aws ec2 describe-images --region "${aws_region}" --output table \
--filters Name=is-public,Values=false \
--query "sort_by(Images[?CreationDate>=\`${START_DATE}\`],&CreationDate)[].[Name,ImageId,ImageType,CreationDate,State]"

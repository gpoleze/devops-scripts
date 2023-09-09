#!/bin/bash
export aws_region="${1:-us-east-1}"

printf "%s\n" "Listing Certificates for region ${aws_region}" 

aws acm list-certificates \
	--region ${aws_region}  \
	--output table \
	--query 'CertificateSummaryList[*].[DomainName,Type,CertificateArn]'


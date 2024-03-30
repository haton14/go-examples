#!/bin/bash
awslocal s3api head-bucket --bucket ${BUCKET_NAME} --region ${AWS_REGION} || awslocal s3api create-bucket --bucket ${BUCKET_NAME} --create-bucket-configuration LocationConstraint=${AWS_REGION}
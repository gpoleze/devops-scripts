#!/usr/bin/python3

import argparse
import os
import subprocess
import json
from prettytable import PrettyTable
from functional import seq


def main(region, repository_name):
    print(
        f"Getting the list of images in {region} for Repository {repository_name}")

    command = f"aws ecr describe-images"\
        f" --repository-name={repository_name}"\
        f" --region={region}" \
        f" --output=json"

    process = subprocess.Popen(command.split(' '),
                               stdout=subprocess.PIPE,
                               stderr=subprocess.PIPE)
    stdout, stderr = process.communicate()
    if stderr:
        print(stderr)
        exit(1)

    images = json.loads(stdout)["imageDetails"]

    tags = []
    for image in images:
        if 'imageTags' not in image.keys():
            continue

        for tag in image['imageTags']:
            tags.append([tag, image["imagePushedAt"]])

    tab = PrettyTable(['tag', 'pushed at'])
    tab.add_rows(tags)
    print(tab)


def get_regions():
    with open("aws-regions.lst", 'r') as file:
        lines = file.readlines()
        return seq(lines)\
            .map(lambda line: line.split("|"))\
            .map(lambda x: x[0])\
            .set()


def parse_args():
    parser = argparse.ArgumentParser(
        prog='aws-ecr-list-images',
        description='List images in a given ECR repository')

    region=os.getenv("AWS_REGION")

    parser.add_argument(
        "--region", default=region, choices=get_regions())
    parser.add_argument("--repository-name", required=True)

    args = parser.parse_args()
    if not region and not args.region:
            parser.error("The region cannot be empty. either declare it defining the AWS_REGION environment variable or using the flag '--region'")
    return args


if __name__ == "__main__":
    args = parse_args()
    main(args.region, args.repository_name)



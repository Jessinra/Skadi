# Cloud Formation Scripts

This directory contains AWS Cloud Formation Scripts related to parnassus.

## How To Run

Modify and execute snippet below:
1. use `create-stack` if the stack doesn't exist (first time creating the stack), use `update-stack` otherwise
2. provide parameters: 
    - `<aws profile>` to your local AWS CLI profile
    - `<stack name>` desired cloudformation stack name
    - `<template.yml>` path to the `.yml` file containing the cloudformation script (e.g. `script/cloudformation/myscript.yml`)
    - `<parameters.json>` path to the `.json` file containing the cloudformation script parameters (e.g. `script/cloudformation/myscript-parameters.json`)

```bash
aws cloudformation <create-stack/update-stack> \
--profile <aws profile> \
--stack-name <stack name> \
--template-body file://<template.yml>  \
--parameters file://<parameters.json> \
--capabilities "CAPABILITY_IAM" "CAPABILITY_NAMED_IAM"
```

Example of complete script

```bash
aws cloudformation create-stack \
--profile skadi-dev \
--stack-name MyStack \
--template-body file://script/cloudformation/mystack.yml  \
--parameters file://script/cloudformation/mystack-parameters.json \
--capabilities "CAPABILITY_IAM" "CAPABILITY_NAMED_IAM"
```

aws cloudformation create-stack \
--profile trivery-dev \
--stack-name skadi-dev-network \
--template-body file://script/cloudformation/network.yml  \
--parameters file://script/cloudformation/network-dev-param.json \
--capabilities "CAPABILITY_IAM" "CAPABILITY_NAMED_IAM"
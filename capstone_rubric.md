# UDCT Capstone Rubric

## Setup Pipeline

- Github repository: https://github.com/Jessinra/TriveryID-Skadi
- Dockerhub: https://hub.docker.com/r/trivery/skadi

## Build Docker Container

- I use 2 linter, golangci-lint (for golang code) and hadolint (dockerfile), both run on circleci 'lint' job.
    - for screenshots see `screenshots` folder
- This project has a Dockerfile, which is used to build the main golang code, run on circleci 'build' job, and push to the dockerhub with latest tag and with commit-hash tag.

## Successful Deployment

- I mainly use AWS elasticbeanstalk (EB) to deploy the application, but for this capstone project I also created another deployment using AWS EKS.
    - Both the EB application and EKS cluster are deployed with CloudFormation, including the networking (VPC, subnets, sg, etc.)
    - I choose to deploy the EKS nodegroup and setup the EKS cluster manually (there's so much to do and it's easier to follow using aws docs).
    - All script available in `script/cloudformation` folder.
- For both EB and EKS deployment, I use rolling update strategy
    - EB -> see the CF script, it has `aws:autoscaling:updatepolicy:rollingupdate` enabled, see screenshots in the `screenshots` folder.
    - EKS -> set node max unavailability to 50%, and use `kubectl rollout restart`, see screenshots in the `screenshots` folder.
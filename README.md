# Source for https://practicalgobook.net

Hugo + go:embed magic.

## Infrastructure

AWS load balancer (443) -> Private EC2 instance (8080) -> Go server

## Deploying

```
ssh-add -K ~/.ssh/<ssh-key>
ssh -A -i ~/.ssh/<ssh key> ec2-user@ec2-3-25-87-17.ap-southeast-2.compute.amazonaws.com
```

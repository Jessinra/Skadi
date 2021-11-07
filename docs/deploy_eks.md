# Deploy to EKS

## Create a new EKS cluster
Follow https://docs.aws.amazon.com/eks/latest/userguide/create-cluster.html


## Add Node Groups
(might not needed)
Follow https://docs.aws.amazon.com/eks/latest/userguide/create-managed-node-group.html

## Configure kubectl
```bash
aws eks --region ap-southeast-1 update-kubeconfig --name skadi
```

Result
```
➜  skadi git:(master) kubectl get svc                                               
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   172.20.0.1   <none>        443/TCP   36m
```

Continue from here https://docs.aws.amazon.com/eks/latest/userguide/sample-deployment.html
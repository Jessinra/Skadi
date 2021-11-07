aws iam create-role \
  --role-name eksSkadiCNIRole \
  --assume-role-policy-document file://"cni-role-trust-policy.json"

aws iam attach-role-policy \
  --policy-arn arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy \
  --role-name eksSkadiCNIRole

aws eks update-addon \
  --region ap-southeast-1 \
  --cluster-name skadi \
  --addon-name vpc-cni \
  --service-account-role-arn arn:aws:iam::442662599070:role/eksSkadiCNIRole

aws iam create-role \
  --role-name eksSkadiNodeRole \
  --assume-role-policy-document file://"node-role-trust-policy.json"

aws iam attach-role-policy \
  --policy-arn arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy \
  --role-name eksSkadiNodeRole
aws iam attach-role-policy \
  --policy-arn arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly \
  --role-name eksSkadiNodeRole


# https://docs.aws.amazon.com/eks/latest/userguide/cni-iam-role.html# annotate
kubectl annotate serviceaccount \
  -n kube-system aws-node \
  eks.amazonaws.com/role-arn=arn:aws:iam::442662599070:role/eksSkadiCNIRole --overwrite


# ============ Enable autoscale ================
aws iam create-policy \
    --policy-name AmazonEKSClusterAutoscalerPolicy \
    --policy-document file://"cluster-autoscaler-policy.json"

kubectl annotate serviceaccount cluster-autoscaler \
  -n kube-system \
  eks.amazonaws.com/role-arn=arn:aws:iam::442662599070:role/eksSkadiClusterAutoscaler

kubectl patch deployment cluster-autoscaler \
  -n kube-system \
  -p '{"spec":{"template":{"metadata":{"annotations":{"cluster-autoscaler.kubernetes.io/safe-to-evict": "false"}}}}}'

kubectl set image deployment cluster-autoscaler \
  -n kube-system \
  cluster-autoscaler=k8s.gcr.io/autoscaling/cluster-autoscaler:v1.21.1
# ==============================================